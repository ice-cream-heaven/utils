package mq

import "time"

type Message[M any] struct {
	Value M `json:"value,omitempty" yaml:"value,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty" yaml:"created_at,omitempty"`
}
