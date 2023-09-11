package inventory

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"strconv"
	"time"
)

type Service struct {
	db *pgxpool.Pool
}

func NewService(db *pgxpool.Pool) (*Service, error) {
	if db == nil {
		return nil, errors.New("please provide a valid connection")
	}
	s := &Service{db: db}
	return s, nil
}

func (s *Service) CreateInventory(ctx context.Context, ni NewShirtInventory, userId string,
	now time.Time) (ShirtInventory, error) {

	inv := ShirtInventory{
		UserId:      userId,
		ItemName:    ni.ItemName,
		Quantity:    ni.Quantity,
		DateCreated: now,
		DateUpdated: now,
	}

	//inserting the core in db for specific core
	const q = `INSERT INTO inventory
		(user_id, item_name, quantity, date_created, date_updated)
		VALUES ( $1, $2, $3, $4, $5)
		Returning shirt_id`

	var shirtId int
	//exec the query // QueryRowContext is used as we are expecting one row back in the result
	row := s.db.QueryRow(ctx, q, userId, inv.ItemName, inv.Quantity, inv.DateCreated, inv.DateUpdated)
	err := row.Scan(&shirtId)

	if err != nil {
		return ShirtInventory{}, fmt.Errorf("inserting inventory %w", err)
	}
	inv.ShirtID = strconv.Itoa(shirtId)
	return inv, nil

}

func (s *Service) ViewAll(ctx context.Context, userId string) ([]ShirtInventory, error) {

	var inv = make([]ShirtInventory, 0, 10)

	const q = `Select shirt_id, user_id, item_name,quantity,date_created,
				date_updated FROM inventory where user_id = $1`

	//QueryContext is used when we are expecting multiple rows to be returned
	rows, err := s.db.Query(ctx, q, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// rows.Next() returns true if more rows are left to read otherwise false which will stop the loop
	for rows.Next() {
		var newInv ShirtInventory
		err = rows.Scan(&newInv.ShirtID, &newInv.UserId, &newInv.ItemName, &newInv.Quantity, &newInv.DateCreated,
			&newInv.DateUpdated)
		if err != nil {
			return nil, err
		}
		//appending the struct into the slice to maintain a whole list
		inv = append(inv, newInv)
	}
	return inv, nil
}
