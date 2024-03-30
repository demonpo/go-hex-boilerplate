package main

import (
	"fmt"
	"io"
	"os"

	"ariga.io/atlas-provider-gorm/gormschema"
	entitiesInfra "goHexBoilerplate/src/modules/user/infra/entities"
)

func main() {
	stmts, err := gormschema.New("postgres").Load(&entitiesInfra.User{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
		os.Exit(1)
	}
	io.WriteString(os.Stdout, stmts)
}
