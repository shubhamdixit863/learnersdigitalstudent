package repository

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"session26/internal/models"
	"strconv"
)

type MysqlRepo struct {
	conn *sqlx.DB
}

func (m MysqlRepo) UpdateUser(ctx context.Context, user models.User) error {
	_, err := m.conn.ExecContext(ctx, `
		UPDATE Users 
		SET UserName = ?, Password = ?,FirstName = ?, SecondName = ?
		WHERE ID = ?`,
		user.Username, user.Password, user.FirstName, user.SecondName, user.ID,
	)
	fmt.Println(err)
	return err

	//panic("implement me")
}

func (m MysqlRepo) DeleteUser(ctx context.Context, user models.User) error {
	_, err := m.conn.ExecContext(ctx, "DELETE FROM Users WHERE ID = ?", user.ID)
	return err
	//panic("implement me")
}

func (m MysqlRepo) CreateUser(ctx context.Context, user models.User) (interface{}, error) {

	result, err := m.conn.ExecContext(ctx, `INSERT INTO Users (FirstName, SecondName, UserName, password) VALUES (?, ?, ?, ?)`,
		user.FirstName, user.SecondName, user.Username, user.Password)

	if err != nil {
		return "", err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return "", err
	}

	return strconv.Itoa(int(id)), nil

}

func (m MysqlRepo) GetUserByUserName(ctx context.Context, userName string) (*models.User, error) {
	var user models.User
	row := m.conn.QueryRowx("select FirstName, SecondName, UserName from Users where userName=?", userName)
	err := row.Scan(&user.FirstName, &user.SecondName, &user.Username)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (m MysqlRepo) GetAllUsers(ctx context.Context) ([]*models.User, error) {

	users := []*models.User{}
	rows, err := m.conn.QueryxContext(ctx, "SELECT * FROM Users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err := rows.StructScan(&user); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func NewMysqlReqo(db *sqlx.DB) DbRepository {
	return &MysqlRepo{ //makes our code loosely coupled
		db,
	}

}
