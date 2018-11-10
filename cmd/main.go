package main

import (
	"fmt"
	"github.com/matthewjamesboyle/testEnsure/internal/system"
	"log"
	"os"
)

func main() {
	result, err := system.Explore(".")
	if err != nil {
		log.Fatalf("Error running program: " + err.Error())
	}
	if len(result) == 0 {
		log.Println("All files are tested!")
		os.Exit(0)
	} else {
		fmt.Println("FAILED: ")
		fmt.Println("Untested files: ")
		for _, untested := range result {
			fmt.Println(untested)
			os.Exit(-1)
		}
	}
}
