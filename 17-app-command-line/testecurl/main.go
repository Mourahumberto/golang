package main

import (
	"net/http"
)

func main() {

	resp, err := http.Get("https://www.canhazip.com")
	if err != nil {
		print(err)
	}
	defer resp.Body.Close()
}

func teste() error {

}
