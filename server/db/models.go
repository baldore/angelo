// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"encoding/json"
)

type Song struct {
	ID     int32           `json:"id"`
	Name   string          `json:"name"`
	Labels json.RawMessage `json:"labels"`
}
