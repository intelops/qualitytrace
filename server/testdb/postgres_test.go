package testdb_test

import (
	"context"
	"testing"
	"time"

	"github.com/intelops/qualitytrace/agent/workers/trigger"
	"github.com/intelops/qualitytrace/server/model"
	"github.com/intelops/qualitytrace/server/pkg/id"
	"github.com/intelops/qualitytrace/server/test"
	"github.com/intelops/qualitytrace/server/testmock"
)

func getDB() (model.Repository, func()) {
	db := testmock.GetTestingDatabase()

	clean := func() {
		defer db.Close()
		err := db.Drop()
		if err != nil {
			panic(err)
		}
	}

	return db, clean
}

func createTestWithName(t *testing.T, db test.Repository, name string) test.Test {
	t.Helper()
	test := test.Test{
		Name:        name,
		Description: "description",
		Trigger: trigger.Trigger{
			Type: trigger.TriggerTypeHTTP,
			HTTP: &trigger.HTTPRequest{
				URL: "http://localhost:3030/hello-instrumented",
			},
		},
	}

	updated, err := db.Create(context.TODO(), test)
	if err != nil {
		panic(err)
	}
	return updated
}

func createTest(t *testing.T, db test.Repository) test.Test {
	return createTestWithName(t, db, "first test")
}

func createRun(t *testing.T, db test.RunRepository, testObj test.Test) test.Run {
	t.Helper()
	run := test.Run{
		TraceID:   id.NewRandGenerator().TraceID(),
		SpanID:    id.NewRandGenerator().SpanID(),
		CreatedAt: time.Now(),
	}
	updated, err := db.CreateRun(context.TODO(), testObj, run)
	if err != nil {
		panic(err)
	}

	return updated
}
