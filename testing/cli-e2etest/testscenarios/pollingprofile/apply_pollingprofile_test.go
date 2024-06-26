package pollingprofile

import (
	"fmt"
	"testing"

	"github.com/intelops/qualitytrace/testing/cli-e2etest/environment"
	"github.com/intelops/qualitytrace/testing/cli-e2etest/helpers"
	"github.com/intelops/qualitytrace/testing/cli-e2etest/qualitytracecli"
	"github.com/intelops/qualitytrace/testing/cli-e2etest/testscenarios/types"
	"github.com/stretchr/testify/require"
)

func TestApplyPollingProfile(t *testing.T) {
	// instantiate require with testing helper
	require := require.New(t)

	// setup isolated e2e environment
	env := environment.CreateAndStart(t)
	defer env.Close(t)

	cliConfig := env.GetCLIConfigPath(t)

	// Given I am a Tracetest CLI user
	// And I have my server recently created

	// When I try to set up a new polling profile
	// Then it should be applied with success
	pollingProfilePath := env.GetTestResourcePath(t, "new-pollingprofile")

	result := qualitytracecli.Exec(t, fmt.Sprintf("apply pollingprofile --file %s", pollingProfilePath), qualitytracecli.WithCLIConfig(cliConfig))
	helpers.RequireExitCodeEqual(t, result, 0)

	// When I try to get a polling profile
	// Then it should return the polling profile applied on the last step
	result = qualitytracecli.Exec(t, "get pollingprofile --id current", qualitytracecli.WithCLIConfig(cliConfig))
	helpers.RequireExitCodeEqual(t, result, 0)

	pollingProfile := helpers.UnmarshalYAML[types.PollingProfileResource](t, result.StdOut)
	require.Equal("PollingProfile", pollingProfile.Type)
	require.Equal("current", pollingProfile.Spec.ID)
	require.Equal("current", pollingProfile.Spec.Name)
	require.True(pollingProfile.Spec.Default)
	require.Equal("periodic", pollingProfile.Spec.Strategy)
	require.Equal("50s", pollingProfile.Spec.Periodic.RetryDelay)
	require.Equal("10m", pollingProfile.Spec.Periodic.Timeout)
}
