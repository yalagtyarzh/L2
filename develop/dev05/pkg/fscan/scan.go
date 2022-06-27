package fscan

import (
	"bufio"
	"os"
)

// Scan читает все строки из файла и возвращает массив прочитанных строк
func Scan(file *os.File) ([]string, error) {
	scanner := bufio.NewScanner(file)

	var strs []string
	for scanner.Scan() {
		strs = append(strs, scanner.Text())
	}

	return strs, scanner.Err()
}
