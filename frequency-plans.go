package main

import (
	"flag"
	"log"

	"github.com/TheThingsNetwork/lorawan-frequency-plans/internal/docs"
	"github.com/TheThingsNetwork/lorawan-frequency-plans/internal/schema"
	"github.com/TheThingsNetwork/lorawan-frequency-plans/internal/validate"
)

var (
	generateDocs   = flag.Bool("docs", false, "Generate docs for the frequency-plans.")
	generateSchema = flag.Bool("schema", false, "Generate the `schema.json` file.")
)

func main() {
	flag.Parse()

	if err := validate.Validate(); err != nil {
		log.Fatal(err)
	}

	if *generateDocs {
		// TODO: update doc generation to support end-device and gateway folders + new structure
		if err := docs.Generate("./frequency-plans.yml", "./docs"); err != nil {
			log.Fatal(err)
		}
	}

	if *generateSchema {
		if err := schema.Generate(); err != nil {
			log.Fatal(err)
		}
	}
}
