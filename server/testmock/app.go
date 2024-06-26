package testmock

import (
	"github.com/intelops/qualitytrace/server/app"
	"github.com/intelops/qualitytrace/server/config"
)

type TestingAppOption func(*config.AppConfig)

func WithServerPrefix(prefix string) TestingAppOption {
	return func(cfg *config.AppConfig) {
		cfg.Set("server.pathPrefix", prefix)
	}
}

func WithHttpPort(port int) TestingAppOption {
	return func(cfg *config.AppConfig) {
		cfg.Set("server.httpPort", port)
	}
}

func GetTestingApp(options ...TestingAppOption) (*app.App, error) {
	cfg, _ := config.New()
	for _, option := range options {
		option(cfg)
	}

	ConfigureDB(cfg)

	return app.New(cfg)
}

func ConfigureDB(cfg *config.AppConfig) {
	db := getTestDatabaseEnvironment()

	cfg.Set("postgres.host", db.container.Host)
	cfg.Set("postgres.user", "qualitytrace")
	cfg.Set("postgres.password", "qualitytrace")
	cfg.Set("postgres.dbname", "postgres")
	cfg.Set("postgres.port", db.container.DefaultPort())
}
