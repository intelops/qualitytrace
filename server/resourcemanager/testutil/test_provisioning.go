package testutil

import (
	"context"
	"testing"

	"github.com/gorilla/mux"
	"github.com/intelops/qualitytrace/server/resourcemanager"
	"github.com/intelops/qualitytrace/server/testmock"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"
)

const (
	OperationProvisioningSuccess          Operation = "ProvisioningSuccess"
	OperationProvisioningError            Operation = "ProvisioningError"
	OperationProvisioningTypeNotSupported Operation = "ProvisioningTypeNotSupported"
)

func testProvisioning(t *testing.T, rt ResourceTypeTest) {
	t.Run("Provisioning", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			db := testmock.CreateMigratedDatabase()
			defer db.Close()

			manager := rt.RegisterManagerFn(mux.NewRouter(), db)
			if rt.Prepare != nil {
				rt.Prepare(t, OperationProvisioningSuccess, manager)
			}

			yamlContents := contentTypeYAML.fromJSON(rt.SampleJSON)
			values := map[string]any{}
			err := yaml.Unmarshal([]byte(yamlContents), &values)
			require.NoError(t, err)

			err = manager.Provision(context.TODO(), values)
			require.NoError(t, err)
		})

		t.Run("UnacceptableType", func(t *testing.T) {
			db := testmock.CreateMigratedDatabase()
			defer db.Close()

			manager := rt.RegisterManagerFn(mux.NewRouter(), db)
			if rt.Prepare != nil {
				rt.Prepare(t, OperationProvisioningTypeNotSupported, manager)
			}

			values := map[string]any{
				"type": "ThisShuoldn'tBeAValidType",
			}

			err := manager.Provision(context.TODO(), values)
			require.ErrorIs(t, err, resourcemanager.ErrTypeNotSupported)
		})
	})
}
