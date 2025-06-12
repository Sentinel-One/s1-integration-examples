package scanner

import (
	"encoding/json"
	"time"
)

// Result holds the result of the scan of an individual file.
type Result struct {
	// Errors holds a list of any errors that occurred during the scan.
	Errors []string `json:"errors"`

	// FileHash is the SHA256 of the file contents.
	FileHash string `json:"file_hash"`

	// FileName is the actual name of the file by itself.
	FileName string `json:"file_name"`

	// Indicators holds a list of indicators that were found in the file that were used to make the decision for
	// the verdict.
	Indicators []string `json:"indicators"`

	// Message contains any additional information to pass back to the requester.
	Message string `json:"message"`

	// ScanDuration indicates the amount of time it took to scan the file.
	ScanDuration Duration `json:"scan_duration"`

	// Verdict indicates the verdict of whether the file is infected, suspicious, benign, etc. This field is
	// interpreted by the
	Verdict Verdict `json:"verdict"`
}

// Duration is just a wrapper for a [time.Duration] object that, when marshaled, produces a duration string rather
// than a raw number.
type Duration time.Duration

// MarshalJSON returns the duration as an elapsed time string formatted as JSON.
func (d Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Duration(d).String())
}

// MarshalText returns the duration as a elapsed time string formatted in plain text.
func (d Duration) MarshalText() ([]byte, error) {
	return []byte(time.Duration(d).String()), nil
}

// String returns the elapsed time as a string.
func (d Duration) String() string {
	return time.Duration(d).String()
}
