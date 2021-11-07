package network

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Package struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
}

type Result struct {
	Results Package `json:"package"`
}

type Response struct {
	Total   int      `json:"total"`
	Results []Result `json:"results"`
}

func GetPackages(search string) Response {
	url := fmt.Sprintf("https://api.npms.io/v2/search?q=%s", search)

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Reading some stuff")
		panic(err)
	}

	defer response.Body.Close()

	b, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var p Response

	err = json.Unmarshal([]byte(b), &p)
	if err != nil {
		panic(err)
	}

	fmt.Println(p)

	return p
}
