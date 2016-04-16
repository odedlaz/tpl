package os

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

// ReadFromStdIn reads text from stdin
func ReadFromStdIn() string {

	var buffer bytes.Buffer
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		buffer.WriteString(fmt.Sprintln(scanner.Text()))
	}
	return buffer.String()
}

// StdInAvailable checks if there's data in stdin
func StdInAvailable() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) == 0
}

// ReadFile reads the file named by filename and returns the contents as string.
// A successful call returns err == nil, not err == EOF. Because ReadFile
// reads the whole file, it does not treat an EOF from Read as an error
// to be reported.
func ReadFile(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
