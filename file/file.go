package file

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// SelfPath gets compiled executable file absolute path
func SelfPath() string {
	path, _ := filepath.Abs(os.Args[0])
	return path
}

// SelfDir gets compiled executable file directory
func SelfDir() string {
	return filepath.Dir(SelfPath())
}

// FileExists reports whether the named file or directory exists.
func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func WriteStringsToFile(data []string, fileName string) error {
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("create map file error: %v\n", err)
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for _, v := range data {
		_, _ = fmt.Fprintln(w, v)
	}
	return w.Flush()
}

func ReadStringsFromFile(filename string) ([]string, error) {
	var output []string

	fi, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return output, err

	}
	defer fi.Close()
	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		output = append(output, string(a))
	}
	return output, nil
}