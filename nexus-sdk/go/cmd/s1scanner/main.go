package main

import (
	"fmt"
	"log/slog"
	"os"
)

func main() {
	exitCode := 0
	err := NewRootCommand().Execute()
	if e, ok := err.(ErrorX); ok {
		exitCode = e.Code
	} else if err != nil {
		// error returned was not an "extended" error so treat it as a usage error
		exitCode = ErrUsage
	}
	if exitCode > 0 && exitCode != ErrUsage {
		slog.Default().Warn(fmt.Sprintf("exiting with non-zero exit code %d", exitCode),
			slog.Int("exit_code", exitCode))
	}
	os.Exit(exitCode)
}
