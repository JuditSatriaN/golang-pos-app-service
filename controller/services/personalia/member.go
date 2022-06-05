package personalia

import (
	"context"

	"github.com/golang_pos_app_service/pkg"

	memberentity "github.com/golang_pos_app_service/entity/personalia"
)

const queryGetAllMember = `
	SELECT id, name  
	FROM members 
	ORDER BY id
`

func GetAllMember(ctx context.Context) ([]memberentity.Member, error) {
	var members []memberentity.Member

	db, err := pkg.GetPgConn()
	if err != nil {
		return members, err
	}

	err = db.SelectContext(ctx, &members, queryGetAllMember)
	if err != nil {
		return members, err
	}

	return members, nil
}
