package preprocessor

import (
	"context"

	"github.com/intelops/qualityTrace/cli/pkg/fileutil"
)

type Preprocessor interface {
	Preprocess(ctx context.Context, input fileutil.File) (fileutil.File, error)
}
