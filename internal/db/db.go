package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/siddarthtv/golang_assignment_kelp/models"
	_ "modernc.org/sqlite"
)

type DBConn struct {
	db *sqlx.DB
}

func (d *DBConn) CompanyFinancials(companyID int) (interface{}, error) {
	var resp = new(models.Financials)
	err := d.db.Get(resp, "SELECT * FROM financials WHERE companyID = ?", companyID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DBConn) CompanySales(companyID int) (interface{}, error) {
	var resp = new(models.Sales)
	err := d.db.Get(resp, "SELECT * FROM sales WHERE companyID = ?", companyID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DBConn) CompanyEmployeeStats(companyID int) (interface{}, error) {
	var resp = new(models.EmpStats)
	err := d.db.Get(resp, "SELECT * FROM employee_stats WHERE companyID = ?", companyID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func InitDB() *DBConn {
	db, err := sqlx.Connect("sqlite", "./sample.db")
	if err != nil {
		log.Fatalf("Error connecting to the database: %s \n", err.Error())
	}
	dbconn := new(DBConn)
	dbconn.db = db
	return dbconn
}
