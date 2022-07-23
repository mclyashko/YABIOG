package main

import (
	"flag"
	"fmt"
	"os"
)

func cpuResolverFuncGen(code []byte) func() {
	var cpu [30000]byte
	var codeIndex, cpuIndex int

	return func() {
		for codeIndex < len(code) {
			switch code[codeIndex] {
			case '+':
				cpu[cpuIndex]++
				codeIndex++
			case '-':
				cpu[cpuIndex]--
				codeIndex++
			case ',':
				_, err := fmt.Scan(&cpu[cpuIndex])
				if err != nil {
					panic("can't get user input: " + err.Error())
				}
				codeIndex++
			case '.':
				fmt.Print(string(cpu[cpuIndex]))
				codeIndex++
			case '[':
				if cpu[cpuIndex] == 0 {
					var brc = 1
					for brc > 0 {
						codeIndex++
						if code[codeIndex] == '[' {
							brc++
						} else if code[codeIndex] == ']' {
							brc--
						}
					}
				} else {
					codeIndex++
				}
			case ']':
				if cpu[cpuIndex] > 0 {
					var brc = 1
					for brc > 0 {
						codeIndex--
						if code[codeIndex] == '[' {
							brc--
						} else if code[codeIndex] == ']' {
							brc++
						}
					}
				} else {
					codeIndex++
				}
			case '>':
				cpuIndex++
				codeIndex++
			case '<':
				cpuIndex--
				codeIndex++
			default:
				codeIndex++
			}
		}
	}
}

func bFInterpreter(fileName string) {
	code, err := os.ReadFile(fileName)
	if err != nil {
		panic("can't read bf source code: " + err.Error())
	}

	cpuResolverFuncGen(code)()
}

func myCLI() (fileName string) {
	flag.StringVar(&fileName, "fileName", "in.bf", "bf source code file path "+
		"(pass only correct source code without comments)")
	flag.Parse()
	return
}

func main() {
	fileName := myCLI()
	bFInterpreter(fileName)
}
