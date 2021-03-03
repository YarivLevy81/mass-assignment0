package main

import (
	"fmt"
	"os"
	"strconv"
)

// ErrIO represents I/O errors: cli, files
const ErrIO int = 1

func main() {
	
	if (len(os.Args) != 3) {
		fmt.Println("USAGE: main $num_processors $filename")
		os.Exit(ErrIO)
	}

	filename := os.Args[2]
	numProcessors, err := strconv.Atoi(os.Args[1]) 
	if err != nil {
        fmt.Println(err)
        os.Exit(ErrIO)
    }	

	fmt.Printf("Processing %s with %d processors", filename, numProcessors) 

}