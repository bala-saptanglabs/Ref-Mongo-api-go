package model

import(
	"time"
)

type Task struct {
	Id *string `json:"id,omitempty" bson:"_id,omitempty"`
	Title *string `json:"title" bson:"title"`
	Completed *bool	`json:"completed" bson:"completed"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdateAt time.Time `json:"updatedAt" bson:"updatedAt"`
}
