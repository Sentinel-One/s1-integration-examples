package main

const (
	ErrUsage = iota + 1
	ErrEngineInit
	ErrEngineVersion
	ErrFileGlob
	ErrNoMatchingFiles
	ErrFileStat
	ErrScanFile
	ErrWalkDir
)

// ErrorX is just an extended error in order to include a unique error code.
type ErrorX struct {
	error

	// Code is a unique code for the error.
	Code int
}
