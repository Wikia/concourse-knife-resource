package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-chef/chef"
)

type Payload struct {
	Source Source `json:"source"`
	Params Params `json:"params"`
}

type Source struct {
	URI         string `json:"uri"`
	Client      string `json:"client"`
	Certificate string `json:"certificate"`
	SkipSSL     bool   `json:"skipssl"`
}

type Params struct {
	Index   string                 `json:"index"`
	Query   string                 `json:"query"`
	Filters map[string]interface{} `json:"filters"`
}

func main() {
	var payload Payload
	json.NewDecoder(os.Stdin).Decode(&payload)

	client, err := chef.NewClient(&chef.Config{
		Name:    payload.Source.Client,
		Key:     payload.Source.Certificate,
		BaseURL: payload.Source.URI,
		SkipSSL: payload.Source.SkipSSL || false,
	})

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating client: ", err)
		os.Exit(1)
	}

	if payload.Params.Filters == nil {
		payload.Params.Filters = make(map[string]interface{})
	}

	result, err := client.Search.PartialExec(payload.Params.Index, payload.Params.Query, payload.Params.Filters)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error executing query ", err)
	}

	json.NewEncoder(os.Stdout).Encode(&result.Rows)
}
