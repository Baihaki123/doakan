package main

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID            int       `db:"id"`
	Nama          string    `db:"nama"`
	Pekerjaan     string    `db:"pekerjaan"`
	Email         string    `db:"email"`
	Password      string    `db:"password"`
	Image_profile string    `db:"image_profile"`
	Role          string    `db:"role"`
	Created_At    time.Time `db:"created_At"`
	Updated_At    time.Time `db:"updated_At"`
}

type Transaction struct {
	ID          int       `db:"id"`
	Donation_ID int       `db:"donation_id"`
	User_ID     int       `db:"user_id"`
	Amount      int       `db:"amount"`
	Status      string    `db:"status"`
	Va_Number   string    `db:"va_number"`
	Bank        string    `db:"bank"`
	Payment_url string    `db:"payment_url"`
	Created_At  time.Time `db:"created_at"`
	Updated_At  time.Time `db:"updated_at"`
}

type Donation struct {
	ID                int       `db:"id"`
	User_ID           int       `db:"user_id"`
	Nama              string    `db:"nama"`
	Deskripsi_Singkat string    `db:"deskripsi_singkat"`
	Deskripsi         string    `db:"deskripsi"`
	Perks             string    `db:"perks"`
	Donatur_Count     int       `db:"donatur_count"`
	Goal_Amount       int       `db:"goal_amount"`
	Current_Amount    int       `db:"current_amount"`
	Slug              string    `db:"slug"`
	Created_At        time.Time `db:"created_at"`
	Updated_At        time.Time `db:"updated_at"`
}

type DonationImage struct {
	ID          int       `db:"id"`
	Donation_ID int       `db:"donation_id"`
	FileName    string    `db:"filename"`
	Is_Primary  string    `db:"is_primary"`
	Created_At  time.Time `db:"created_at"`
	Updated_At  time.Time `db:"updated_at"`
}

// Migrate digunakan untuk melakukan migrasi database dengan data yang dibutuhkan
func Migrate() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "../doakan.db")
	if err != nil {
		panic(err)
	}

	sqlStmt := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		nama VARCHAR(128),
		pekerjaan VARCHAR(128),
		email VARCHAR(128),
		password VARCHAR(128),
		image_profile VARCHAR(128),
		role VARCHAR(128),
		created_at DATETIME,
		updated_at DATETIME
	);`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	sqlStmt = `CREATE TABLE IF NOT EXISTS transactions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		donation_id INTEGER,
		user_id INTEGER,
		amount INTEGER,
		status VARCHAR(128),
		va_number VARCHAR(128),
		bank VARCHAR(128),
		payment_url VARCHAR(128),
		created_at DATETIME,
		updated_at DATETIME
	);`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	sqlStmt = `CREATE TABLE IF NOT EXISTS donations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		nama VARCHAR(128),
		deskripsi_singkat VARCHAR(128),
		deskripsi TEXT,
		perks VARCHAR(128),
		donatur_count INTEGER,
		goal_amount INTEGER,
		current_amount INTEGER,
		slug VARCHAR(128),
		created_at DATETIME,
		updated_at DATETIME
	);`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	sqlStmt = `CREATE TABLE IF NOT EXISTS donation_images (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		donation_id INTEGER,
		filename VARCHAR(128),
		is_primary VARCHAR(128),
		created_at DATETIME,
		updated_at DATETIME
	);`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Rollback(db *sql.DB) {
	sqlStmt := `DROP TABLE users;`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		panic(err)
	}

	sqlStmt = `DROP TABLE transactions;`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		panic(err)
	}

	sqlStmt = `DROP TABLE donations;`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		panic(err)
	}

	sqlStmt = `DROP TABLE donation_images;`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		panic(err)
	}
}

// Jalankan main untuk melakukan migrasi database
func main() {
	_, err := Migrate()
	if err != nil {
		panic(err)
	}

	// Use Rollback() // to rollback the changes
	// Rollback(db)
}
