package config_test

import (
	"os"
	"testing"

	"github.com/intelops/qualityTrace/server/testmock"
)

func TestMain(m *testing.M) {
	testmock.StartTestEnvironment()

	exitVal := m.Run()

	testmock.StopTestEnvironment()

	os.Exit(exitVal)
}
