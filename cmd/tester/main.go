package main

import (
	"fmt"
	"net/http"

	"github.com/inconshreveable/go-update"
)

func main() {
	url := "https://github.com/1337jazz/go-update-test/releases/download/0.0.3/go-update-test_Linux_x86_64.tar.gz"
	err := doUpdate(url)
	if err != nil {
		panic(err)
	}

	fmt.Println("Update successful, version 0.0.5 now!")
}

func doUpdate(url string) error {

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = update.Apply(resp.Body, update.Options{})
	if err != nil {
		// error handling
	}
	return err
}
