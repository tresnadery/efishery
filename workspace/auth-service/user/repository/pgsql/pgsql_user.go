package pgsql

import(	
	"log"
	"time"
	"context"
	"database/sql"		
	"auth-service/domain"
)

type PgsqlUserRepository struct{
	Conn *sql.DB
}

func NewPgsqlUserRepository(conn *sql.DB) *PgsqlUserRepository{
	return &PgsqlUserRepository{
		Conn: conn,
	}
}
func (p *PgsqlUserRepository) GetByPhoneNumber(phoneNumber string)(*domain.User, error){
	var user domain.User
	query := `SELECT users.id, phone_number, users.name, password, roles.name, users.created_at FROM users INNER JOIN roles ON (users.role_id = roles.id) WHERE phone_number = $1`	
	log.Println(query)
	log.Println("Phone number : " + phoneNumber)
	row := p.Conn.QueryRow(query, phoneNumber)
	err := row.Scan(&user.ID, &user.PhoneNumber, &user.Name, &user.Password, &user.RoleName, &user.CreatedAt)
	switch err{
	case sql.ErrNoRows:
		return nil, domain.ErrNotFound
	case nil:
		return &user, nil
	default:
		return nil, domain.ErrInternalServerError
	}
}
func (p *PgsqlUserRepository) GetByNameANDPhoneNumber(name, phoneNumber string)(*domain.User, error){		
	var user domain.User
	query := `SELECT users.id, phone_number, users.name, roles.name FROM users INNER JOIN roles ON (users.role_id = roles.id) WHERE LOWER(users.name) LIKE LOWER($1) AND phone_number = $2;`
	row := p.Conn.QueryRow(query, name, phoneNumber)
	err := row.Scan(&user.ID, &user.PhoneNumber, &user.Name, &user.RoleName)
	switch err{
	case sql.ErrNoRows:
		return nil, domain.ErrNotFound
	case nil:
		return &user, nil
	default:
		return nil, domain.ErrInternalServerError
	}	
}

func (p *PgsqlUserRepository) Store(ctx context.Context, payload *domain.User)(error){
	query := `INSERT INTO users(phone_number, name, role_id, password, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6)`
	stmt, err := p.Conn.PrepareContext(ctx, query)
	if err != nil{		
		return err
	}
	_, err = stmt.ExecContext(ctx, payload.PhoneNumber, payload.Name, payload.RoleID, payload.Password, time.Now(), time.Now())
	if err != nil{		
		return err
	}
	return nil
}

func (p *PgsqlUserRepository) UpdatePassword(ctx context.Context, userID string, password string)(error){
	log.Println(userID, password)
	query := `UPDATE users SET password = $1 WHERE id = $2 `
	stmt, err := p.Conn.PrepareContext(ctx, query)
	if err != nil{		
		return err
	}
	_, err = stmt.ExecContext(ctx, password, userID)
	if err != nil{		
		return err
	}
	return nil
}