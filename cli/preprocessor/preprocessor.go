package preprocessor

import (
	"context"

	"github.com/intelops/qualitytrace/cli/pkg/fileutil"
)

type Preprocessor interface {
	Preprocess(ctx context.Context, input fileutil.File) (fileutil.File, error)
}
