package file

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

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