package main

import (
	"fmt"
	"log"

	"github.com/fachrunwira/go-query-builder/builder"
	"github.com/fachrunwira/go-query-builder/examples"
)

func main() {
	db, err := examples.Init()
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	result, err := builder.Make(db).Table("users").Limit(1).Get()
	if err != nil {
		log.Fatalln(err)
	}

	for _, m := range result {
		for col, val := range m {
			fmt.Printf("%s: %v\n", col, val)
		}
	}

	// fmt.Println(result)
}
