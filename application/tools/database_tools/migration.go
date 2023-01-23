package main

import (
	"context"
	"log"
	"os"

	atlas "ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/go-sql-driver/mysql"
	"mvc-model-sample/application/model/orm/ent/migrate"
)

func main() {
    ctx := context.Background()
    // Create a local migration directory able to understand Atlas migration file format for replay.
    dir, err := atlas.NewLocalDir("migration/")
    if err != nil {
        log.Fatalf("failed creating atlas migration directory: %v", err)
    }
    // Migrate diff options.
    opts := []schema.MigrateOption{
        schema.WithDir(dir),                         // provide migration directory
        schema.WithMigrationMode(schema.ModeInspect), // provide migration mode
        schema.WithDialect(dialect.MySQL),           // Ent dialect to use
        schema.WithFormatter(atlas.DefaultFormatter),
    }
    if len(os.Args) != 2 {
        log.Fatalln("migration name is required. Use: 'go run -mod=mod ent/migrate/main.go <name>'")
    }
    // Generate migrations using Atlas support for MySQL (note the Ent dialect option passed above).
    err = migrate.NamedDiff(ctx, "mysql://root:pass@localhost:3306/test", os.Args[1], opts...)
    if err != nil {
        log.Fatalf("failed generating migration file: %v", err)
    }
}
