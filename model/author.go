package model

import "time"

type Author struct {
	Author_ID    int
	Name         string
	Pen_Name     string
	Email        string
	Password     string
	Is_Disabled  bool
	Created_Time time.Time
}
