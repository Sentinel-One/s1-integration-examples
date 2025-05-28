package scanner

// #cgo CFLAGS: -g -Wall -I${SRCDIR}/nexus/include
// #cgo linux arm64 LDFLAGS: -L${SRCDIR}/nexus/lib/linux/arm64 -ldfi
// #cgo linux amd64 LDFLAGS: -L${SRCDIR}/nexus/lib/linux/amd64 -ldfi
// #cgo windows amd64 LDFLAGS: -L${SRCDIR}/nexus/lib/windows/amd64 -ldfi
// #cgo windows 386 LDFLAGS: -L${SRCDIR}/nexus/lib/windows/x86 -ldfi
// #include <stdio.h>
// #include <stdint.h>
// #include <stdlib.h>
// #include "libdfi.h"
import "C"
import (
	"bytes"
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
	"unsafe"
)

// Engine implements the SentinelOne Static AI file scanning engine.
type Engine struct {
	isInit bool
}

// NewEngine creates a new [Engine] object.
//
// Be sure to defer a call to [Engine.Cleanup] when it is no longer needed in order to clean up any memory allocated
// by the scan engine.
func NewEngine() *Engine {
	return &Engine{}
}

// Cleanup should be called to cleanup any memory allocated by the scan engine.
func (e *Engine) Cleanup() {
	if e.isInit {
		C.cleanup()
	}
}

// Init must be called to initialize the scan engine prior to scanning any files.
func (e *Engine) Init() error {
	if e.isInit {
		return nil
	}

	// initialize Nexus SDK
	result := C.init()
	if result != C.DFI_SUCCESS {
		return fmt.Errorf("static AI engine initialization returned error code %d", int(result))
	}

	e.isInit = true
	return nil
}

// IsInitialized returns whether or not the scan engine has already been initialized.
func (e *Engine) IsInitialized() bool {
	return e.isInit
}

