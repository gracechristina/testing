package testing

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"

    _"github.com/lib/pq"
)

const(
    DB_USER = "gc160606"
    DB_PASSWORD = "Ayamgoreng123"
    DB_NAME = "kero"
    TCP = "192.168.100.126"
)

func main() {
    db := NewDB()
    log.Println("Listening on :8080")
    http.ListenAndServe(":8080", ShowBooks(db))
}

func ShowBooks(db *sql.DB) http.Handler {
    return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
        var district_name, postcode string
        var district_id int
        err := db.QueryRow("SELECT district_id, district_name,postcode FROM kk_postcodes").Scan(&district_id, &district_name,&postcode)
        if err != nil {
            panic(err)
        }

        fmt.Fprintf(rw, "The information is %d and %s and %s", district_id, district_name, postcode)
    })
}

func NewDB() *sql.DB {
    dbinfo := fmt.Sprintf("user = %s password= %s dbname=%s sslmode=disable host=192.168.100.126 port=5432", 
    DB_USER,DB_PASSWORD,DB_NAME)
    db, err := sql.Open("postgres", dbinfo)
    if err != nil {
        panic(err)
    }

    /*_, err = db.Exec("create table if not exists books(title text, author text)")
    if err != nil {
        panic(err)
    }*/

    return db
}
