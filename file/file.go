package file

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
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
func IsExists(name string) bool {
	_, err := os.Stat(name)
	return err == nil || os.IsExist(err)
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

func ReadByteFromFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

func ReadLinesFromFile(filename string) ([]string, error) {
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
