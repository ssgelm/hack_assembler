package code

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ssgelm/hack_assembler/symboltable"
)

type InstructionType int

const (
	AINST InstructionType = iota
	CINST
)

var compTable = map[string]string{
	"0":   "0101010",
	"1":   "0111111",
	"-1":  "0111010",
	"D":   "0001100",
	"A":   "0110000",
	"!D":  "0001101",
	"!A":  "0110001",
	"-D":  "0001111",
	"-A":  "0110011",
	"D+1": "0011111",
	"A+1": "0110111",
	"D-1": "0001110",
	"A-1": "0110010",
	"D+A": "0000010",
	"D-A": "0010011",
	"A-D": "0000111",
	"D&A": "0000000",
	"D|A": "0010101",
	"M":   "1110000",
	"!M":  "1110001",
	"-M":  "1110011",
	"M+1": "1110111",
	"M-1": "1110010",
	"D+M": "1000010",
	"D-M": "1010011",
	"M-D": "1000111",
	"D&M": "1000000",
	"D|M": "1010101",
}

var destTable = map[string]string{
	"M":   "001",
	"D":   "010",
	"MD":  "011",
	"A":   "100",
	"AM":  "101",
	"AD":  "110",
	"AMD": "111",
}

var jumpTable = map[string]string{
	"JGT": "001",
	"JEQ": "010",
	"JGE": "011",
	"JLT": "100",
	"JNE": "101",
	"JLE": "110",
	"JMP": "111",
}

type Instruction interface {
	InstructionType() InstructionType
}

type Ainst struct {
	InstType InstructionType
	Value    string
}

func (inst Ainst) InstructionType() InstructionType {
	return inst.InstType
}

func (inst Ainst) String() string {
	var num int
	val, err := strconv.Atoi(inst.Value)
	if err != nil {
		num = symboltable.FetchVariable(inst.Value)
	} else {
		num = val
	}
	binaryNum := strconv.FormatInt(int64(num), 2)
	return "0" + padLeft(binaryNum, "0", 15)
}

type Cinst struct {
	InstType InstructionType
	Comp     string
	Dest     string
	Jump     string
}

func (inst Cinst) InstructionType() InstructionType {
	return inst.InstType
}

func (inst Cinst) String() string {
	var comp string
	var dest string
	var jump string
	var exist bool

	if inst.Comp != "" {
		comp, exist = compTable[inst.Comp]
		if !exist {
			fmt.Printf("Error, %v is an invalid instruction", inst.Comp)
			os.Exit(1)
		}
	} else {
		comp = "0000000"
	}

	if inst.Dest != "" {
		dest, exist = destTable[inst.Dest]
		if !exist {
			fmt.Printf("Error, %v is an invalid dest", inst.Dest)
			os.Exit(1)
		}
	} else {
		dest = "000"
	}

	if inst.Jump != "" {
		jump, exist = jumpTable[inst.Jump]
		if !exist {
			fmt.Printf("Error, %v is an invalid jump", inst.Jump)
			os.Exit(1)
		}
	} else {
		jump = "000"
	}

	return "111" + comp + dest + jump
}

var Program []Instruction

func Init() {
	Program = make([]Instruction, 0)
}

func AddLine(inst Instruction) {
	Program = append(Program, inst)
}

func PrintProgram() {
	for _, line := range Program {
		fmt.Println(line)
	}
}

func WriteProgram(filename string) {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer f.Close()

	for _, line := range Program {
		f.WriteString(fmt.Sprintln(line))
	}
	f.Sync()
}

func padLeft(str, pad string, length int) string {
	for {
		if len(str) == length {
			return str[0:length]
		}
		str = pad + str
	}
}
