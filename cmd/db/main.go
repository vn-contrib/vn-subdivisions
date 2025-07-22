package main

import (
	"context"
	"os"
	"strings"

	"github.com/uptrace/bun/migrate"
	"github.com/urfave/cli/v3"
	"github.com/vn-contrib/vn-subdivisions/cmd/db/migrations"
	"github.com/vn-contrib/vn-subdivisions/db"

	_ "github.com/errybase/go-dotenv/autoload"
)

func main() {
	db := db.NewDB()
	defer db.Close()

	migrator := migrate.NewMigrator(db, migrations.Migrations)

	cmd := &cli.Command{
		Commands: []*cli.Command{
			{
				Name:    "generate",
				Aliases: []string{"g"},
				Usage:   "Generate .sql migration",
				Action: func(ctx context.Context, c *cli.Command) error {
					name := strings.Join(c.Args().Slice(), "_")
					_, err := migrator.CreateSQLMigrations(ctx, name)
					return err
				},
			},
			{
				Name:  "init",
				Usage: "create migration tables",
				Action: func(ctx context.Context, c *cli.Command) error {
					return migrator.Init(ctx)
				},
			},
			{
				Name:    "migrate",
				Aliases: []string{"m"},
				Usage:   "migrate database",
				Action: func(ctx context.Context, c *cli.Command) error {
					_, err := migrator.Migrate(ctx)
					return err
				},
			},
			{
				Name:    "rollback",
				Aliases: []string{"rb"},
				Usage:   "rollback last migration",
				Action: func(ctx context.Context, c *cli.Command) error {
					_, err := migrator.Rollback(ctx)
					return err
				},
			},
		},
	}

	cmd.Run(context.Background(), os.Args)
}
