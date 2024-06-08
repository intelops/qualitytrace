package datastore

import (
	"fmt"
	"testing"

	"github.com/intelops/qualitytrace/testing/cli-e2etest/environment"
	"github.com/intelops/qualitytrace/testing/cli-e2etest/helpers"
	"github.com/intelops/qualitytrace/testing/cli-e2etest/qualitytracecli"
	"github.com/intelops/qualitytrace/testing/cli-e2etest/testscenarios/types"
	"github.com/stretchr/testify/require"
)

func TestDeleteDatastore(t *testing.T) {
	// instantiate require with testing helper
	require := require.New(t)

	// setup isolated e2e environment
	env := environment.CreateAndStart(t)
	defer env.Close(t)

	cliConfig := env.GetCLIConfigPath(t)

	// Given I am a Tracetest CLI user
	// And I have my server recently created

	// When I try to set up a new datastore
	// Then it should be applied with success
	dataStorePath := env.GetEnvironmentResourcePath(t, "data-store")

	result := qualitytracecli.Exec(t, fmt.Sprintf("apply datastore --file %s", dataStorePath), qualitytracecli.WithCLIConfig(cliConfig))
	helpers.RequireExitCodeEqual(t, result, 0)

	// When I try to get a datastore
	// Then it should return the datastore applied on the last step
	result = qualitytracecli.Exec(t, "get datastore --id current", qualitytracecli.WithCLIConfig(cliConfig))
	helpers.RequireExitCodeEqual(t, result, 0)

	dataStore := helpers.UnmarshalYAML[types.DataStoreResource](t, result.StdOut)
	require.Equal("DataStore", dataStore.Type)
	require.Equal("current", dataStore.Spec.ID)
	require.True(dataStore.Spec.Default)

	// When I try to delete the datastore
	// Then it should return a error message, showing that we cannot delete a datastore
	result = qualitytracecli.Exec(t, "delete datastore --id current", qualitytracecli.WithCLIConfig(cliConfig))
	helpers.RequireExitCodeEqual(t, result, 1)
	require.Contains(result.StdErr, "resource DataStore does not support the action")
}
