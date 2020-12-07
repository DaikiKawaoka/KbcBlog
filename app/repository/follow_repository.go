package repository

import (
	"time"
	"app/model"
	)

	// フォローする人は「follower」
	// フォローされる人は「followed」


	// フォローしているかチェック
	func CheckFollow(followerId int, followedId int) (int, error) {
		query := `SELECT COUNT(*) FROM follows WHERE followerid = ? AND followedid = ?;`
		var count int
		if err := db.Get(&count, query, followerId, followedId); err != nil {
			return 0, err
		}
		return count, nil
	}

	// 自分のフォローユーザの情報を取得
	func GetFollowerInfoList(userId int) ([]*model.FUser, error) {
		query := `SELECT u.id,u.name,u.comment FROM users u,follows f WHERE u.id = f.followedid AND f.followerid = ?;`
		var users []*model.FUser
		if err := db.Select(&users, query, userId); err != nil {
			return nil, err
		}
		return users, nil
	}

	// 自分のフォロワーの情報を取得
	func GetFollowedInfoList(userId int) ([]*model.FUser, error) {
		query := `SELECT u.id,u.name,u.comment FROM users u,follows f WHERE u.id = f.followerid AND f.followedid = ?;`
		var users []*model.FUser
		if err := db.Select(&users, query, userId); err != nil {
			return nil, err
		}
		return users, nil
	}

	// フォロー人数を取得
	func GetFollowerCount(userId int) (int, error) {
		query := `SELECT COUNT(*) FROM follows WHERE followerid = ?;`
		var count int
		if err := db.Get(&count, query, userId); err != nil {
			return 0, err
		}
		return count, nil
	}

	// フォロワー人数を取得
	func GetFollowedCount(userId int) (int, error) {
		query := `SELECT COUNT(*) FROM follows WHERE followedid = ?;`
		var count int
		if err := db.Get(&count, query, userId); err != nil {
			return 0, err
		}
		return count, nil
	}

	func Follow(follow *model.Following) error {
		now := time.Now()
		follow.Created = now.Format("2006/01/02 15:04:05")
		query := `INSERT INTO follows (followerid, followedid, created)
		VALUES (:followerid, :followedid, :created);`
		tx := db.MustBegin()
		_, err := tx.NamedExec(query, follow)
		if err != nil {
			tx.Rollback()
			return err
		}
		tx.Commit()
		return nil
	}

	func UnFollow(followerid int,followedid int) error {
		query := "DELETE FROM follows WHERE followerid = ? AND followedid = ?"
		tx := db.MustBegin()
		if _, err := tx.Exec(query,followerid,followedid); err != nil {
			tx.Rollback()
			return err
		}
		return tx.Commit()
	}