package cmd

import (
	"context"
	"fmt"

	"github.com/intelops/qualitytrace/cli/cloud/runner"
	"github.com/intelops/qualitytrace/cli/cmdutil"
	"github.com/intelops/qualitytrace/cli/config"
	"github.com/intelops/qualitytrace/cli/formatters"
	"github.com/intelops/qualitytrace/cli/pkg/resourcemanager"
	"github.com/intelops/qualitytrace/cli/preprocessor"

	cliRunner "github.com/intelops/qualitytrace/cli/runner"
)

func RunMultipleFiles(ctx context.Context, httpClient *resourcemanager.HTTPClient, runParams *cmdutil.RunParameters, cliConfig *config.Config, runnerRegistry cliRunner.Registry, format string) (int, error) {
	if cliConfig.Jwt == "" {
		return cliRunner.ExitCodeGeneralError, fmt.Errorf("you should be authenticated to run multiple files, please run 'qualitytrace configure'")
	}

	variableSetPreprocessor := preprocessor.VariableSet(cmdutil.GetLogger())

	runGroup := runner.RunGroup(config.GetAPIClient(*cliConfig))
	formatter := formatters.MultipleRun[cliRunner.RunResult](func() string { return cliConfig.UI() }, true)

	orchestrator := runner.MultiFileOrchestrator(
		cmdutil.GetLogger(),
		runGroup,
		GetVariableSetClient(httpClient, variableSetPreprocessor),
		runnerRegistry,
		formatter,
	)

	return orchestrator.Run(ctx, runner.RunOptions{
		IDs:             runParams.IDs,
		ResourceName:    runParams.ResourceName,
		DefinitionFiles: runParams.DefinitionFiles,
		VarsID:          runParams.VarsID,
		SkipResultWait:  runParams.SkipResultWait,
		JUnitOuptutFile: runParams.JUnitOuptutFile,
		RequiredGates:   runParams.RequiredGates,
		RunGroupID:      runParams.RunGroupID,
	}, format)
}

func GetVariableSetClient(httpClient *resourcemanager.HTTPClient, preprocessor preprocessor.Preprocessor) resourcemanager.Client {
	variableSetClient := resourcemanager.NewClient(
		httpClient, cmdutil.GetLogger(),
		"variableset", "variablesets",
		resourcemanager.WithTableConfig(resourcemanager.TableConfig{
			Cells: []resourcemanager.TableCellConfig{
				{Header: "ID", Path: "spec.id"},
				{Header: "NAME", Path: "spec.name"},
				{Header: "DESCRIPTION", Path: "spec.description"},
			},
		}),
		resourcemanager.WithResourceType("VariableSet"),
		resourcemanager.WithApplyPreProcessor(preprocessor.Preprocess),
		resourcemanager.WithDeprecatedAlias("Environment"),
	)

	return variableSetClient
}
