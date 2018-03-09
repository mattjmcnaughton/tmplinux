package reporter

import (
	"fmt"

	"os"
)

type Reporter interface {
	ReportIfError(error, string, ...interface{})
}

// Can additionally have a mock one for testing.
type FmtReporter struct {
}

func (f FmtReporter) ReportIfError(err error, format string, a ...interface{}) {
	if err != nil {
		fmt.Printf(fmt.Sprintf("%s\n", format), a...)
		fmt.Printf("Error: %v\n", err)

		// @TODO(mattjmcnaughton) Make exit code more accurate.
		os.Exit(1)
	}
}
