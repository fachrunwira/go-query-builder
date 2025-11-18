package main

import (
	"fmt"
	"log"

	"github.com/fachrunwira/go-query-builder/lib"
	"github.com/fachrunwira/go-query-builder/lib/builder"
)

func main() {
	db, err := lib.Init()
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	err = builder.Make(db).Table("users").WhereRaw("id in (?,?)", 10001, 10000).Delete().Save()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("done")
}
