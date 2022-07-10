package compiler

import (
	"fmt"

	"github.com/kyleconroy/sqlc/internal/config"
	"github.com/kyleconroy/sqlc/internal/engine/dolphin"
	"github.com/kyleconroy/sqlc/internal/engine/postgresql"
	"github.com/kyleconroy/sqlc/internal/engine/sqlite"
	"github.com/kyleconroy/sqlc/internal/opts"
	"github.com/kyleconroy/sqlc/internal/sql/catalog"
)

// Compiler translates SQL statements to the configured programming
// language and for the targeted database engine. It uses the sqlc config file
// to identify the out put programming language and the targeted database engine.
//
// It specifies the SQL parser which generates the schema catalog and the query statement
// results. The generated outputs are held within the compiler.
type Compiler struct {
	conf    config.SQL
	combo   config.CombinedSettings
	catalog *catalog.Catalog
	parser  Parser
	result  *Result
}

// New initializes a new compiler using the sqlc configuration file information.
// The returned compiler will have a new parser and a new catalog initialized.
// It panics if the database engine provided is not supported.
func New(conf config.SQL, combo config.CombinedSettings) *Compiler {
	c := &Compiler{conf: conf, combo: combo}
	switch conf.Engine {
	case config.EngineXLemon:
		c.parser = sqlite.NewParser()
		c.catalog = sqlite.NewCatalog()
	case config.EngineMySQL:
		c.parser = dolphin.NewParser()
		c.catalog = dolphin.NewCatalog()
	case config.EnginePostgreSQL:
		c.parser = postgresql.NewParser()
		c.catalog = postgresql.NewCatalog()
	default:
		panic(fmt.Sprintf("unknown engine: %s", conf.Engine))
	}
	return c
}

// Catalog returns the catalog used by the compiler
func (c *Compiler) Catalog() *catalog.Catalog {
	return c.catalog
}

// ParseCatalog finds all valid .sql schema files in the provided paths.
// It removes rollback statements and proceeds to parse the schema.
// Finally, it generates the catalog and updates the compiler catalog value.
//
// Errors found during schema file processing are accumulated and returned
// when ParseCatalog completes its execution.
func (c *Compiler) ParseCatalog(schema []string) error {
	return c.parseCatalog(schema)
}

// ParseQueries finds all valid .sql query files in the provided paths.
// It then proceeds to parse the queries in the query files and generates []*Query,
// which are added to the returned Result.
//
// Errors found during query file processing are accumulated and returned
// when ParseQueries completes its execution.
func (c *Compiler) ParseQueries(queries []string, o opts.Parser) error {
	r, err := c.parseQueries(o)
	if err != nil {
		return err
	}
	c.result = r
	return nil
}

// Result returns the result from the compiler
func (c *Compiler) Result() *Result {
	return c.result
}
