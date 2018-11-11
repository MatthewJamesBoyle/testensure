package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/matthewjamesboyle/testEnsure/internal/system"
	"log"
	"os"
)

type skippablepackages struct {
	Packages []string
}

func main() {
	var filesToSkip skippablepackages
	if _, err := toml.DecodeFile("./testensure.toml", &filesToSkip); err != nil {
		log.Fatalf(err.Error())
	}

	result, err := system.Explore(".")
	fmt.Println(filesToSkip)
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
