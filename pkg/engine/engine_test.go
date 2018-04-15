package engine_test

import (
	"strings"
	"testing"

	"github.com/mattjmcnaughton/wutils/pkg/executor"
	"github.com/mattjmcnaughton/wutils/pkg/reporter"
)

func assertKeywordIncludedInCommand(t *testing.T, mockExecutor executor.MockShellExecutor, keyword string) {
	execCmds := mockExecutor.GetExecutedCommands()

	anyMatches := false

	for _, cmd := range execCmds {
		if strings.Contains(cmd, keyword) {
			anyMatches = true
		}
	}

	if !anyMatches {
		t.Fatalf("%v should include the keyword %s", execCmds, keyword)
	}
}

func assertCommandIssuedInSubdirectoryOf(t *testing.T, mockExecutor executor.MockShellExecutor, parentDir string) {
	execCmdDir := mockExecutor.GetExecutedCommandDir()

	if !strings.HasPrefix(execCmdDir, parentDir) {
		t.Fatalf("Command executed in %s, which is not a subdir of %s", execCmdDir, parentDir)
	}
}

func assertReporterCalled(t *testing.T, mockReporter reporter.MockReporter) {
	if !mockReporter.Reported() {
		t.Fatalf("Reporter should have been called")
	}
}

func assertReporterNotCalled(t *testing.T, mockReporter reporter.MockReporter) {
	if mockReporter.Reported() {
		t.Fatalf("Reporter should not have been called")
	}
}
