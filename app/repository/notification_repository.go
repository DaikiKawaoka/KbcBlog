package repository

import (
	// "database/sql"
	"time"
	"app/model"
	)

// NotificationCreate ...
func NotificationCreate(notification *model.CreateNotification) error {
	now := time.Now()
  notification.Created = now.Format("2006/01/02 15:04:05")
  notification.Updated = now.Format("2006/01/02 15:04:05")

	var query string
	switch notification.Action {
	case "alike":
		query = `INSERT INTO notifications (articleid, visiterid, visitedid, action, created, updated)
    VALUES (:articleid, :visiterid, :visitedid, :action, :created, :updated);`
	case "qlike":
		query = `INSERT INTO notifications (questionid, visiterid, visitedid, action, created, updated)
    VALUES (:questionid, :visiterid, :visitedid, :action, :created, :updated);`
	case "follow":
		query = `INSERT INTO notifications (visiterid, visitedid, action, created, updated)
    VALUES (:visiterid, :visitedid, :action, :created, :updated);`
	case "aclike":
		query = `INSERT INTO notifications (articleid,a_commentid,visiterid, visitedid, action, created, updated)
    VALUES (:articleid,:a_commentid,:visiterid, :visitedid, :action, :created, :updated);`
	case "qclike":
		query = `INSERT INTO notifications (questionid,q_commentid,visiterid, visitedid, action, created, updated)
    VALUES (:questionid,:q_commentid,:visiterid, :visitedid, :action, :created, :updated);`
	case "acomment":
		query = `INSERT INTO notifications (articleid,a_commentid,visiterid, visitedid, action, created, updated)
    VALUES (:articleid,:a_commentid,:visiterid, :visitedid, :action, :created, :updated);`
	case "qcomment":
		query = `INSERT INTO notifications (questionid,q_commentid,visiterid, visitedid, action, created, updated)
    VALUES (:questionid,:q_commentid,:visiterid, :visitedid, :action, :created, :updated);`
}
  tx := db.MustBegin()
  _,err := tx.NamedExec(query, notification)
  if err != nil {
    tx.Rollback()
    return err
  }
  tx.Commit()
  return nil
}

// GetNotificationCount ...
func GetNotificationCount(userID int) (int,error){
	query := `SELECT COUNT(*) FROM notifications WHERE visitedid = ? AND checked = false;`
	var count int
	if err := db.Get(&count, query, userID); err != nil {
		return 0, err
	}
	return count, nil
}

// GetNotification ...
func GetNotification(userID int,cursor int) ([]*model.Notification , error){
	var notifications []*model.Notification
	notifications = make([]*model.Notification, 0, 4)
	query := `SELECT n.id id, a.id articleid,q.id questionid, n.visiterid userid, u2.name name,ac.text a_text,qc.text q_text, n.action action, n.checked checked, n.created created, n.updated updated FROM notifications n left join users u on n.visitedid = u.id left join users u2 on n.visiterid = u2.id left join articles a on n.articleid = a.id left join questions q on n.questionid = q.id left join article_comments ac on n.a_commentid = ac.id left join question_comments qc on n.q_commentid = qc.id WHERE u.id = ? ORDER BY n.id desc LIMIT ?,4;`
	if err := db.Select(&notifications, query, userID, cursor); err != nil {
		return nil, err
	}
	return notifications, nil
}

// NotificationChecked ...
func NotificationChecked(userID int) error{
	var notification model.CreateNotification
	now := time.Now()
	notification.Updated = now.Format("2006/01/02 15:04:05")
	notification.Visitedid = userID

	query := `UPDATE notifications
	SET  updated= :updated,
			checked = true
	WHERE visitedid = :visitedid AND checked = false ;`
	tx := db.MustBegin()
	_ , err := tx.NamedExec(query, notification)

	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// NotificationDelete ...
func NotificationDelete(userID int) error {
	query := "DELETE FROM notifications WHERE visitedid = ?"
	tx := db.MustBegin()
	if _, err := tx.Exec(query, userID); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}
