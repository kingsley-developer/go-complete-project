package mysqlconn
import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// ...


db, err := sql.Open("mysql", "user:password@/dbname")
if err != nil {
	panic(err)
}
// See "Important settings" section.
db.SetConnMaxLifetime(time.Minute * 3)
db.SetMaxOpenConns(10)
db.SetMaxIdleConns(10)
/*
type MySQL struct {
	Name string
}
*/
