package variableset

import (
	"fmt"
	"testing"

	"github.com/intelops/qualityTrace/testing/cli-e2etest/environment"
	"github.com/intelops/qualityTrace/testing/cli-e2etest/helpers"
	"github.com/intelops/qualityTrace/testing/cli-e2etest/qualityTracecli"
	"github.com/intelops/qualityTrace/testing/cli-e2etest/testscenarios/types"
	"github.com/stretchr/testify/require"
)

func TestDeleteVariableSet(t *testing.T) {
	// instantiate require with testing helper
	require := require.New(t)

	// setup isolated e2e environment
	env := environment.CreateAndStart(t)
	defer env.Close(t)

	cliConfig := env.GetCLIConfigPath(t)

	// Given I am a Tracetest CLI user
	// And I have my server recently created

	// When I try to delete a variable set that don't exist
	// Then it should return an error and say that this resource does not exist
	result := qualityTracecli.Exec(t, "delete variableset --id .env", qualityTracecli.WithCLIConfig(cliConfig))
	helpers.RequireExitCodeEqual(t, result, 1)
	require.Contains(result.StdErr, "Resource variableset with ID .env not found")

	// When I try to set up a new environment
	// Then it should be applied with success
	newEnvironmentPath := env.GetTestResourcePath(t, "new-varSet")

	result = qualityTracecli.Exec(t, fmt.Sprintf("apply variableset --file %s", newEnvironmentPath), qualityTracecli.WithCLIConfig(cliConfig))
	helpers.RequireExitCodeEqual(t, result, 0)

	environmentVars := helpers.UnmarshalYAML[types.VariableSetResource](t, result.StdOut)
	require.Equal("VariableSet", environmentVars.Type)
	require.Equal(".env", environmentVars.Spec.ID)

	// When I try to delete the environment
	// Then it should delete with success
	result = qualityTracecli.Exec(t, "delete variableset --id .env", qualityTracecli.WithCLIConfig(cliConfig))
	helpers.RequireExitCodeEqual(t, result, 0)
	require.Contains(result.StdOut, "✔ Variableset successfully deleted")

	// When I try to get an environment again
	// Then it should return a message saying that the environment was not found
	result = qualityTracecli.Exec(t, "get variableset --id .env --output yaml", qualityTracecli.WithCLIConfig(cliConfig))
	helpers.RequireExitCodeEqual(t, result, 0)
	require.Contains(result.StdOut, "Resource variableset with ID .env not found")
}
