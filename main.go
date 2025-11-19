package main

import (
	"fmt"
	"log"

	"github.com/fachrunwira/go-query-builder/lib"
)

func main() {
	db, err := lib.Init()
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	fmt.Println("done")
}
