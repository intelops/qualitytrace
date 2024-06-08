package executor_test

import (
	"context"
	"testing"

	"github.com/intelops/qualitytrace/server/assertions/comparator"
	"github.com/intelops/qualitytrace/server/executor"
	"github.com/intelops/qualitytrace/server/expression"
	"github.com/intelops/qualitytrace/server/pkg/id"
	"github.com/intelops/qualitytrace/server/pkg/maps"
	"github.com/intelops/qualitytrace/server/test"
	"github.com/intelops/qualitytrace/server/traces"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/trace"
)

func TestAssertion(t *testing.T) {

	spanID := id.NewRandGenerator().SpanID()
	cases := []struct {
		name              string
		testDef           test.Specs
		trace             traces.Trace
		expectedResult    maps.Ordered[test.SpanQuery, []test.AssertionResult]
		expectedAllPassed bool
	}{
		{
			name: "CanAssert",
			testDef: test.Specs{
				{
					Selector: test.SpanQuery(`span[service.name="Pokeshop"]`),
					Assertions: []test.Assertion{
						`attr:qualitytrace.span.duration = 2000ns`,
					},
				},
			},
			trace: traces.Trace{
				RootSpan: traces.Span{
					ID: spanID,
					Attributes: traces.NewAttributes(map[string]string{
						"service.name":               "Pokeshop",
						"qualitytrace.span.duration": "2000",
					}),
				},
			},
			expectedAllPassed: true,
			expectedResult: (maps.Ordered[test.SpanQuery, []test.AssertionResult]{}).MustAdd(`span[service.name="Pokeshop"]`, []test.AssertionResult{
				{
					Assertion: `attr:qualitytrace.span.duration = 2000ns`,
					Results: []test.SpanAssertionResult{
						{
							SpanID:        &spanID,
							ObservedValue: "2us",
							CompareErr:    nil,
						},
					},
				},
			}),
		},
		{
			name: "CanAssertOnSpanMatchCount",
			testDef: test.Specs{
				{
					Selector: test.SpanQuery(`span[service.name="Pokeshop"]`),
					Assertions: []test.Assertion{
						`attr:qualitytrace.selected_spans.count = 1`,
					},
				},
				{
					Selector: test.SpanQuery(`span[service.name="NotExists"]`),
					Assertions: []test.Assertion{
						`attr:qualitytrace.selected_spans.count = 0`,
					},
				},
			},
			trace: traces.Trace{
				RootSpan: traces.Span{
					ID: spanID,
					Attributes: traces.NewAttributes(map[string]string{
						"service.name":               "Pokeshop",
						"qualitytrace.span.duration": "2000",
					}),
				},
			},
			expectedAllPassed: true,
			expectedResult: (maps.Ordered[test.SpanQuery, []test.AssertionResult]{}).MustAdd(`span[service.name="Pokeshop"]`, []test.AssertionResult{
				{
					Assertion: `attr:qualitytrace.selected_spans.count = 1`,
					Results: []test.SpanAssertionResult{
						{
							SpanID:        &spanID,
							ObservedValue: "1",
							CompareErr:    nil,
						},
					},
				},
			}).MustAdd(`span[service.name="NotExists"]`, []test.AssertionResult{
				{
					Assertion: `attr:qualitytrace.selected_spans.count = 0`,
					Results: []test.SpanAssertionResult{
						{
							SpanID:        nil,
							ObservedValue: "0",
							CompareErr:    nil,
						},
					},
				},
			}),
		},
		// https://github.com/intelops/qualitytrace/issues/617
		{
			name: "ContainsWithJSON",
			testDef: test.Specs{
				{
					Selector: test.SpanQuery(`span[service.name="Pokeshop"]`),
					Assertions: []test.Assertion{
						`attr:http.response.body contains 52`,
						`attr:qualitytrace.span.duration <= 21ms`,
					},
				},
			},
			trace: traces.Trace{
				RootSpan: traces.Span{
					ID: spanID,
					Attributes: traces.NewAttributes(map[string]string{
						"service.name":               "Pokeshop",
						"http.response.body":         `{"id":52}`,
						"qualitytrace.span.duration": "21000000",
					}),
				},
			},
			expectedAllPassed: true,
			expectedResult: (maps.Ordered[test.SpanQuery, []test.AssertionResult]{}).MustAdd(`span[service.name="Pokeshop"]`, []test.AssertionResult{
				{
					Assertion: `attr:http.response.body contains 52`,
					Results: []test.SpanAssertionResult{
						{
							SpanID:        &spanID,
							ObservedValue: `{"id":52}`,
							CompareErr:    nil,
						},
					},
				},
				{
					Assertion: `attr:qualitytrace.span.duration <= 21ms`,
					Results: []test.SpanAssertionResult{
						{
							SpanID:        &spanID,
							ObservedValue: "21ms",
							CompareErr:    nil,
						},
					},
				},
			}),
		},
		// https://github.com/intelops/qualitytrace/issues/1203
		{
			name: "DurationComparison",
			testDef: test.Specs{
				{
					Selector: test.SpanQuery(`span[service.name="Pokeshop"]`),
					Assertions: []test.Assertion{
						`attr:qualitytrace.span.duration <= 25ms`,
					},
				},
			},
			trace: traces.Trace{
				RootSpan: traces.Span{
					ID: spanID,
					Attributes: traces.NewAttributes(map[string]string{
						"service.name":               "Pokeshop",
						"http.response.body":         `{"id":52}`,
						"qualitytrace.span.duration": "25187564", // 25ms
					}),
				},
			},
			expectedAllPassed: true,
			expectedResult: (maps.Ordered[test.SpanQuery, []test.AssertionResult]{}).MustAdd(`span[service.name="Pokeshop"]`, []test.AssertionResult{
				{
					Assertion: `attr:qualitytrace.span.duration <= 25ms`,
					Results: []test.SpanAssertionResult{
						{
							SpanID:        &spanID,
							ObservedValue: "25ms",
							CompareErr:    nil,
						},
					},
				},
			}),
		},
		// https://github.com/intelops/qualitytrace/issues/1421
		{
			name: "FailedAssertionsConvertDurationFieldsIntoDurationFormat",
			testDef: test.Specs{
				{
					Selector: test.SpanQuery(`span[service.name="Pokeshop"]`),
					Assertions: []test.Assertion{
						`attr:qualitytrace.span.duration <= 25ms`,
					},
				},
			},
			trace: traces.Trace{
				RootSpan: traces.Span{
					ID: spanID,
					Attributes: traces.NewAttributes(map[string]string{
						"service.name":               "Pokeshop",
						"http.response.body":         `{"id":52}`,
						"qualitytrace.span.duration": "35000000", // 35ms
					}),
				},
			},
			expectedAllPassed: false,
			expectedResult: (maps.Ordered[test.SpanQuery, []test.AssertionResult]{}).MustAdd(`span[service.name="Pokeshop"]`, []test.AssertionResult{
				{
					Assertion: `attr:qualitytrace.span.duration <= 25ms`,
					Results: []test.SpanAssertionResult{
						{
							SpanID:        &spanID,
							ObservedValue: "35ms",
							CompareErr:    comparator.ErrNoMatch,
						},
					},
				},
			}),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			cl := c

			executor := executor.NewAssertionExecutor(trace.NewNoopTracerProvider().Tracer("tracer"))
			actual, allPassed := executor.Assert(context.Background(), cl.testDef, cl.trace, []expression.DataStore{})

			assert.Equal(t, cl.expectedAllPassed, allPassed)

			cl.expectedResult.ForEach(func(expectedSel test.SpanQuery, expectedAssertionResults []test.AssertionResult) error {
				actualAssertionResults := actual.Get(expectedSel)
				assert.NotEmpty(t, actualAssertionResults, `expected selector "%s" not found`, expectedSel)
				for i := 0; i < len(expectedAssertionResults); i++ {
					expectedAR := expectedAssertionResults[i]
					actualAR := actualAssertionResults[i]

					assert.Equal(t, expectedAR.Assertion, actualAR.Assertion)
					require.Len(t, actualAR.Results, len(expectedAR.Results))

					for i, expectedSpanRes := range expectedAR.Results {
						actualSpanRes := actualAR.Results[i]
						assert.Equal(t, expectedSpanRes.ObservedValue, actualSpanRes.ObservedValue)
						assert.Equal(t, expectedSpanRes.CompareErr, actualSpanRes.CompareErr)
					}
				}

				return nil
			})

		})
	}
}
