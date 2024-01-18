package funky_funcs

import (
	"log"
	"os"
)

func RollWriteOut() {
	f, err := os.Create("output.txt")

	if err != nil {
		log.Fatalf("error %s", err)
	}

	defer f.Close()

	s := []byte("Hello gophers!")

	_, err = f.Write(s)
	if err != nil {
		log.Fatalf("error %s", err)
	}
}
