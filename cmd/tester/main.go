package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/inconshreveable/go-update"
)

type Response struct {
	Url string `json:"url"`
}

func main() {

	var latestReleaseUrl = "https://api.github.com/repos/1337jazz/go-update-test/releases/latest"
	resp, err := http.Get(latestReleaseUrl)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// Decode the JSON response into our struct type.
	var release Response
	err = json.NewDecoder(resp.Body).Decode(&release)
	if err != nil {
		panic(err)
	}

	url := release.Url

	err = doUpdate(url)
	if err != nil {
		panic(err)
	}

	// get the version from version.txt

	data, err := os.ReadFile("version.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Updated to version: ", string(data))
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
