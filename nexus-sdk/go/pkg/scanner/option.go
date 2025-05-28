package scanner

import "time"

// Option is used for storing and passing around optional arguments for scanning.
type Option struct {
	maxScanDepth int           // maximum scan depth
	maxScanTime  time.Duration // maximum time to take to scan a file
}

// newOption creates a new [Option] object with default values.
func newOption() *Option {
	return &Option{
		maxScanDepth: 2,
		maxScanTime:  5 * time.Second,
	}
}

// WithMaxScanDepth limits the maximum depth of files to scan within an archive.
//
// The default maximum scan depth is 2. The maximum depth is only limited to the memory available to the OS, however,
// it is recommended not to go larger than 10.
func WithMaxScanDepth(depth int) func(*Option) {
	return func(o *Option) {
		o.maxScanDepth = depth
	}
}

// WithTimeout limits the amount of time scanning a file.
//
// The default maximum scan duration is 5 seconds.
func WithTimeout(scanTime time.Duration) func(*Option) {
	return func(o *Option) {
		o.maxScanTime = scanTime
	}
}
