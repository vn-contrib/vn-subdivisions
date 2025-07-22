package fixtures

import (
	"context"

	"github.com/uptrace/bun"
)

var Fixtures = &fixtures{}

type fixture func(ctx context.Context, db *bun.DB) error

type fixtures []fixture

func (fs *fixtures) Register(f fixture) {
	*fs = append(*fs, f)
}
