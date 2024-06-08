package testscenarios

import (
	"testing"

	"github.com/intelops/qualitytrace/testing/cli-e2etest/environment"
	"github.com/intelops/qualitytrace/testing/cli-e2etest/helpers"
	"github.com/intelops/qualitytrace/testing/cli-e2etest/qualitytracecli"
	"github.com/stretchr/testify/require"
)

func TestVersionCommand(t *testing.T) {
	// instantiate require with testing helper
	require := require.New(t)

	// setup test server environment
	env := environment.CreateAndStart(t)
	defer env.Close(t)

	cliConfig := env.GetCLIConfigPath(t)

	// Given I am a Tracetest CLI user
	// When I try to check the qualitytrace version
	// Then I should receive a version string with success
	result := qualitytracecli.Exec(t, "version", qualitytracecli.WithCLIConfig(cliConfig))

	helpers.RequireExitCodeEqual(t, result, 0)
	require.Greater(len(result.StdOut), 0)
}
