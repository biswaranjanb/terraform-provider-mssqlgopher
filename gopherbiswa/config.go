package gopherbiswa

import (
	"database/sql"
	"fmt"

  //"fmt"
)
var server = "INLT-HYY3ZY2\\SQL2014"
var user = "sa"
var password = "Sa@123"

// CheckDB checks if the database "strDBName" exists on the MSSQL database engine.
func CheckDB(db *sql.DB, strDBName string) (bool, error) {
	// Does the database exist?
	result, err := db.Query("SELECT db_id('" + strDBName + "')")
	defer result.Close()

	if err != nil {
		return false, err
	}

	for result.Next() {
		var s sql.NullString
		err := result.Scan(&s)
		if err != nil {
			return false, err
		}
		// Check result
		if s.Valid {
			return true, nil
		}
	}

	// This return() should never be hit...
	return false, err
}

//CreateConnection creates the connection to the database
func CreateConnection()(db *sql.DB, err error) {
	db, err1 := sql.Open("mssql", "server="+server+";Initial Catalog=terrform;user id="+user+";password="+password+";encrypt=disable;")
	if err1 != nil {
		fmt.Println("yahan")
		fmt.Println(err1)
	}
	//defer db.Close()
	boolDBExist, err := CheckDB(db, "terrform")
	fmt.Println(boolDBExist)
	if !boolDBExist {
		fmt.Println("Not there")
	}
	return db,err
}
