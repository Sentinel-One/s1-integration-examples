package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/lmittmann/tint"
	"github.com/mattn/go-colorable"
	"github.com/mattn/go-isatty"
	"github.com/spf13/cobra"

	"s1scanner-example/pkg/scanner"
)

const (
	_SampleFilesFolder = "/opt/s1scanner/sample-files" // path to sample files
	_BaseScanFolder    = "/mnt"                        // root folder where host filesystem is mounted
)

// RootCommand is the base command for the application.
type RootCommand struct {
	cobra.Command

	baseScanFolder string // root folder from which to base file paths
	debug          bool   // enable debug mode
	jsonOutput     bool   // use JSON output
}

// NewRootCommand creates a new Command object.
func NewRootCommand() *RootCommand {
	cmd := &RootCommand{}
	cmd.Use = "s1scanner [flags] [file_or_dir ...]"
	cmd.Short = "Runs the s1scanner"
	cmd.Long = "Scan a file or directory using the s1scanner"
	cmd.RunE = cmd.runE

	// add flags
	flags := cmd.Flags()
	flags.String("base-scan-folder", "/mnt", "override the root volume mount path")
	flags.MarkHidden("base-scan-folder")
	flags.Bool("debug", false, "enable debug mode")
	flags.Bool("demo", false, "runs the scanner demo and exits")
	flags.Int("max-depth", 2, "maximum depth to scan inside archive files (more than 10 is not recommended)")
	flags.String("max-scan-duration", "5s", "maximum duration to spend scanning a single file")
	flags.BoolP("json", "j", false, "print scan results as JSON instead of text")
	flags.BoolP("recurse", "r", false, "recursively scan into directories")
	flags.BoolP("version", "v", false, "print the version and exit")

	return cmd
}

// runE is responsible for actually running the application.
func (c *RootCommand) runE(cmd *cobra.Command, args []string) error {
	if cmd == nil {
		panic("cmd parameter should never be nil")
	}
	if args == nil {
		panic("args parameter should never be nil")
	}
	flags := cmd.Flags()

	// initialize the output logger
	c.debug, _ = flags.GetBool("debug")
	if c.jsonOutput, _ = flags.GetBool("json"); c.jsonOutput {
		c.createJSONLogger()
	} else {
		c.createTextLogger()
	}

	// override base scan folder
	c.baseScanFolder, _ = flags.GetString("base-scan-folder")

	// version only
	if showVersion, _ := flags.GetBool("version"); showVersion {
		cmd.SilenceErrors = true
		cmd.SilenceUsage = true
		return c.showVersion()
	}

	// demo mode
	if runDemo, _ := flags.GetBool("demo"); runDemo {
		cmd.SilenceErrors = true
		cmd.SilenceUsage = true
		options := []func(*scanner.Option){}
		return c.scan("", _SampleFilesFolder, options, true)
	}

	// make sure we have a one or more args left
	if len(args) == 0 {
		errx := ErrorX{
			error: errors.New("you must specify 1 or more files or folders to be scanned"),
			Code:  ErrUsage,
		}
		return errx
	}
	cmd.SilenceErrors = true
	cmd.SilenceUsage = true

	// setup scan options
	scanOptions := []func(*scanner.Option){}
	depth, _ := flags.GetInt("max-depth")
	scanOptions = append(scanOptions, scanner.WithMaxScanDepth(depth))
	duration, _ := flags.GetString("max-scan-duration")
	maxScanDuration, err := time.ParseDuration(duration)
	if err != nil {
		errx := ErrorX{
			error: fmt.Errorf("failed to parse maximum scan duration: %w", err),
			Code:  ErrUsage,
		}
		return errx
	}
	scanOptions = append(scanOptions, scanner.WithTimeout(maxScanDuration))
	recurse, _ := flags.GetBool("recurse")
	slog.Default().Debug("configured scan options",
		slog.Int("max_scan_depth", depth),
		slog.String("max_scan_duration", maxScanDuration.String()),
		slog.Bool("recursive_scan", recurse),
	)

	// scan each of the files/folders specified
	var lastError *ErrorX
	for _, path := range args {
		if errx := c.scan(c.baseScanFolder, path, scanOptions, recurse); errx != nil {
			lastError = errx
		}
	}
	return lastError
}

// createTextLogger creates a pretty text logger and sets it as default.
func (c *RootCommand) createTextLogger() {
	level := slog.LevelInfo
	addSource := false
	if c.debug {
		level = slog.LevelDebug
		addSource = true
	}
	handlerOptions := &tint.Options{
		AddSource:  addSource,
		NoColor:    !isatty.IsTerminal(os.Stdout.Fd()),
		Level:      level,
		TimeFormat: "03:04:05PM",
	}
	handler := tint.NewHandler(colorable.NewColorable(os.Stdout), handlerOptions)
	logger := slog.New(handler)
	slog.SetDefault(logger)
}

// createJSONLogger creates a JSON logger and sets it as default.
func (c *RootCommand) createJSONLogger() {
	level := slog.LevelInfo
	addSource := false
	if c.debug {
		level = slog.LevelDebug
		addSource = true
	}
	handlerOptions := &slog.HandlerOptions{
		AddSource: addSource,
		Level:     level,
	}
	handler := slog.NewJSONHandler(os.Stdout, handlerOptions)
	logger := slog.New(handler)
	slog.SetDefault(logger)
}

