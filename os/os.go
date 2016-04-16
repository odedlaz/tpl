package os

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

func ReadFromStdIn() string {

	var buffer bytes.Buffer
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		buffer.WriteString(fmt.Sprintln(scanner.Text()))
	}
	return buffer.String()
}

func StdInAvailable() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) == 0
}

func ReadFile(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
