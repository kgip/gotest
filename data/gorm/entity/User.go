package entity

import (
	"time"
)

type User struct {
	ID     int
	Name   string
	Age    int
	Gender byte
	Birth  time.Time
}
