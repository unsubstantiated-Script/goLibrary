package funky_funcs

import (
	"bytes"
	"fmt"
)

func RollByteBuffer() {
	b := bytes.NewBufferString("Hello ")
	fmt.Println(b.String())

	b.WriteString("Gophers!")
	fmt.Println(b.String())

	b.Reset()
	b.WriteString("It's Thursday")
	fmt.Println(b.String())

	b.Write([]byte("Happy Happy!"))
	fmt.Println(b.String())
}
