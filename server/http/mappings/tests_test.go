package mappings_test

import (
	"testing"

	"github.com/intelops/qualityTrace/server/http/mappings"
	"github.com/intelops/qualityTrace/server/openapi"
	"github.com/intelops/qualityTrace/server/test"
	"github.com/intelops/qualityTrace/server/traces"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_OpenApiToModel_Outputs(t *testing.T) {
	in := openapi.Test{
		Outputs: []openapi.TestOutput{
			{
				Name: "OUTPUT",
				SelectorParsed: openapi.Selector{
					Query: `span[name="root"]`,
				},
				Value: "attr:qualityTrace.selected_spans.count",
			},
		},
	}

	expected := test.Outputs{
		{
			Name:     "OUTPUT",
			Selector: `span[name="root"]`,
			Value:    "attr:qualityTrace.selected_spans.count",
		},
	}

	m := mappings.New(traces.NewConversionConfig(), nil)

	actual, err := m.In.Test(in)
	require.NoError(t, err)

	assert.Equal(t, expected, actual.Outputs)
}
