package executor

import (
	"context"

	"github.com/intelops/qualityTrace/server/executor/testrunner"
	"github.com/intelops/qualityTrace/server/pkg/pipeline"
	"github.com/intelops/qualityTrace/server/test"
	"github.com/intelops/qualityTrace/server/testsuite"
	"github.com/intelops/qualityTrace/server/variableset"
)

type TestSuitesPipeline struct {
	*pipeline.Pipeline[Job]
	runs testSuiteRunRepo
}

type testSuiteRunRepo interface {
	CreateRun(context.Context, testsuite.TestSuiteRun) (testsuite.TestSuiteRun, error)
}

func NewTestSuitePipeline(
	pipeline *pipeline.Pipeline[Job],
	runs testSuiteRunRepo,
) *TestSuitesPipeline {
	return &TestSuitesPipeline{
		Pipeline: pipeline,
		runs:     runs,
	}
}

func (p *TestSuitesPipeline) Run(ctx context.Context, tran testsuite.TestSuite, metadata test.RunMetadata, variableSet variableset.VariableSet, requiredGates *[]testrunner.RequiredGate) testsuite.TestSuiteRun {
	tranRun := tran.NewRun()
	tranRun.Metadata = metadata
	tranRun.VariableSet = variableSet
	tranRun.RequiredGates = requiredGates

	tranRun, _ = p.runs.CreateRun(ctx, tranRun)

	job := NewJob()
	job.TestSuite = tran
	job.TestSuiteRun = tranRun

	p.Pipeline.Begin(ctx, job)

	return tranRun
}
