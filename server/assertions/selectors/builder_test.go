package selectors_test

import (
	"testing"

	"github.com/intelops/qualityTrace/server/assertions/selectors"
	"github.com/stretchr/testify/assert"
)

func TestSimpleSelectorBuilder(t *testing.T) {
	testCases := []struct {
		Name             string
		Expression       string
		ShouldSucceed    bool
		ExpectedSelector selectors.Selector
	}{
		{
			Name:          "Selector with single attribute",
			Expression:    "span[qualityTrace.span.type=\"http\"]",
			ShouldSucceed: true,
		},
		{
			Name:          "Selector with multiple attributes",
			Expression:    "span[service.name=\"Pokeshop\" qualityTrace.error=true]",
			ShouldSucceed: true,
		},
		{
			Name:          "Selector with child selector",
			Expression:    "span[qualityTrace.span.duration=0.5] span[qualityTrace.span.type contains \"http\"]",
			ShouldSucceed: true,
		},
		{
			Name:          "Selector with pseudo class filter",
			Expression:    "span[http.status_code=200]:nth_child(3)",
			ShouldSucceed: true,
		},
		{
			Name:          "Selector with invalid syntax",
			Expression:    "span.qualityTrace.span.type=\"http\"",
			ShouldSucceed: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			_, err := selectors.New(testCase.Expression)
			if testCase.ShouldSucceed {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
