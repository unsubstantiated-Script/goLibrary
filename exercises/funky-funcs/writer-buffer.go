package funky_funcs

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type personz struct {
	first string
}

func (p personz) writeOut(w io.Writer) error {
	_, err := w.Write([]byte(p.first))
	return err
}

func RollWriterBuffer() {
	p := personz{
		first: "Kakky",
	}

	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("error %s", err)
	}
	defer f.Close()

	var b bytes.Buffer

	//Writing out to the output.txt file
	p.writeOut(f)

	//Writing out to the buffer
	p.writeOut(&b)

	fmt.Println(b.String())
}
