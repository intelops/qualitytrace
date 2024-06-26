package qualitytracecli

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/intelops/qualitytrace/testing/cli-e2etest/command"
	"github.com/intelops/qualitytrace/testing/cli-e2etest/config"
	"golang.org/x/exp/slices"

	"github.com/intelops/qualitytrace/cli/cmd"
	"github.com/stretchr/testify/require"
)

type ExecOption func(*executionState)

type executionState struct {
	cliConfigFile string
}

func Exec(t *testing.T, qualitytraceSubCommand string, options ...ExecOption) *command.ExecResult {
	state := &executionState{}
	for _, option := range options {
		option(state)
	}

	if state.cliConfigFile != "" {
		// append config at the start of the command
		qualitytraceSubCommand = fmt.Sprintf("--config %s %s", state.cliConfigFile, qualitytraceSubCommand)
	}

	qualitytraceCommand := config.GetConfigAsEnvVars().TracetestCommand
	qualitytraceSubCommands := strings.Split(qualitytraceSubCommand, " ")

	if config.GetConfigAsEnvVars().EnableCLIDebug {
		return runTracetestAsInternalCommand(t, qualitytraceCommand, qualitytraceSubCommands)
	}

	result, err := command.Exec(qualitytraceCommand, qualitytraceSubCommands...)
	require.NoError(t, err)

	return result
}

func WithCLIConfig(cliConfig string) ExecOption {
	return func(es *executionState) {
		es.cliConfigFile = cliConfig
	}
}

func runTracetestAsInternalCommand(t *testing.T, qualitytraceCommand string, qualitytraceSubCommands []string) *command.ExecResult {
	// This code calls the CLI as a library to enable Go debugger to step into CLI statements and help a dev to debug CLI problems found on CLI tests
	//, but emulates this call as an executable call intercepting data sent to stdout, stderr and part of the os.Exit commands

	// keep backup of the real stdout
	stdoutBackup := os.Stdout
	stdoutRead, stdoutWriter, _ := os.Pipe()
	os.Stdout = stdoutWriter

	// keep backup of the real stderr
	stderrBackup := os.Stderr
	stderrRead, stderrWriter, _ := os.Pipe()
	os.Stderr = stderrWriter

	argsBackup := os.Args
	os.Args = slices.Insert(qualitytraceSubCommands, 0, qualitytraceCommand)

	exitCode := 0
	cmd.RegisterCLIExitInterceptor(func(i int) {
		exitCode = i
	})

	cmd.Execute()

	os.Args = argsBackup

	stdoutChannel := make(chan string)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, stdoutRead)
		stdoutChannel <- buf.String()
	}()

	// back to normal state
	stdoutWriter.Close()
	os.Stdout = stdoutBackup // restoring the real stdout

	stderrChannel := make(chan string)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, stderrRead)
		stderrChannel <- buf.String()
	}()

	// back to normal state
	stderrWriter.Close()
	os.Stderr = stderrBackup // restoring the real stderr

	return &command.ExecResult{
		CommandExecuted: fmt.Sprintf("%s %s", qualitytraceCommand, strings.Join(qualitytraceSubCommands, " ")),
		StdOut:          <-stdoutChannel,
		StdErr:          <-stderrChannel,
		ExitCode:        exitCode,
	}
}
