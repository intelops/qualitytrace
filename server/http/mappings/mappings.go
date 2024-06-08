package mappings

import (
	"github.com/intelops/qualitytrace/server/assertions/comparator"
	"github.com/intelops/qualitytrace/server/traces"
)

type Mappings struct {
	In  Model
	Out OpenAPI
}

func New(tcc traces.ConversionConfig, cr comparator.Registry) Mappings {
	return Mappings{
		In: Model{
			comparators:           cr,
			traceConversionConfig: tcc,
		},
		Out: OpenAPI{
			traceConversionConfig: tcc,
		},
	}
}
