package types

import (
	"golang.org/x/time/rate"
)

type Client struct {
	Limiter *rate.Limiter
}

type NoResponse struct {
	Lang   string `json:"lang"`
	Reason string `json:"reason"`
}
