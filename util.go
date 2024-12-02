package adventofcode2024

import (
	"bufio"
	"os"
)

func ReadToScanner(filename string, f func(*bufio.Scanner)) {
	file, err := os.Open(filename)
	if err != nil {
		panic("jweejflkwe$@#$@#")
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic("failed to close file")
		}
	}(file)

	f(bufio.NewScanner(file))
}
