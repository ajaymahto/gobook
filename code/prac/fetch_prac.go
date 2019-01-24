package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	urls := os.Args[1:]
	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error fetching Url : %s", url)
			continue
		}
		b, err := io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Printf("Error printing to Stdout", err)
			continue
		}
		fmt.Println("%s", b)
	}
}
