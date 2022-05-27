package personalia

import (
	"github.com/golang_pos_app_service/pkg"

	userentity "github.com/golang_pos_app_service/entity/personalia"
)

const queryGetAllUser = `
	SELECT user_id, user_name, full_name, password, is_admin  
	FROM users 
	ORDER BY user_id
`

func GetAllUserController() ([]userentity.User, error) {
	var users []userentity.User

	db, err := pkg.GetPgConn()
	if err != nil {
		return users, err
	}
	err = db.Select(&users, queryGetAllUser)
	if err != nil {
		return users, err
	}
	return users, nil
}
