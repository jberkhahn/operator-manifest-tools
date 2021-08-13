package main

import (
	"log"

	"github.com/operator-framework/operator-manifest-tools/cmd"
	"github.com/spf13/cobra/doc"
)

func main() {
	err := doc.GenMarkdownTree(cmd.Root(), "../../../docs")
	if err != nil {
		log.Fatal(err)
	}
}
