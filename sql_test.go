package databases

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestInsert(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	_, err := db.ExecContext(ctx, "INSERT INTO customer(id, name, email, balance, rating, birth_date, married) VALUES('third', 'Drama', 'null', 156984, 80.0, '1992-02-01', true)")
	if err != nil {
		panic(err)
	}

	fmt.Println("Success Insert New Customer")
}

func TestSelect(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	rows, err := db.QueryContext(ctx, "SELECT * FROM customer")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birth_date, created_at time.Time
		var married bool
		err := rows.Scan(&id, &name, &email, &balance, &rating, &created_at, &birth_date, &married)
		if err != nil {
			panic(err)
		}

		fmt.Println("ID : ", id, "Name: ", name, "Email: ", email, "Balance: ", balance, "Rating: ", rating, "Birth Date: ", birth_date, "Married: ", married, "Created At: ", created_at)
	}

	defer rows.Close()
}

func TestSQLInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	uname := "admin'; #"
	pass := "salah"
	ctx := context.Background()
	rows, err := db.QueryContext(ctx, "SELECT username FROM user WHERE username = '"+uname+"' AND password = '"+pass+"' LIMIT 1")
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login", username)
	} else {
		fmt.Println("Error Login")
	}
}

func TestSQLInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	uname := "admin'; #"
	pass := "salah"
	ctx := context.Background()
	rows, err := db.QueryContext(ctx, "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1", uname, pass)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login", username)
	} else {
		fmt.Println("Error Login")
	}
}

func TestInsertSqlParams(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	_, err := db.ExecContext(ctx, "INSERT INTO user(username, password) VALUES(?,?)", "member", "member")
	if err != nil {
		panic(err)
	}

	fmt.Println("Success Insert New User")
}

func TestLastInsertID(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	email := "go@gmail.com"
	comment := "BLEH BLEH BLEH"
	result, err := db.ExecContext(ctx, "INSERT INTO comments(email, comment) VALUES(?,?)", email, comment)
	if err != nil {
		panic(err)
	}

	insertID, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Success Insert New Comment", insertID)
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	stmt, err := db.PrepareContext(ctx, "INSERT INTO comments(email, comment) VALUES(?,?)")
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	for i := 0; i < 10; i++ {
		email := "Dimas" + strconv.Itoa(i) + "@gmail.com"
		comment := "Komentar Ke-" + strconv.Itoa(i)

		res, err := stmt.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := res.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Last ID-", id)
	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	for i := 0; i < 10; i++ {
		email := "Dimas" + strconv.Itoa(i) + "@gmail.com"
		comment := "Komentar Ke-" + strconv.Itoa(i)

		res, err := tx.ExecContext(ctx, "INSERT INTO comments(email, comment) VALUES(?,?)", email, comment)
		if err != nil {
			panic(err)
		}

		id, err := res.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Last ID-", id)
	}

	// To Commit the transaction
	err = tx.Commit()
	// To Cancel Commit the transaction
	//err = tx.Rollback()
	if err != nil {
		panic(err)
	}
}
