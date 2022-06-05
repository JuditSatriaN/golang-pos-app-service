package personalia

import (
	"context"
	"database/sql"

	"github.com/golang_pos_app_service/pkg"

	memberentity "github.com/golang_pos_app_service/entity/personalia"
)

const queryGetAllMember = `
	SELECT id, name, phone  
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

const queryGetMemberByID = `
	SELECT id, name, phone  
	FROM members 
	WHERE id = $1
`

func GetMemberByID(ctx context.Context, ID string) (memberentity.Member, bool, error) {
	var member memberentity.Member

	db, err := pkg.GetPgConn()
	if err != nil {
		return member, false, err
	}
	err = db.GetContext(ctx, &member, queryGetMemberByID, ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return member, false, nil
		}
		return member, false, err
	}
	return member, true, nil
}

const insertMember = `
	INSERT INTO members (id, name, phone) 
	VALUES (:id, :name, :phone)
`

func InsertMember(ctx context.Context, member memberentity.Member) error {

	db, err := pkg.GetPgConn()
	if err != nil {
		return err
	}

	_, err = db.NamedQueryContext(ctx, insertMember, member)
	if err != nil {
		return err
	}
	return nil
}

const updateMember = `
	UPDATE members SET 
		name = :name,
	    phone = :phone,
		update_time = NOW()
	WHERE id = :id
`

func UpdateMember(ctx context.Context, member memberentity.Member) error {

	db, err := pkg.GetPgConn()
	if err != nil {
		return err
	}

	_, err = db.NamedQueryContext(ctx, updateMember, member)
	if err != nil {
		return err
	}
	return nil
}

const deleteMember = `
	DELETE FROM members 
	WHERE id = $1
`

func DeleteMember(ctx context.Context, ID string) error {

	db, err := pkg.GetPgConn()
	if err != nil {
		return err
	}

	_, err = db.ExecContext(ctx, deleteMember, ID)
	if err != nil {
		return err
	}
	return nil
}
