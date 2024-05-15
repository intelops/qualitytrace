package config

import "os"

type EnvironmentVars struct {
	EnableCLIDebug   bool
	TracetestCommand string
	TestEnvironment  string
}

var instance *EnvironmentVars

func GetConfigAsEnvVars() *EnvironmentVars {
	if instance != nil {
		return instance
	}

	enableCLIDebug := (os.Getenv("ENABLE_CLI_DEBUG") == "true")

	qualitytraceCommand := os.Getenv("TRACETEST_CLI")
	if qualitytraceCommand == "" {
		qualitytraceCommand = "qualitytrace"
	}

	testEnvironment := os.Getenv("TEST_ENVIRONMENT")
	if testEnvironment == "" {
		testEnvironment = "jaeger"
	}

	return &EnvironmentVars{
		EnableCLIDebug:   enableCLIDebug,
		TracetestCommand: qualitytraceCommand,
		TestEnvironment:  testEnvironment,
	}
}
