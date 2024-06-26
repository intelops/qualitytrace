package integration_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/goccy/go-yaml"
	"github.com/intelops/qualitytrace/server/app"
	"github.com/intelops/qualitytrace/server/datastore"
	"github.com/intelops/qualitytrace/server/resourcemanager"
	"github.com/intelops/qualitytrace/server/test"
	"github.com/intelops/qualitytrace/server/testmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func getQualitytraceApp(options ...testmock.TestingAppOption) (*app.App, error) {
	qualitytraceApp, err := testmock.GetTestingApp(options...)
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		qualitytraceApp.Start()
		time.Sleep(1 * time.Second)
		wg.Done()
	}()

	wg.Wait()

	return qualitytraceApp, nil
}

func TestServerPrefix(t *testing.T) {
	_, err := getQualitytraceApp(
		testmock.WithServerPrefix("/qualitytrace"),
		testmock.WithHttpPort(8000),
	)
	require.NoError(t, err)

	expectedEndpoint := "http://localhost:8000/qualitytrace"
	tests := getTests(t, expectedEndpoint)
	assert.NotNil(t, tests)

	dataStores := getDatastores(t, expectedEndpoint)
	assert.NotNil(t, dataStores)
	assert.GreaterOrEqual(t, dataStores.Count, 1)
}

func getTests(t *testing.T, endpoint string) resourcemanager.ResourceList[test.Test] {
	url := fmt.Sprintf("%s/api/tests", endpoint)
	resp, err := http.Get(url)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	bodyJsonBytes, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	var tests resourcemanager.ResourceList[test.Test]
	err = json.Unmarshal(bodyJsonBytes, &tests)
	require.NoError(t, err)

	return tests
}

func getDatastores(t *testing.T, endpoint string) resourcemanager.ResourceList[datastore.DataStore] {
	url := fmt.Sprintf("%s/api/datastores", endpoint)
	resp, err := http.Get(url)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	bodyJsonBytes, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	var dataStores resourcemanager.ResourceList[datastore.DataStore]
	err = yaml.Unmarshal(bodyJsonBytes, &dataStores)
	require.NoError(t, err)

	return dataStores
}
