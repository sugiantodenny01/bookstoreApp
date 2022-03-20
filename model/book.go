package model

import "time"

type Book struct {
	Book_ID      int
	Author_ID    int
	Title        string
	Summary      string
	Stock        int
	Price        int
	Cover_URL    string
	Created_Time time.Time
}
