package entity

import "time"

type Transaction struct {
	date   time.Time
	status Status
}
