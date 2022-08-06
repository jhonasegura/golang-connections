package postgre

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "first_db"
)

func main() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	defer db.Close()

	insertStmt := `insert into "Employee"("Name", "EmpId") values('Robit', 21)`
	_, e := db.Exec(insertStmt)
	CheckError(e)

	insertDynStmt := `insert into "Employee"("Name", "EmpId") values($1, $2)`
	_, e = db.Exec(insertDynStmt, "krish", 03)
	CheckError(e)

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
