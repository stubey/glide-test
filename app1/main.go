package main

import (
	"github.com/stubey/glide-test/pkg1"
	"github.com/stubey/glide-test/pkg2"
	"log"
)

func main() {
	log.Printf("hello from %s", pkg1.Name)
	log.Printf("hellp from %s", pkg2.Name)
}
