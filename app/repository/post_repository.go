package repository

import (
"app/model"
)

func PostCount(user *model.User) error{
	query := `SELECT COUNT(*) AS PostCount
						FROM
						(
							SELECT * FROM articles WHERE userid = ?
							UNION
							SELECT * FROM questions WHERE userid = ?
							) AS posttable`;
	if err := db.Get(user, query, user.ID, user.ID); err != nil {
		return err
	}
	return nil
}