package testrunner

import (
	"testing"

	"github.com/intelops/qualitytrace/testing/cli-e2etest/environment"
	"github.com/intelops/qualitytrace/testing/cli-e2etest/helpers"
	"github.com/intelops/qualitytrace/testing/cli-e2etest/qualitytracecli"
	"github.com/stretchr/testify/require"
)

func TestDeleteTestRunner(t *testing.T) {
	// instantiate require with testing helper
	require := require.New(t)

	// setup isolated e2e environment
	env := environment.CreateAndStart(t)
	defer env.Close(t)

	cliConfig := env.GetCLIConfigPath(t)

	// Given I am a Tracetest CLI user
	// And I have my server recently created

	// When I try to delete the testrunner
	// Then it should return a error message, showing that we cannot delete a testrunner
	result := qualitytracecli.Exec(t, "delete testrunner --id current", qualitytracecli.WithCLIConfig(cliConfig))
	helpers.RequireExitCodeEqual(t, result, 1)
	require.Contains(result.StdErr, "resource TestRunner does not support the action")
}
