package model

import "github.com/uptrace/bun"

type Area struct {
	bun.BaseModel `bun:"table:subdivisions"`

	ID       int64 `bun:",nullzero"`
	Name     string
	Unit     string
	Level    int8
	GsoID    string `bun:",nullzero"`
	ParentID int64  `bun:",nullzero"`

	ParentGsoID string `bun:"-"`
}
