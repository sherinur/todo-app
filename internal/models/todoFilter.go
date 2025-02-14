package models

import "time"

type TodoFilter struct {
	Done      bool      `json:"status"`
	Priority  string    `json:"priority"`
	DueBefore time.Time `json:"due_before"`
}
