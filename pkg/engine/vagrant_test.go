package engine_test

import (
	"testing"

	"github.com/mattjmcnaughton/tmplinux/pkg/engine"
	"github.com/mattjmcnaughton/tmplinux/pkg/executor"
	"github.com/mattjmcnaughton/tmplinux/pkg/reporter"
)

func TestVagrantEngineStartCommandSuccess(t *testing.T) {
	mockExecutor := executor.CreateSuccessMockExecutor()
	mockReporter := reporter.MockReporter{}

	testVagrantEngine := engine.NewCustomVagrantEngine(&mockExecutor, &mockReporter)
	testVagrantEngine.Start()

	assertKeywordIncludedInCommand(t, mockExecutor, "init")
	assertCommandIssuedInSubdirectoryOf(t, mockExecutor, "/tmp")
	assertKeywordIncludedInCommand(t, mockExecutor, "up")
	assertCommandIssuedInSubdirectoryOf(t, mockExecutor, "/tmp")

	assertReporterNotCalled(t, mockReporter)
}

func TestVagrantEngineCommandFailed(t *testing.T) {
	mockExecutor := executor.CreateFailureMockExecutor()
	mockReporter := reporter.MockReporter{}

	testVagrantEngine := engine.NewCustomVagrantEngine(&mockExecutor, &mockReporter)
	testVagrantEngine.Start()

	assertReporterCalled(t, mockReporter)
}

func TestVagrantEngineSSHCommandSuccess(t *testing.T) {
	mockExecutor := executor.CreateSuccessMockExecutor()
	mockReporter := reporter.MockReporter{}

	testVagrantEngine := engine.NewCustomVagrantEngine(&mockExecutor, &mockReporter)
	testVagrantEngine.SSH()

	assertKeywordIncludedInCommand(t, mockExecutor, "ssh")
	assertCommandIssuedInSubdirectoryOf(t, mockExecutor, "/tmp")
	assertReporterNotCalled(t, mockReporter)
}

func TestVagrantEngineStopCommandSuccess(t *testing.T) {
	mockExecutor := executor.CreateSuccessMockExecutor()
	mockReporter := reporter.MockReporter{}

	testVagrantEngine := engine.NewCustomVagrantEngine(&mockExecutor, &mockReporter)
	testVagrantEngine.Stop()

	assertKeywordIncludedInCommand(t, mockExecutor, "suspend")
	assertCommandIssuedInSubdirectoryOf(t, mockExecutor, "/tmp")
	assertReporterNotCalled(t, mockReporter)
}

func TestVagrantEngineRmCommandSuccess(t *testing.T) {
	mockExecutor := executor.CreateSuccessMockExecutor()
	mockReporter := reporter.MockReporter{}

	testVagrantEngine := engine.NewCustomVagrantEngine(&mockExecutor, &mockReporter)
	testVagrantEngine.Rm()

	assertKeywordIncludedInCommand(t, mockExecutor, "destroy -f")
	assertCommandIssuedInSubdirectoryOf(t, mockExecutor, "/tmp")
	assertReporterNotCalled(t, mockReporter)

	assertKeywordIncludedInCommand(t, mockExecutor, "rm")
	assertCommandIssuedInSubdirectoryOf(t, mockExecutor, "/tmp")
	assertReporterNotCalled(t, mockReporter)
}

func TestVagrantEngineValidateCommandSuccess(t *testing.T) {
	mockExecutor := executor.CreateSuccessMockExecutor()
	mockReporter := reporter.MockReporter{}

	testVagrantEngine := engine.NewCustomVagrantEngine(&mockExecutor, &mockReporter)
	testVagrantEngine.Validate()

	assertKeywordIncludedInCommand(t, mockExecutor, "global-status")
	assertCommandIssuedInSubdirectoryOf(t, mockExecutor, "/tmp")
	assertReporterNotCalled(t, mockReporter)
}
