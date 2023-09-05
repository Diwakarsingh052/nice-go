package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/username/reponame/exercise/interface/data"
)

type Conn struct {
	db *sql.DB
}

// NewConn setups the db connection
func NewConn(db *sql.DB) *Conn {
	return &Conn{db: db}
}

func (p *Conn) Create(ctx context.Context, usr data.User) error {
	fmt.Println("adding to mysql", usr)
	return nil
}
func (p *Conn) Update(ctx context.Context, usr data.User) error {
	fmt.Println("updating in mysql", usr)
	return nil
}
func (p *Conn) Delete(ctx context.Context, usr data.User) error {
	fmt.Println("deleting in mysql", usr)
	return nil
}
