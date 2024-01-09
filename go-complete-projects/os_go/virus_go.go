package virus

import (
	"fmt"
	"log"
	"os"
)

func ReadFile(path string) {
	file, err := os.Open(path) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %d bytes: %q\n", count, data[:count])
}
