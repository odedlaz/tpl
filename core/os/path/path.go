package path

import "os"

// IsDir true if a path is a directory
func IsDir(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), nil
}

// FileMode returns the file mode bits of a file
func FileMode(path string) (os.FileMode, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	// file mode bits
	return fileInfo.Mode(), nil
}
