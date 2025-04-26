package main

import (
	"fmt"
	"io"
	"net/http"
)

type logWriter struct{}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}

func main() {
	resp, error := http.Get("https://www.sijosam.in/blog/use-sync-external-store/")

	if error != nil {
		panic(error)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))

}
