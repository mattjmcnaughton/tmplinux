package engine_test

import (
	"strings"
	"testing"

	"github.com/mattjmcnaughton/tmplinux/pkg/engine"
	"github.com/mattjmcnaughton/tmplinux/pkg/executor"
	"github.com/mattjmcnaughton/tmplinux/pkg/reporter"
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

func assertKeywordIncludedInCommand(t *testing.T, mockExecutor executor.MockShellExecutor, keyword string) {
	execCmd := mockExecutor.GetExecutedCommand()

	if !strings.Contains(execCmd, keyword) {
		t.Fatalf("%s should include the keyword %s", execCmd, keyword)
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
