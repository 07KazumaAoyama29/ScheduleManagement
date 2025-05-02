package model

import "time"

type Event struct {
    ID      string    `db:"id" json:"id"`
    Title   string    `db:"title" json:"title"`
    StartAt time.Time `db:"start_at" json:"start_at"`
    EndAt   time.Time `db:"end_at" json:"end_at"`
    RRule   string    `db:"rrule" json:"rrule"`
}