package personalia

import (
	"context"
	"database/sql"
	"github.com/golang_pos_app_service/pkg"

	userentity "github.com/golang_pos_app_service/entity/personalia"
)

const queryGetAllUser = `
	SELECT user_id, user_name, full_name, password, is_admin  
	FROM users 
	ORDER BY user_id
`

func GetAllUser(ctx context.Context) ([]userentity.User, error) {
	var users []userentity.User

	db, err := pkg.GetPgConn()
	if err != nil {
		return users, err
	}
	err = db.SelectContext(ctx, &users, queryGetAllUser)
	if err != nil {
		return users, err
	}
	return users, nil
}

const queryGetUserByUserID = `
	SELECT user_id, user_name, full_name, password, is_admin  
	FROM users 
	WHERE user_id = $1
`

func GetUserByUserID(ctx context.Context, userID string) (userentity.User, bool, error) {
	var user userentity.User

	db, err := pkg.GetPgConn()
	if err != nil {
		return user, false, err
	}
	err = db.GetContext(ctx, &user, queryGetUserByUserID, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, false, nil
		}
		return user, false, err
	}
	return user, true, nil
}

const insertUser = `
	INSERT INTO users (user_id, user_name, full_name, password, is_admin) 
	VALUES (:user_id, :user_name, :full_name, :password, :is_admin)
`

func InsertUser(ctx context.Context, user userentity.User) error {

	db, err := pkg.GetPgConn()
	if err != nil {
		return err
	}

	_, err = db.NamedQueryContext(ctx, insertUser, user)
	if err != nil {
		return err
	}
	return nil
}

const updateUser = `
	UPDATE users SET 
		user_name = :user_name, 
		full_name = :full_name, 
		password = :password, 
		is_admin = :is_admin 
	WHERE user_id = :user_id
`

func UpdateUser(ctx context.Context, user userentity.User) error {

	db, err := pkg.GetPgConn()
	if err != nil {
		return err
	}

	_, err = db.NamedQueryContext(ctx, updateUser, user)
	if err != nil {
		return err
	}
	return nil
}

const deleteUser = `
	DELETE FROM users 
	WHERE user_id = $1
`

func DeleteUser(ctx context.Context, userID string) error {

	db, err := pkg.GetPgConn()
	if err != nil {
		return err
	}

	_, err = db.ExecContext(ctx, deleteUser, userID)
	if err != nil {
		return err
	}
	return nil
}
