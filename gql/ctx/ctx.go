package ctx

import (
	"context"

	"github.com/uptrace/bun"
)

func NewContext(ctx context.Context, db *bun.DB) Context {
	return Context{
		Context: ctx,
		db:      db,
	}
}

type Context struct {
	context.Context

	db *bun.DB
}

func (ctx Context) DB() *bun.DB {
	return ctx.db
}
