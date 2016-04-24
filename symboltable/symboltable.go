package symboltable

type SymbolTable map[string]int

var Table SymbolTable
var nextVar int = 16

func Init() {
  Table = SymbolTable{
    "SP": 0,
    "LCL": 1,
    "ARG": 2,
    "THIS": 3,
    "THAT": 4,
    "R0": 0,
    "R1": 1,
    "R2": 2,
    "R3": 3,
    "R4": 4,
    "R5": 5,
    "R6": 6,
    "R7": 7,
    "R8": 8,
    "R9": 9,
    "R10": 10,
    "R11": 11,
    "R12": 12,
    "R13": 13,
    "R14": 14,
    "R15": 15,
    "SCREEN": 16384,
    "KBD": 24576,
  }
}

func FetchVariable(v string) int {
  val, exists := Table[v]
  if exists {
    return val
  }
  Table[v] = nextVar
  nextVar += 1
  return Table[v]
}
