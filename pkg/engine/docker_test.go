package engine_test

import (
	"testing"

	"github.com/mattjmcnaughton/tmplinux/pkg/engine"
	"github.com/mattjmcnaughton/wutils/pkg/executor"
	"github.com/mattjmcnaughton/wutils/pkg/reporter"
)

func TestDockerEngineStartCommandSuccess(t *testing.T) {
	mockExecutor := executor.CreateSuccessMockExecutor()
	mockReporter := reporter.MockReporter{}

	testDockerEngine := engine.NewCustomDockerEngine(&mockExecutor, &mockReporter)
	testDockerEngine.Start()

	assertKeywordIncludedInCommand(t, mockExecutor, "run")
	assertReporterNotCalled(t, mockReporter)
}

func TestDockerEngineCommandFailed(t *testing.T) {
	mockExecutor := executor.CreateFailureMockExecutor()
	mockReporter := reporter.MockReporter{}

	testDockerEngine := engine.NewCustomDockerEngine(&mockExecutor, &mockReporter)
	testDockerEngine.Start()

	assertReporterCalled(t, mockReporter)
}

func TestDockerEngineSSHCommandSuccess(t *testing.T) {
	mockExecutor := executor.CreateSuccessMockExecutor()
	mockReporter := reporter.MockReporter{}

	testDockerEngine := engine.NewCustomDockerEngine(&mockExecutor, &mockReporter)
	testDockerEngine.SSH()

	assertKeywordIncludedInCommand(t, mockExecutor, "exec")
	assertReporterNotCalled(t, mockReporter)
}

func TestDockerEngineStopCommandSuccess(t *testing.T) {
	mockExecutor := executor.CreateSuccessMockExecutor()
	mockReporter := reporter.MockReporter{}

	testDockerEngine := engine.NewCustomDockerEngine(&mockExecutor, &mockReporter)
	testDockerEngine.Stop()

	assertKeywordIncludedInCommand(t, mockExecutor, "kill")
	assertReporterNotCalled(t, mockReporter)
}

func TestDockerEngineRmCommandSuccess(t *testing.T) {
	mockExecutor := executor.CreateSuccessMockExecutor()
	mockReporter := reporter.MockReporter{}

	testDockerEngine := engine.NewCustomDockerEngine(&mockExecutor, &mockReporter)
	testDockerEngine.Rm()

	assertKeywordIncludedInCommand(t, mockExecutor, "rm")
	assertReporterNotCalled(t, mockReporter)
}
