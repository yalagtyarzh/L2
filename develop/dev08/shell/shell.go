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

// StartShell начинает работу кастомной shell программы
func StartShell() {
	reader := bufio.NewReader(os.Stdin)
	for {
		// Получаем нынешнюю директорию и выводим ее на экран
		path, _ := filepath.Abs(".")
		fmt.Printf("%s> ", path)

		// Получаем команду
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		// Обрабатываем команду
		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

// execInput получает заданную команду в консоли, делит ее на аргументы и запускает на исполнение
func execInput(input string) error {
	// удаляем символ новой строки
	input = strings.TrimSuffix(input, "\n")

	// Получаем аргументы команды
	args := strings.Split(input, " ")

	switch args[0] {
	// При команде cd меняем директорию на заданную
	case "cd":
		if len(args) < 2 {
			return errors.New("path required")
		}

		return os.Chdir(args[1])
	// При команде получаем директорию, на который находиться пользователь и выводим ее в консоль
	case "pwd":
		pwd, err := os.Getwd()
		if err != nil {
			return err
		}

		fmt.Println(pwd)
		return nil
	// При команде echo выводим все аргументы команды, кроме первого
	case "echo":
		for i := 1; i < len(args); i++ {
			fmt.Print(args[i], " ")
		}

		fmt.Println()
		return nil
	// При команде kill находим pid процесса и убиваем процесс
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
	// При команде ps находим все работающие процессы и выводим их на экран
	case "ps":
		procs, err := ps.Processes()
		if err != nil {
			return err
		}

		for _, proc := range procs {
			fmt.Printf("%d\t%s\n", proc.Pid(), proc.Executable())
		}

		return nil
	// При команде \quit просто выходим из программы
	case "\\quit":
		os.Exit(0)
	// В иных случаях запускаем unix программы
	default:
		cmd := exec.Command(args[0], args[1:]...)

		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout

		return cmd.Run()
	}

	return nil
}
