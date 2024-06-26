package demo

import (
	"fmt"
	"testing"

	"github.com/intelops/qualitytrace/testing/cli-e2etest/environment"
	"github.com/intelops/qualitytrace/testing/cli-e2etest/helpers"
	"github.com/intelops/qualitytrace/testing/cli-e2etest/qualitytracecli"
	"github.com/intelops/qualitytrace/testing/cli-e2etest/testscenarios/types"
	"github.com/stretchr/testify/require"
)

func TestApplyDemo(t *testing.T) {
	// instantiate require with testing helper
	require := require.New(t)

	// setup isolated e2e environment
	env := environment.CreateAndStart(t)
	defer env.Close(t)

	cliConfig := env.GetCLIConfigPath(t)

	// Given I am a Tracetest CLI user
	// And I have my server recently created

	// When I try to set up a new demo
	// Then it should be applied with success
	newDemoPath := env.GetTestResourcePath(t, "new-demo")

	result := qualitytracecli.Exec(t, fmt.Sprintf("apply demo --file %s", newDemoPath), qualitytracecli.WithCLIConfig(cliConfig))
	helpers.RequireExitCodeEqual(t, result, 0)

	demo := helpers.UnmarshalYAML[types.DemoResource](t, result.StdOut)

	require.Equal("Demo", demo.Type)
	require.Equal("dev", demo.Spec.Name)
	require.Equal("otelstore", demo.Spec.Type)
	require.True(demo.Spec.Enabled)
	require.Equal("http://dev-cart:8082", demo.Spec.OTelStore.CartEndpoint)
	require.Equal("http://dev-checkout:8083", demo.Spec.OTelStore.CheckoutEndpoint)
	require.Equal("http://dev-frontend:9000", demo.Spec.OTelStore.FrontendEndpoint)
	require.Equal("http://dev-product:8081", demo.Spec.OTelStore.ProductCatalogEndpoint)

	// When I try to get the demo applied on the last step
	// Then it should return it
	command := fmt.Sprintf("get demo --id %s --output yaml", demo.Spec.Id)
	result = qualitytracecli.Exec(t, command, qualitytracecli.WithCLIConfig(cliConfig))
	helpers.RequireExitCodeEqual(t, result, 0)

	demo = helpers.UnmarshalYAML[types.DemoResource](t, result.StdOut)

	require.Equal("Demo", demo.Type)
	require.Equal("dev", demo.Spec.Name)
	require.Equal("otelstore", demo.Spec.Type)
	require.True(demo.Spec.Enabled)
	require.Equal("http://dev-cart:8082", demo.Spec.OTelStore.CartEndpoint)
	require.Equal("http://dev-checkout:8083", demo.Spec.OTelStore.CheckoutEndpoint)
	require.Equal("http://dev-frontend:9000", demo.Spec.OTelStore.FrontendEndpoint)
	require.Equal("http://dev-product:8081", demo.Spec.OTelStore.ProductCatalogEndpoint)

	// When I try to update the last demo
	// Then it should be applied with success
	updatedNewDemoPath := env.GetTestResourcePath(t, "updated-new-demo")
	helpers.Copy(updatedNewDemoPath+".tpl", updatedNewDemoPath)
	helpers.InjectIdIntoDemoFile(t, updatedNewDemoPath, demo.Spec.Id)

	result = qualitytracecli.Exec(t, fmt.Sprintf("apply demo --file %s", updatedNewDemoPath), qualitytracecli.WithCLIConfig(cliConfig))
	helpers.RequireExitCodeEqual(t, result, 0)

	updatedDemo := helpers.UnmarshalYAML[types.DemoResource](t, result.StdOut)
	require.Equal("Demo", updatedDemo.Type)
	require.Equal("dev-updated", updatedDemo.Spec.Name)
	require.Equal("otelstore", updatedDemo.Spec.Type)
	require.True(updatedDemo.Spec.Enabled)
	require.Equal("http://dev-updated-cart:8082", updatedDemo.Spec.OTelStore.CartEndpoint)
	require.Equal("http://dev-updated-checkout:8083", updatedDemo.Spec.OTelStore.CheckoutEndpoint)
	require.Equal("http://dev-updated-frontend:9000", updatedDemo.Spec.OTelStore.FrontendEndpoint)
	require.Equal("http://dev-updated-product:8081", updatedDemo.Spec.OTelStore.ProductCatalogEndpoint)

	// When I try to get the demo applied on the last step
	// Then it should return it
	command = fmt.Sprintf("get demo --id %s --output yaml", updatedDemo.Spec.Id)
	result = qualitytracecli.Exec(t, command, qualitytracecli.WithCLIConfig(cliConfig))
	helpers.RequireExitCodeEqual(t, result, 0)

	updatedDemo = helpers.UnmarshalYAML[types.DemoResource](t, result.StdOut)
	require.Equal("Demo", updatedDemo.Type)
	require.Equal("dev-updated", updatedDemo.Spec.Name)
	require.Equal("otelstore", updatedDemo.Spec.Type)
	require.True(updatedDemo.Spec.Enabled)
	require.Equal("http://dev-updated-cart:8082", updatedDemo.Spec.OTelStore.CartEndpoint)
	require.Equal("http://dev-updated-checkout:8083", updatedDemo.Spec.OTelStore.CheckoutEndpoint)
	require.Equal("http://dev-updated-frontend:9000", updatedDemo.Spec.OTelStore.FrontendEndpoint)
	require.Equal("http://dev-updated-product:8081", updatedDemo.Spec.OTelStore.ProductCatalogEndpoint)
}
