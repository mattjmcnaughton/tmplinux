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

	mockExecutor.AssertKeywordIncludedInCommand(t, "run")
	mockReporter.AssertNotCalled(t)
}

func TestDockerEngineCommandFailed(t *testing.T) {
	mockExecutor := executor.CreateFailureMockExecutor()
	mockReporter := reporter.MockReporter{}

	testDockerEngine := engine.NewCustomDockerEngine(&mockExecutor, &mockReporter)
	testDockerEngine.Start()

	mockReporter.AssertCalled(t)
}

func TestDockerEngineSSHCommandSuccess(t *testing.T) {
	mockExecutor := executor.CreateSuccessMockExecutor()
	mockReporter := reporter.MockReporter{}

	testDockerEngine := engine.NewCustomDockerEngine(&mockExecutor, &mockReporter)
	testDockerEngine.SSH()

	mockExecutor.AssertKeywordIncludedInCommand(t, "exec")
	mockReporter.AssertNotCalled(t)
}

func TestDockerEngineStopCommandSuccess(t *testing.T) {
	mockExecutor := executor.CreateSuccessMockExecutor()
	mockReporter := reporter.MockReporter{}

	testDockerEngine := engine.NewCustomDockerEngine(&mockExecutor, &mockReporter)
	testDockerEngine.Stop()

	mockExecutor.AssertKeywordIncludedInCommand(t, "kill")
	mockReporter.AssertNotCalled(t)
}

func TestDockerEngineRmCommandSuccess(t *testing.T) {
	mockExecutor := executor.CreateSuccessMockExecutor()
	mockReporter := reporter.MockReporter{}

	testDockerEngine := engine.NewCustomDockerEngine(&mockExecutor, &mockReporter)
	testDockerEngine.Rm()

	mockExecutor.AssertKeywordIncludedInCommand(t, "rm")
	mockReporter.AssertNotCalled(t)
}
