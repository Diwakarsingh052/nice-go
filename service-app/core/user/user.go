package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
	"service-app/auth"
	"strconv"
	"time"
)

type Service struct {
	db *pgxpool.Pool
}

func NewService(db *pgxpool.Pool) (*Service, error) { //db *sql.DB
	if db == nil {
		return nil, errors.New("please provide a valid connection")
	}
	s := &Service{db: db}
	return s, nil
}

func (s *Service) CreateUser(ctx context.Context, nu NewUser, now time.Time) (User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(nu.Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, fmt.Errorf("generating password hash %w", err)
	}

	u := User{
		Name:         nu.Name,
		Email:        nu.Email,
		Roles:        nu.Roles,
		PasswordHash: hash,
		DateCreated:  now,
		DateUpdated:  now,
	}

	const q = `INSERT INTO users
		(name, email, password_hash, roles, date_created, date_updated)
		VALUES ( $1, $2, $3, $4, $5, $6)
		Returning id`
	/*
		//Query acquires a connection and executes a query that is expected to return multiple rows
		s.db.Query()

		//QueryRow acquires a connection and executes a query that is expected to return at **most** one row
		s.db.QueryRow()

		//exec the query doesn't return anything
		s.db.Exec()

	*/

	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	row := s.db.QueryRow(ctx, q, u.Name, u.Email, u.PasswordHash, u.Roles, u.DateCreated, u.DateUpdated)
	var id int
	err = row.Scan(&id)
	if err != nil {
		return User{}, fmt.Errorf("inserting user %w", err)
	}
	u.ID = strconv.Itoa(id)
	return u, nil
}

func (s *Service) Authenticate(ctx context.Context, email, password string, now time.Time) (auth.Claims, error) {

	//this query is used to check whether user exist in the db or not
	const q = `SELECT id,name,email,roles,password_hash FROM users WHERE email = $1`
	var u User
	row := s.db.QueryRow(ctx, q, email)
	err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Roles, &u.PasswordHash)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return auth.Claims{}, errors.New("authentication failed")
		}
		return auth.Claims{}, err
	}

	err = bcrypt.CompareHashAndPassword(u.PasswordHash, []byte(password))
	if err != nil {
		return auth.Claims{}, errors.New("authentication failed")
	}
	claims := auth.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "service project",
			Subject:   u.ID,
			Audience:  jwt.ClaimStrings{"students"},
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
		Roles: u.Roles,
	}

	return claims, nil
}
