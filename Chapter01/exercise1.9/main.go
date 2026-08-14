// Измените программу fetch так, чтобы она выводила код состо­яния HTTP,
// содержащийся в resp.Status.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		_, _ = io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
		fmt.Println(resp.Status)
	}
}
