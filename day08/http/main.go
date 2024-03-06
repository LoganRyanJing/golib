package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// func main() {
// 	resp, err := http.Get("https://api.github.com/users/microsoft/repos?page=15&per_page=3")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		os.Exit(1)
// 	}

// 	io.Copy(os.Stdout, resp.Body)
// }

type GitHubResponse []struct {
	fullName string `json:"full_name"`
}

func (w customWriter) Write(p []byte) (n int, err error) {
	var resp GitHubResponse
	json.Unmarshal(p, &resp)
	for _, r := range resp {
		fmt.Println(r.fullName)
	}
	return len(p), nil
}

func main() {
	resp, err := http.Get()
	if er != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	witer := customWriter{}
	io.Copy(writer, resp.Body)
}
