package gopherbiswa

import (
	"context"
	"database/sql"
	"log"
)

var (
	id      int
	name    string
	address string
	age     int
	city    string
	ctx     context.Context
	test 	int
)

//CreateData creates a record in the database and return s a ID.
func CreateData(db *sql.DB, req Information) (Information, error) {
	err := db.QueryRow("EXEC terrform.dbo.terraform_insert @id=?, @name=?, @e_address=?, @age=?, @city=?", req.ID, req.Name, req.Address, req.Age, req.City).Scan(&id, &name, &address, &age, &city)
	if err != nil {
		log.Fatal(err)
	}
	resp := Information{
		ID:      id,
		Name:    name,
		Address: address,
		Age:     age,
		City:    city,
	}
	return resp, err
}

//DeleteData deletes the data from the database.
func DeleteData(db *sql.DB, req Information) {
	_, err := db.Query("EXEC terrform.dbo.terraform_delete @id=? ",req.ID)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Record has been deleted")
}

//UpdateData updates the record in the database
func UpdateData(db *sql.DB, req Information) (Information, error) {
	err := db.QueryRow("EXEC terrform.dbo.terraform_update @id=?, @name=?,@e_address=?,@age=?,@city=?", req.ID, req.Name, req.Address, req.Age, req.City).Scan(&id, &name, &address, &age, &city)
	if err != nil {
		log.Fatal(err)
	}
	resp := Information{
		ID:      id,
		Name:    name,
		Address: address,
		Age:     age,
		City:    city,
	}
	return resp, err
}

//ReadData deletes the data from the database.
func ReadData(db *sql.DB, req Information) ([]Information, error) {
	rows, err := db.Query("Select * from terrform.dbo.terraform_sql")
	if err != nil {
		log.Fatal(err)
	}
	var infoArr []Information

	for rows.Next() {
		err := rows.Scan(&id, &name, &address, &age, &city)
		if err != nil {
			log.Fatal(err)
		}
		i := Information{
			ID:      id,
			Name:    name,
			Address: address,
			Age:     age,
			City:    city,
		}
		infoArr = append(infoArr, i)
		//log.Println(id, name, address, age, city)
	}
	return infoArr, err
}
