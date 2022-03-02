package main

import (
	"encoding/json"
	"time"

	todo "github.com/petrostrak/command-line-apps-in-Go/Todo"
)

type todoResponse struct {
	Results todo.List `json:"results"`
}

func (r *todoResponse) MarshalJSON() ([]byte, error) {
	resp := struct {
		Results      todo.List `json:"results"`
		Date         int64     `json:"date"`
		TotalResults int       `json:"total_results"`
	}{
		r.Results,
		time.Now().Unix(),
		len(r.Results),
	}

	return json.Marshal(resp)
}
