// Измените программу echo так, чтобы она выводила также os.Args[0], имя выполняемой команды.
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	s, sep := filepath.Base(os.Args[0]), " "

	for _, arg := range os.Args[1:] {
		s += sep + arg
	}
	fmt.Println(s)
}
