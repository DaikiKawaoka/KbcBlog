package model

import (
	"database/sql"
)

// CreateNotification ...
type CreateNotification struct {
	ID         int        `db:"id" json:"id"`
	Articleid  int        `db:"articleid" json:"articleid"`
	Questionid int        `db:"questionid" json:"questionid"`
	Visiterid  int        `db:"visiterid" json:"visiterid"`
	Visitedid  int        `db:"visitedid" json:"visitedid"`
	ACommentid int        `db:"a_commentid" json:"a_commentid"`
	QCommentid int        `db:"q_commentid" json:"q_commentid"`
	Action     string     `db:"action" json:"action"`  // acomment qcomment alike qlike aclike qclike follow
	Checked    bool       `db:"checked" json:"checked"`
	Created    string     `db:"created"`
	Updated    string     `db:"updated"`
}

// Notification ...
type Notification struct{
	ID         int        `db:"id" json:"id"`
	Articleid  sql.NullInt32         `db:"articleid" json:"articleid"`
	Questionid sql.NullInt32         `db:"questionid" json:"questionid"`
	Userid     int        `db:"userid" json:"userid"`
	Name       string     `db:"name" json:"name"`
	// ACommentid sql.NullInt32         `db:"a_commentid" json:"a_commentid"`
	// QCommentid sql.NullInt32       `db:"q_commentid" json:"q_commentid"`
	Atext    sql.NullString       `db:"a_text" json:"a_text"`
	Qtext    sql.NullString       `db:"q_text" json:"q_text"`
	Action     string     `db:"action" json:"action"`  // acomment qcomment alike qlike acommentlike qcommentlike follow
	Checked    bool       `db:"checked" json:"checked"`
	Created    string     `db:"created"`
	Updated    string     `db:"updated"`
}