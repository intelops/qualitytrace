package model_test

import (
	"fmt"
	"testing"

	"github.com/intelops/qualityTrace/server/model"
	"github.com/intelops/qualityTrace/server/pkg/id"
	"github.com/stretchr/testify/assert"
)

func TestTestRunEvent_ResourceID(t *testing.T) {
	testID := id.NewRandGenerator().ID()
	runID := 1

	event := model.TestRunEvent{TestID: testID, RunID: runID}

	assert.Equal(t, event.ResourceID(), fmt.Sprintf("test/%s/run/%d/event", testID, runID))
}
