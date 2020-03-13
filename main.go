package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	complete, err := fileComplete(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}

	if complete {
		fmt.Println("The file is complete")
	} else {
		fmt.Println("The file is not complete")
	}
}

func fileComplete(name string) (complete bool, err error) {

	archivo, err := os.Open(name)
	if err != nil {
		return
	}
	defer archivo.Close()

	bufer := make([]byte, 512)

	for {
		_, err = archivo.Read(bufer)

		if err != nil {
			if err == io.EOF {
				err = nil
				complete = true
			}

			break
		}
	}

	return
}
