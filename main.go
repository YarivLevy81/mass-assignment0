package main

import (
	"fmt"
	"os"
	"strconv"
	"bufio"
)

// ErrIO represents I/O errors: cli, files
const ErrIO int = 1

func readFile(filename string, numbers *[]int){
	f, err := os.Open(filename)
	if err != nil {
        fmt.Printf("File %s doesn't exist", filename)
		os.Exit(ErrIO)
    }

	println(f)

	scanner := bufio.NewScanner(f)
	i := 0
	for scanner.Scan() {
		lineStr := scanner.Text()
		num, _ := strconv.Atoi(lineStr)
		if i <= 50 {
			fmt.Println(lineStr, num)
			i++
		}
		*numbers = append(*numbers, num)
	}

	fmt.Println("Done reading file")
}

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

	fmt.Printf("Processing %s with %d processors\n", filename, numProcessors) 

	var numbers []int = make([]int, 0, 1000)
	readFile(filename, &numbers)

	fmt.Println(numbers[0])
	fmt.Println(numbers[len(numbers)-1])
	
	Sort(&numbers)
}