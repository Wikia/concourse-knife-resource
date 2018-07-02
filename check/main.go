package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-chef/chef"
)

type Payload struct {
	Source  Source `json:"source"`
	Version string `json:"version"`
}

type Source struct {
	URI         string                 `json:"uri"`
	Client      string                 `json:"client"`
	Certificate string                 `json:"certificate"`
	SkipSSL     bool                   `json:"skipssl"`
	Index       string                 `json:"index"`
	Query       string                 `json:"query"`
	Filters     map[string]interface{} `json:"filters"`
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

	if payload.Source.Filters == nil {
		payload.Source.Filters = make(map[string]interface{})
	}

	result, err := client.Search.PartialExec(payload.Source.Index, payload.Source.Query, payload.Source.Filters)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error executing query: ", err)
	}

	version, err := json.Marshal(result.Rows)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to marshal version: ", err)
	}

	payload.Version = string(version)
	json.NewEncoder(os.Stdout).Encode(&payload)
}
