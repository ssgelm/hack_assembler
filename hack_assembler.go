package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/ssgelm/hack_assembler/code"
	"github.com/ssgelm/hack_assembler/parser"
	"github.com/ssgelm/hack_assembler/symboltable"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %v [inputfile]\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Input file is missing.")
		os.Exit(1)
	}
	asm_filename := args[0]

	file, err := os.Open(asm_filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer file.Close()

	symboltable.Init()
	code.Init()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parser.Parse(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}

	outfile := asm_filename[0:len(asm_filename)-3] + "hack"
	code.WriteProgram(outfile)
}
