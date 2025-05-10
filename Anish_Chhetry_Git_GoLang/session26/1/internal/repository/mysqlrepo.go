package repository

import (
	"1/internal/models"
	"context"

	"github.com/jmoiron/sqlx"
)

type MysqlRepo struct {
	conn *sqlx.DB
}

func (m MysqlRepo) UpdateUser(ctx context.Context, user models.User) error {
	_, err := m.conn.ExecContext(ctx, `UPDATE Users SET FirstName=?, SecondName=?, Password=?, Email=?, ID=? WHERE Username=?`,
		user.FirstName, user.SecondName, user.Password, user.Email, user.ID, user.Username)
	if err != nil {
		return err
	}
	return nil
}

func (m MysqlRepo) DeleteUser(ctx context.Context, user models.User) error {
	_, err := m.conn.ExecContext(ctx, `DELETE FROM Users WHERE Username = ?`, user.Username)
	if err != nil {
		return err
	}
	return nil
}

func (m MysqlRepo) CreateUser(ctx context.Context, user models.User) (string, error) {

	_, err := m.conn.ExecContext(ctx, `INSERT INTO Users (FirstName, SecondName, UserName, Password, Email, ID) VALUES (?, ?, ?, ?, ?, ?)`,
		user.FirstName, user.SecondName, user.Username, user.Password, user.Email, user.ID)

	if err != nil {
		return "", err
	}

	return user.ID, nil

}

func (m MysqlRepo) GetUserByUserName(ctx context.Context, userName string) (*models.User, error) {
	var user models.User
	row := m.conn.QueryRowx("select Username,FirstName,SecondName,Password, Email,ID from Users where userName=?", userName)
	err := row.Scan(&user.Username, &user.FirstName, &user.SecondName, &user.Password, &user.Email, &user.ID)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (m MysqlRepo) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	var users []*models.User
	row, err := m.conn.Queryx("select Username,FirstName,SecondName,Password, Email,ID from Users")
	if err != nil {
		return nil, err
	}
	for row.Next() {
		var user models.User
		row.Scan(&user.Username, &user.FirstName, &user.SecondName, &user.Password, &user.Email, &user.ID)

		users = append(users, &user)
	}
	return users, nil
}

func NewMysqlReqo(db *sqlx.DB) DbRepository {
	return &MysqlRepo{
		db,
	}

}
