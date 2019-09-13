package main

import (
	"fmt"
	"log"

	"aqwari.net/xml/xsdgen"
)

func main() {
	var cfg xsdgen.Config
	cfg.Option(xsdgen.PackageName("postal"))

	out, err := cfg.GenSource("ocil-2.0.xsd")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", out)

	// Output: package postal
	//
	// // May be no more than 10 items long
	// type Zipcode string
}


