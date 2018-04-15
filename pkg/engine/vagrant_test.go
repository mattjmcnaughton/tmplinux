package engine_test

import (
	"testing"

	"github.com/mattjmcnaughton/tmplinux/pkg/engine"
	"github.com/mattjmcnaughton/wutils/pkg/executor"
	"github.com/mattjmcnaughton/wutils/pkg/reporter"
)

func TestVagrantEngineStartCommandSuccess(t *testing.T) {
	mockExecutor := executor.CreateSuccessMockExecutor()
	mockReporter := reporter.MockReporter{}

	testVagrantEngine := engine.NewCustomVagrantEngine(&mockExecutor, &mockReporter)
	testVagrantEngine.Start()

	mockExecutor.AssertKeywordIncludedInCommand(t, "init")
	mockExecutor.AssertCommandIssuedInSubdirectoryOf(t, "/tmp")
	mockExecutor.AssertKeywordIncludedInCommand(t, "up")
	mockExecutor.AssertCommandIssuedInSubdirectoryOf(t, "/tmp")
	mockReporter.AssertNotCalled(t)
}

func TestVagrantEngineCommandFailed(t *testing.T) {
	mockExecutor := executor.CreateFailureMockExecutor()
	mockReporter := reporter.MockReporter{}

	testVagrantEngine := engine.NewCustomVagrantEngine(&mockExecutor, &mockReporter)
	testVagrantEngine.Start()

	mockReporter.AssertCalled(t)
}

func TestVagrantEngineSSHCommandSuccess(t *testing.T) {
	mockExecutor := executor.CreateSuccessMockExecutor()
	mockReporter := reporter.MockReporter{}

	testVagrantEngine := engine.NewCustomVagrantEngine(&mockExecutor, &mockReporter)
	testVagrantEngine.SSH()

	mockExecutor.AssertKeywordIncludedInCommand(t, "ssh")
	mockExecutor.AssertCommandIssuedInSubdirectoryOf(t, "/tmp")
	mockReporter.AssertNotCalled(t)
}

func TestVagrantEngineStopCommandSuccess(t *testing.T) {
	mockExecutor := executor.CreateSuccessMockExecutor()
	mockReporter := reporter.MockReporter{}

	testVagrantEngine := engine.NewCustomVagrantEngine(&mockExecutor, &mockReporter)
	testVagrantEngine.Stop()

	mockExecutor.AssertKeywordIncludedInCommand(t, "suspend")
	mockExecutor.AssertCommandIssuedInSubdirectoryOf(t, "/tmp")
	mockReporter.AssertNotCalled(t)
}

func TestVagrantEngineRmCommandSuccess(t *testing.T) {
	mockExecutor := executor.CreateSuccessMockExecutor()
	mockReporter := reporter.MockReporter{}

	testVagrantEngine := engine.NewCustomVagrantEngine(&mockExecutor, &mockReporter)
	testVagrantEngine.Rm()

	mockExecutor.AssertKeywordIncludedInCommand(t, "destroy -f")
	mockExecutor.AssertCommandIssuedInSubdirectoryOf(t, "/tmp")
	mockReporter.AssertNotCalled(t)

	mockExecutor.AssertKeywordIncludedInCommand(t, "rm")
	mockExecutor.AssertCommandIssuedInSubdirectoryOf(t, "/tmp")
	mockReporter.AssertNotCalled(t)
}

func TestVagrantEngineValidateCommandSuccess(t *testing.T) {
	mockExecutor := executor.CreateSuccessMockExecutor()
	mockReporter := reporter.MockReporter{}

	testVagrantEngine := engine.NewCustomVagrantEngine(&mockExecutor, &mockReporter)
	testVagrantEngine.Validate()

	mockExecutor.AssertKeywordIncludedInCommand(t, "global-status")
	mockExecutor.AssertCommandIssuedInSubdirectoryOf(t, "/tmp")
	mockReporter.AssertNotCalled(t)
}
