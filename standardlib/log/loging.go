package main

import (
	"fmt"
	"log"
)

func init() {
	fmt.Println(8)
	log.SetFlags(15)
}
func main() {
	log.SetPrefix("TRACE: ")
	log.Println("Joss")

	log.Printf("Set %v printf", 100)

	// log.Fatalln("fatal")
	// log.Panicln("Panic!")
}