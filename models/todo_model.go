package models

import "time"

type TodoModel struct {
	ID          string    `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
