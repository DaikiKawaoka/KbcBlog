package model

import "time"

  // Article ...
  type Article struct {
		ID      int       `db:"id" json:"id"`
		Userid  int       `db:"userid" json:"userid"`
		Title   string    `db:"title" json:"title"`
		Body    string    `db:"body" json:"body"`
		Created time.Time `db:"created"`
		Updated time.Time `db:"updated"`
  }