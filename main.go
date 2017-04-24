package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"sort"
)






func inMemorySort(p sort.Interface, r io.Reader, w io.Writer) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	var split = bytes.Split(b, []byte("\n"))

	sortBytes(split)

	for _, v := range split {
		const newLine = byte('\n')
		_, err = w.Write(append(v, newLine))
		if err != nil {
			return err
		}
	}

	return nil
}

func merge(a, b io.Reader, w io.Writer) {

}

func main () {
	err := inMemorySort(os.Stdin, os.Stdout)
	if err != nil {
		panic(err)
	}
}
