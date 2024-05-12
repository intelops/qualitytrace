package cmd

import (
	"context"
	"fmt"

	"github.com/intelops/qualityTrace/cli/cloud/runner"
	"github.com/intelops/qualityTrace/cli/config"
	"github.com/intelops/qualityTrace/cli/formatters"
	cliRunner "github.com/intelops/qualityTrace/cli/runner"
)

func Wait(ctx context.Context, cliConfig *config.Config, runGroupID, format string) (int, error) {
	rungroupWaiter := runner.RunGroup(config.GetAPIClient(*cliConfig))
	runGroup, err := rungroupWaiter.WaitForCompletion(ctx, runGroupID)
	if err != nil {
		return runner.ExitCodeGeneralError, err
	}

	formatter := formatters.MultipleRun[cliRunner.RunResult](func() string { return cliConfig.UI() }, true)
	runnerGetter := func(resource any) (formatters.Runner[cliRunner.RunResult], error) {
		return nil, nil
	}

	output := formatters.MultipleRunOutput[cliRunner.RunResult]{
		Runs:         []cliRunner.RunResult{},
		Resources:    []any{},
		RunGroup:     runGroup,
		RunnerGetter: runnerGetter,
		HasResults:   true,
	}

	fmt.Println(formatter.Format(output, formatters.Output(format)))

	exitCode := runner.ExitCodeSuccess
	if runGroup.GetStatus() == "failed" {
		exitCode = runner.ExitCodeTestNotPassed
	}

	return exitCode, nil
}