// logResult simply sends the result to the output logger.
func (c *RootCommand) logResult(result scanner.Result) {
	var msg string
	attrs := []slog.Attr{
		slog.String("file_name", result.FileName),
		slog.String("file_hash", result.FileHash),
		slog.String("verdict", result.Verdict.String()),
		slog.String("scan_duration", result.ScanDuration.String()),
	}
	if result.Message != "" {
		attrs = append(attrs, slog.String("info", result.Message))
	}
	if len(result.Indicators) > 0 {
		attrs = append(attrs, slog.Any("indicators", result.Indicators))
	}
	if len(result.Errors) > 0 {
		msg = "file scan finished with errors"
		attrs = append(attrs, slog.Any("errors", result.Errors))
	} else {
		msg = "file scan finished successfully"
	}
	slog.Default().LogAttrs(context.Background(), slog.LevelInfo, msg, attrs...)
}

// scan submits the file(s) or folder(s) to the scan engine for scanning.
func (c *RootCommand) scan(basePath, scanPath string, opts []func(*scanner.Option), recurse bool) *ErrorX {
	logger := slog.Default()

	// the actual scan path inside the container is the basePath + scanPath
	// make sure it exists and get info about it to determine if it's a file or directory
	path := path.Join(basePath, scanPath)
	logger.Debug("scan requested", slog.String("container_path", path))
	matches, err := filepath.Glob(path)
	if err != nil {
		errx := &ErrorX{
			error: fmt.Errorf("failed to scan the file/folder '%s': %w", scanPath, err),
			Code:  ErrFileGlob,
		}
		logger.Error(errx.Error())
		return errx
	}
	logger.Debug(fmt.Sprintf("matched %d files/folders", len(matches)), slog.String("container_path", path))
	if len(matches) == 0 {
		err := errors.New("no matching files or folders were found")
		errx := &ErrorX{
			error: fmt.Errorf("failed to scan the file/folder '%s': %w", scanPath, err),
			Code:  ErrNoMatchingFiles,
		}
		logger.Error(errx.Error())
		return errx
	}

	// initialize the engine
	engine := scanner.NewEngine()
	defer engine.Cleanup()
	if err := engine.Init(); err != nil {
		errx := &ErrorX{
			error: fmt.Errorf("failed to initialize the scan engine: %w", err),
			Code:  ErrEngineInit,
		}
		logger.Error(errx.Error())
		return errx
	}

	// scan matching files/folders
	for _, match := range matches {
		fileInfo, err := os.Stat(match)
		if err != nil {
			path := strings.TrimPrefix(match, basePath)
			errx := &ErrorX{
				error: fmt.Errorf("failed to scan the file/folder '%s': %w", path, err),
				Code:  ErrFileStat,
			}
			logger.Error(errx.Error())
			return errx
		}

		// just scan a file
		if !fileInfo.IsDir() {
			result, err := engine.ScanFile(match, opts...)
			if err != nil {
				path := strings.TrimPrefix(match, basePath)
				errx := &ErrorX{
					error: fmt.Errorf("failed to scan the file/folder '%s': %w", path, err),
					Code:  ErrScanFile,
				}
				logger.Error(errx.Error())
				return errx
			}
			c.logResult(result)
			continue
		}

		// walk the directory
		if err := filepath.WalkDir(match, func(path string, d os.DirEntry, err error) error {
			logger.Debug("walking directory", slog.String("container_path", path))
			if d.Name() == "." || d.Name() == ".." { // skip self and parent
				return nil
			}
			if !recurse && d.IsDir() && path != match {
				logger.Debug("skipping directory as recursion is disabled", slog.String("container_path", path))
				return filepath.SkipDir
			}
			if !d.IsDir() {
				result, err := engine.ScanFile(path, opts...)
				if err != nil {
					path := strings.TrimPrefix(path, basePath)
					errx := &ErrorX{
						error: fmt.Errorf("failed to scan the file/folder '%s': %w", path, err),
						Code:  ErrScanFile,
					}
					logger.Error(errx.Error())
					return errx
				}
				c.logResult(result)
			}
			return nil
		}); err != nil {
			path := strings.TrimPrefix(match, basePath)
			errx := &ErrorX{
				error: fmt.Errorf("failed to scan the file/folder '%s': %w", path, err),
				Code:  ErrWalkDir,
			}
			logger.Error(errx.Error())
			return errx
		}
	}
	return nil
}

// showVersion prints the version of the SDK and exits
func (c *RootCommand) showVersion() error {
	engine := scanner.NewEngine()
	defer engine.Cleanup()

	version, hash, err := engine.Version()
	if err != nil {
		errx := ErrorX{
			error: fmt.Errorf("failed to retrieve SDK version: %w", err),
			Code:  ErrEngineVersion,
		}
		slog.Default().Error(errx.Error())
		return errx
	}
	fmt.Printf("Nexus SDK version %s (%s)\n", version, hash)
	return nil
}
