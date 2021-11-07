package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Package struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Version     string `json:"version"`
}

func GetPackages(search string) Package {
	req, err := http.Get("https://api.npms.io/v2/search?q=cross+spawn")
	if err != nil {
		panic(err)
	}

	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)

	p := Package{}
	jResp := json.Unmarshal([]byte(body), &p)
	if jResp != nil {
		panic(jResp)
	}

	fmt.Println(p)

	return p
}
