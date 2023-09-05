package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/username/reponame/exercise/interface/data"
)

type Conn struct {
	db *sql.DB
}

func NewConn(db *sql.DB) *Conn {
	return &Conn{db: db}
}

func (p *Conn) Create(ctx context.Context, usr data.User) error {
	fmt.Println("adding to postgres", usr)
	return nil
}
func (p *Conn) Update(ctx context.Context, usr data.User) error {
	fmt.Println("updating in postgres", usr)
	return nil
}
func (p *Conn) Delete(usr data.User) error {
	fmt.Println("deleting in postgres", usr)
	return nil
}
