package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">> ")
		// Read the keyboard input
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		// Handle the execution of the input
		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

var ErrNoPath = errors.New("Path Required")

func execInput(input string) error {
	// Remove the new line character
	input = strings.TrimSuffix(input, "\n")

	// Split the input to seperate the command and the arguments
	args := strings.Split(input, " ")

	// Pass the program and the arguments separately
	cmd := exec.Command(args[0], args[1:]...)

	// Set the correct device
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// Execute the command and return the error
	return cmd.Run()
}
