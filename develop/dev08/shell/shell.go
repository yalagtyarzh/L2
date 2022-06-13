package shell

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/mitchellh/go-ps"
)

func StartShell() {
	reader := bufio.NewReader(os.Stdin)
	for {
		path, _ := filepath.Abs(".")
		fmt.Printf("%s> ", path)

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func execInput(input string) error {
	input = strings.TrimSuffix(input, "\n")

	args := strings.Split(input, " ")

	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return errors.New("path required")
		}

		return os.Chdir(args[1])
	case "pwd":
		pwd, err := os.Getwd()
		if err != nil {
			return err
		}

		fmt.Println(pwd)
		return nil
	case "echo":
		for i := 1; i < len(args); i++ {
			fmt.Print(args[i], " ")
		}

		fmt.Println()
		return nil
	case "kill":
		pid, err := strconv.Atoi(args[1])
		if err != nil {
			return err
		}

		proc, err := os.FindProcess(pid)
		if err != nil {
			return err
		}

		return proc.Kill()
	case "ps":
		procs, err := ps.Processes()
		if err != nil {
			return err
		}

		for _, proc := range procs {
			fmt.Printf("%d\t%s\n", proc.Pid(), proc.Executable())
		}

		return nil
	case "\\quit":
		os.Exit(0)
	default:
		cmd := exec.Command(args[0], args[1:]...)

		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout

		return cmd.Run()
	}

	return nil
}
