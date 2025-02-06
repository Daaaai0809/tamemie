package main

import (
	"fmt"
	"os"

	"github.com/Daaaai0809/tamemie/config"
	"github.com/Daaaai0809/tamemie/repositories/mysql"
	"gorm.io/gen"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./gen/query",
		Mode:    gen.WithQueryInterface,
	})

	c, err := config.NewConfig()
	if err != nil {
		fmt.Errorf("failed to load config: %v", err)
		os.Exit(1)
	}

	db, err := repository.Conn(c, "localhost")
	if err != nil {
		fmt.Errorf("failed to connect to database: %v", err)
		os.Exit(1)
	}

	g.UseDB(db)

	all := g.GenerateAllTable()

	g.ApplyBasic(all...)

	g.Execute()
}
