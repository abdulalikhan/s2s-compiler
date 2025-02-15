package main

import (
	"bufio"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"io/ioutil"
	"log"
	"miniimpplus/constants"
	"miniimpplus/parser"
	"miniimpplus/utils"
	"miniimpplus/visitors"
	"os"
	"path/filepath"
	"strings"
)

func compileMipFile(inputFile string) string {
	input, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("[-] Could not read file %s: %v", inputFile, err)
	}

	// create ANTLR input stream
	is := antlr.NewInputStream(string(input))
	lexer := parser.NewMiniImpLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewMiniImpParser(stream)

	// the input will now be parsed from the `prog` rule
	tree := p.Prog().(*parser.ProgContext)
	pv := visitors.NewPythonVisitor()
	output := pv.Generate(tree)

	// write the python code to the output folder
	outputDir := filepath.Join(filepath.Dir(inputFile), "output")
	err = os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		log.Fatalf("[-] Error creating output directory: %v", err)
	}
	baseName := strings.TrimSuffix(filepath.Base(inputFile), constants.MipFileExt)
	outFilename := filepath.Join(outputDir, baseName+".py")

	err = ioutil.WriteFile(outFilename, []byte(output), 0644)
	if err != nil {
		log.Fatalf("[-] Error writing output: %v", err)
	}

	fmt.Printf("\n[+] Python code written to: %s\n", constants.StyledYellow(outFilename))
	return outFilename
}

func main() {
	utils.PrintBannerOnConsole()

	// 1 - the file to be compiled can be either supplied via a CLI argument
	if len(os.Args) > 1 {
		// Process the supplied file directly
		inputFile := os.Args[1]
		if !strings.HasSuffix(inputFile, constants.MipFileExt) {
			log.Fatalf("[-] The provided file is not a .mip file: %v", inputFile)
		}
		compileMipFile(inputFile)
		return
	}

	// or
	// 2- if no arguments are supplied - run the tool in interactive mode
	miFiles, err := utils.ListMIFiles()
	if err != nil {
		log.Fatalf("[-] Error reading directory: %v", err)
	}

	if len(miFiles) == 0 {
		log.Fatalf("[-] No MiniImp+ (.mip) files found in the current directory")
	}

	fmt.Print(constants.StyledGreen("\n[+] Available MiniImp+ Files:"))
	fmt.Println()
	for i, file := range miFiles {
		fmt.Printf("%d: %s\n", i+1, constants.StyledYellow(file))
	}

	fmt.Print(constants.StyledGreen("\n[+] Enter the number of the file to compile: "))
	var choice int
	_, err = fmt.Scanf("%d", &choice)
	if err != nil || choice < 1 || choice > len(miFiles) {
		log.Fatalf("[-] Invalid choice.")
	}
	bufio.NewReader(os.Stdin).ReadString('\n')

	inputFile := miFiles[choice-1]

	// compile the .mip file into Python source code
	outputFile := compileMipFile(inputFile)

	// output the Python source code
	fmt.Println(constants.StyledGreen("\n[+] Generated Python Source Code:\n"))
	pythonCode, err := os.ReadFile(outputFile)
	if err != nil {
		log.Fatalf("[-] Failed to read compiled Python file: %v", err)
	}
	fmt.Println(constants.StyledYellow(string(pythonCode)))

	// the user will be prompted - they can either execute the python file right here or on their own
	utils.AskUserToExecute(outputFile)
}
