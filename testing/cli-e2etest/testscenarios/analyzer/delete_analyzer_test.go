package analyzer

import (
	"testing"

	"github.com/intelops/qualityTrace/testing/cli-e2etest/environment"
	"github.com/intelops/qualityTrace/testing/cli-e2etest/helpers"
	"github.com/intelops/qualityTrace/testing/cli-e2etest/qualityTracecli"
	"github.com/stretchr/testify/require"
)

func TestDeleteAnalyzer(t *testing.T) {
	// instantiate require with testing helper
	require := require.New(t)

	// setup isolated e2e environment
	env := environment.CreateAndStart(t)
	defer env.Close(t)

	cliConfig := env.GetCLIConfigPath(t)

	// Given I am a Tracetest CLI user
	// And I have my server recently created

	// When I try to delete the analyzer
	// Then it should return a error message, showing that we cannot delete a analyzer
	result := qualityTracecli.Exec(t, "delete analyzer --id current", qualityTracecli.WithCLIConfig(cliConfig))
	helpers.RequireExitCodeEqual(t, result, 1)
	require.Contains(result.StdErr, "resource Analyzer does not support the action")
}