// ScanBytes scans the raw bytes passed into the function and returns the result.
//
// You must pass the entire contents of the file to be scanned through the buffer. You can not perform a partial
// scan of the contents using this function.
//
// Be sure to call the [Engine.Init] function to initialze the scan engine before calling this function.
func (e *Engine) ScanBytes(buf *bytes.Buffer, options ...func(*Option)) (Result, error) {
	result := Result{
		Errors:     []string{},
		FileHash:   fmt.Sprintf("%x", sha256.Sum256(buf.Bytes())),
		FileName:   "<unknown>",
		Indicators: []string{},
		Verdict:    VerdictUnknown,
	}
	if !e.isInit {
		return result, errors.New("scan engine has not been initialized")
	}
	if buf == nil {
		buf = bytes.NewBuffer([]byte{})
	}

	// initialize options
	scanOptions := newOption()
	for _, o := range options {
		o(scanOptions)
	}

	// prep memory for sharing
	var scanStatusCode C.int
	var verdict *C.char
	contents := (*C.uint8_t)(C.CBytes(buf.Bytes()))
	defer C.free(unsafe.Pointer(contents))
	indicators := (*C.char)(C.CBytes(make([]byte, 512)))
	defer C.free(unsafe.Pointer(indicators))

	// scan the file
	now := time.Now()
	if scanOptions.maxScanTime > 0 {
		ctx, cancel := context.WithTimeout(context.Background(), scanOptions.maxScanTime)
		defer cancel()
		resultCh := make(chan C.int)
		go func() {
			resultCh <- C.scan_file_with_depth(contents, C.uint(buf.Len()), 0, &verdict, indicators, 512,
				C.uint(scanOptions.maxScanDepth))
		}()
		select {
		case scanStatusCode = <-resultCh:
			break
		case <-ctx.Done():
			return result, fmt.Errorf("timeout while scanning file: %w", ctx.Err())
		}
	} else {
		scanStatusCode = C.scan_file_with_depth(contents, C.uint(buf.Len()), 0, &verdict, indicators, 512,
			C.uint(scanOptions.maxScanDepth))
	}
	result.ScanDuration = Duration(time.Since(now))

	// parse the results
	var err error
	switch scanStatusCode {
	case C.DFI_SUCCESS:
		if C.GoString(indicators) != "" {
			result.Indicators = strings.Split(C.GoString(indicators), ",")
		}
		switch strings.ToLower(C.GoString(verdict)) {
		case "benign":
			result.Verdict = VerdictBenign
		case "suspicious":
			result.Verdict = VerdictSuspicious
		case "malware":
			result.Verdict = VerdictMalicious
		}
	case C.DFI_NOT_INITIALIZED: // should NOT happen
		err = errors.New("can engine not initialized")
		result.Errors = append(result.Errors, err.Error())
		result.Message = err.Error()
	case C.DFI_WRONG_PARAMETER:
		err = errors.New("one or more parameters were not initialized correctly")
		result.Errors = append(result.Errors, err.Error())
		result.Message = err.Error()
	case C.DFI_TIMEOUT_REACHED:
		err = errors.New("timeout reached while scanning the file")
		result.Errors = append(result.Errors, err.Error())
		result.Message = err.Error()
	case C.DFI_NO_LICENSE:
		err = errors.New("no SentinelOne license was found")
		result.Errors = append(result.Errors, err.Error())
		result.Message = err.Error()
	case C.DFI_MINIMUM_FILE_SIZE_ERROR:
		err = errors.New("the file is too small for scanning")
		result.Errors = append(result.Errors, err.Error())
		result.Message = err.Error()
	case C.DFI_UNKNOWN_FILE_TYPE:
		err = errors.New("the file type could not be determined")
		result.Errors = append(result.Errors, err.Error())
		result.Message = err.Error()
	case C.DFI_DATA_IS_MISSING:
		err = errors.New("data is missing from the file")
		result.Errors = append(result.Errors, err.Error())
		result.Message = err.Error()
	case C.DFI_FILE_CORRUPTED:
		err = errors.New("the file appears to be corrupted")
		result.Errors = append(result.Errors, err.Error())
		result.Message = err.Error()
	case C.DFI_INVALID_ARCHIVE:
		err = errors.New("%s: the archive format could not be scanned")
		result.Errors = append(result.Errors, err.Error())
		result.Message = err.Error()
	case C.DFI_MAX_ARCHIVE_DEPTH_ERROR:
		err = errors.New("the maximum depth was reached while scanning the archive file")
		result.Errors = append(result.Errors, err.Error())
		result.Message = err.Error()
	case C.DFI_FILE_TOO_SHORT_FOR_VECTOR:
		err = errors.New("the file is too short for vector")
		result.Errors = append(result.Errors, err.Error())
		result.Message = err.Error()
	case C.DFI_UNKNOWN_VECTOR_TYPE:
		err = errors.New("unknown vector type")
		result.Errors = append(result.Errors, err.Error())
		result.Message = err.Error()
	case C.DFI_SECTION_INDEX_OUT_OF_RANGE:
		err = errors.New("out of range error during parsing")
		result.Errors = append(result.Errors, err.Error())
		result.Message = err.Error()
	case C.DFI_SECTION_INDEX_OUT_OF_RANGE_DISASM:
		err = errors.New("out of range error during disassembly")
		result.Errors = append(result.Errors, err.Error())
		result.Message = err.Error()
	case C.DFI_SIGNATURES_SCAN_FAILURE:
		err = errors.New("signature scanning failed")
		result.Errors = append(result.Errors, err.Error())
		result.Message = err.Error()
	case C.DFI_INDICATORS_SCAN_FAILURE:
		err = errors.New("scanning for static indicators failed")
		result.Errors = append(result.Errors, err.Error())
		result.Message = err.Error()
	case C.DFI_RUNTIME_ERROR:
		err = errors.New("a runtime error occurred while scanning the file")
		result.Errors = append(result.Errors, err.Error())
		result.Message = err.Error()
	case C.DFI_LENGTH_ERROR:
		err = errors.New("internal file length error")
		result.Errors = append(result.Errors, err.Error())
		result.Message = err.Error()
	case C.DFI_BAD_ALLOC_ERROR:
		err = errors.New("bad allocation during file parsing")
		result.Errors = append(result.Errors, err.Error())
		result.Message = err.Error()
	case C.DFI_SFX_TYPE_UNSUPPORTED:
		err = errors.New("self-extracting file type not supported")
		result.Errors = append(result.Errors, err.Error())
		result.Message = err.Error()
	case C.DFI_SFX_ARCHIVE_MISSING:
		err = errors.New("could not extract self-extracting archive")
		result.Errors = append(result.Errors, err.Error())
		result.Message = err.Error()
	case C.DFI_EML_INTERNAL_ERROR:
		err = errors.New("%s: internal error parsing email message")
		result.Errors = append(result.Errors, err.Error())
		result.Message = err.Error()
	default:
		err = fmt.Errorf("an unexpected error has occurred (code: %d)", scanStatusCode)
		result.Errors = append(result.Errors, err.Error())
		result.Message = err.Error()
	}
	return result, nil
}

// ScanFile opens and scans the given file and returns the result.
//
// Be sure to call the [Engine.Init] function to initialze the scan engine before calling this function.
func (e *Engine) ScanFile(path string, options ...func(*Option)) (Result, error) {
	result := Result{
		Verdict: VerdictUnknown,
	}
	file, err := os.Open(path)
	if err != nil {
		return result, fmt.Errorf("failed to open the file '%s' for reading: %w", path, err)
	}
	defer file.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, file)
	if err != nil {
		return result, fmt.Errorf("failed to read the contents of the file '%s': %w", path, err)
	}
	result, err = e.ScanBytes(&buf)
	result.FileName = path
	return result, err
}

// Version returns the scan engine version number and hash.
func (e *Engine) Version() (string, string, error) {
	var version, hash *C.char
	C.get_version(&version, &hash)
	return C.GoString(version), C.GoString(hash), nil
}
