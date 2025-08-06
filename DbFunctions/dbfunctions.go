package dbfunctions

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

type Merch struct {
	Name     string  `json: "name"`
	Price    float64 `json: "price"`
	Type     string  `json: "type"`
	Size     string  `json: "size"`
	Quantity int64   `json: "quantity"`
	Color    string  `json: "color"`
}

// type Merch struct {
// 	Name     string
// 	Type     string
// 	Price    float64
// 	Size     string
// 	Quantity int64
// }

var db *sql.DB

func DBConnect() (bool, error) {
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DBUSER")
	cfg.Passwd = os.Getenv("DBPASS")
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "merchandise"

	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return false, fmt.Errorf("DBConnect: error at sql.Open %v", err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		return false, fmt.Errorf("DBConnect: error at db.Ping %v", pingErr)
	}
	return true, nil
}

func AddMerchToDb(merch Merch) (int64, error) {
	//Error handler
	fail := func(err error) (int64, error) {
		return 0, fmt.Errorf("error in SaveMerchToDb: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := db.BeginTx(ctx, nil) //Using transaction since we have 2 inserts
	if err != nil {
		return fail(err)
	}
	defer tx.Rollback() //incase one insert fails

	//Get merchId if available
	row := tx.QueryRowContext(ctx, "Select id from merch where name=? and type=?", merch.Name, merch.Type)

	var merchId int64
	merchId = 0
	if err := row.Scan(&merchId); err != nil {
		if err == sql.ErrNoRows {
			//Add merch if not available
			result, err := tx.ExecContext(ctx, "Insert into merch (name, type, price) values (?,?,?)", merch.Name, merch.Type, merch.Price)
			if err != nil {
				return fail(err)
			}
			merchId, err = result.LastInsertId()
			if err != nil {
				return fail(err)
			}
		} else {
			return fail(err)
		}
	}

	var stockID int64
	if merchId != 0 {
		//Add stock
		result, err := tx.ExecContext(ctx, "Insert into stock (merch_fkid, size, quantity) values (?,?,?)", merchId, merch.Size, merch.Quantity)
		if err != nil {
			return fail(err)
		}

		stockID, err = result.LastInsertId()
		if err != nil {
			return fail(err)
		}

	}

	if err = tx.Commit(); err != nil {
		return fail(err)
	}

	return stockID, nil
}

func GetAllMerchWithQuantity() ([]Merch, error) {
	var allmerch []Merch

	//Error handler
	fail := func(err error) ([]Merch, error) {
		return nil, fmt.Errorf("error in GetAllMerchWithQuantity: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := db.BeginTx(ctx, nil) //Using transaction since we have 2 inserts
	if err != nil {
		return fail(err)
	}
	defer tx.Rollback() //incase one insert fails

	rows, err := tx.QueryContext(ctx, "select m.name, m.type, m.price, m.color, sum(s.quantity) as quantity from merch m inner join stock s on m.id = s.merch_fkid group by merch_fkid")
	if err != nil {
		return fail(err)
	}

	if err := rows.Err(); err != nil {
		return fail(fmt.Errorf("getBoth: Error in Overall Query %v", err))
	}

	defer rows.Close()

	for rows.Next() {
		var merch Merch
		if err := rows.Scan(&merch.Name, &merch.Type, &merch.Price, &merch.Color, &merch.Quantity); err != nil {
			if err == sql.ErrNoRows {
				return fail(fmt.Errorf("getBoth(): No Merchandise Found %v", err))
			}
			return fail(err)
		}

		allmerch = append(allmerch, merch)
	}

	return allmerch, nil
}
