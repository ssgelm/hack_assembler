package parser

import (
	"strings"

	"github.com/ssgelm/hack_assembler/code"
	"github.com/ssgelm/hack_assembler/symboltable"
)

func stripComments(line string) string {
	splitLine := strings.SplitN(line, "//", 2)
	return strings.TrimSpace(splitLine[0])
}

func Parse(line string) {
	strippedLine := stripComments(line)
	if len(strippedLine) == 0 {
		return
	}
	var inst code.Instruction
	if strippedLine[0] == '@' {
		location := strings.TrimPrefix(strippedLine, "@")
		inst = code.Ainst{code.AINST, location}
	} else if strippedLine[0] == '(' {
		symbol := strings.Trim(strippedLine, "()")
		symboltable.Table[symbol] = len(code.Program)
		return
	} else {
		// first determine if there is a dest
		var dest string
		checkEq := strings.Split(strippedLine, "=")
		if len(checkEq) == 2 {
			dest = checkEq[0]
			strippedLine = checkEq[1]
		}

		// next, check for a jump
		var jump string
		checkSemi := strings.Split(strippedLine, ";")
		if len(checkSemi) == 2 {
			jump = checkSemi[1]
			strippedLine = checkSemi[0]
		}

		// finally, handle comp
		comp := strippedLine

		inst = code.Cinst{code.CINST, comp, dest, jump}
	}
	code.AddLine(inst)
}
