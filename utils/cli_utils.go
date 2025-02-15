package utils

import (
	"bufio"
	"fmt"
	"miniimpplus/constants"
	"os"
	"os/exec"
	"strings"
)

// PrintBannerOnConsole displays a colored console banner for the tool in interactive mode
// MiniImp+ ASCII art courtesy of http://www.network-science.de/ascii/
func PrintBannerOnConsole() {
	fmt.Println(constants.StyledGreen(`
	##     ## #### ##    ## #### #### ##     ## ########         
	###   ###  ##  ###   ##  ##   ##  ###   ### ##     ##   ##
	#### ####  ##  ####  ##  ##   ##  #### #### ##     ##   ##   
	## ### ##  ##  ## ## ##  ##   ##  ## ### ## ########  ###### 
	##     ##  ##  ##  ####  ##   ##  ##     ## ##          ##   
	##     ##  ##  ##   ###  ##   ##  ##     ## ##          ##   
	##     ## #### ##    ## #### #### ##     ## ##
    `))
	fmt.Println(constants.StyledYellow("[+] MiniImp+ to Python source-to-source compiler, built with ANTLRv4 and Go"))
	fmt.Println(constants.StyledYellow("------------------------------------------------------------------------"))
}

// AskUserToExecute prompts the user if they would like to run the output Python source code
func AskUserToExecute(outputFile string) {
	for {
		fmt.Print(constants.StyledGreen("\n[+] Would you like to execute the compiled Python file? (y/n): "))

		reader := bufio.NewReader(os.Stdin)
		response, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading input: %v\n", err)
			return
		}
		response = strings.TrimSpace(response)

		if response == "y" {
			cmd := exec.Command(constants.PyExecCommand, outputFile)
			// make sure Python process gets user input from terminal
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			fmt.Printf("\n[+] Running Python file with command `%v`\n", constants.StyledGreen(cmd.String()))
			fmt.Printf(constants.StyledYellow("------------------------------------------------------------------------------------"))
			fmt.Println()
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			err := cmd.Run()
			if err != nil {
				fmt.Printf("[-] Error executing Python file: %v\n", err)
			}
			return
		} else if response == "n" {
			fmt.Println("Compilation completed")
			return
		} else {
			fmt.Println("Invalid input - please enter either `y` or `n`")
		}
	}
}
