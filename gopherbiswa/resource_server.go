package gopherbiswa

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
)

var server1 = "INLT-HYY3ZY2\\SQL2014"
var user1 = "sa"
var password1 = "Sa@123"

//PingServer pings the database server
func PingServer(db *sql.DB) string {

	err := db.Ping()
	if err != nil {
		return ("From Ping() Attempt: " + err.Error())
	}

	return ("Database Ping Worked...")

}

func resourceServer() *schema.Resource {
	return &schema.Resource{ // This defines the data schema and CRUD operations of the resource. This is the only thing we require to create a resource.
		Create: resourceCreate,
		Read:   resourceRead,
		Update: resourceUpdate,
		Delete: resourceDelete,
		Schema: map[string]*schema.Schema{
			"fields": {
				Type: schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"address": { // TForm's schema automatically enforces validation and type castng.
							Type:     schema.TypeString,
							Required: true,
						},
						"age": { // TForm's schema automatically enforces validation and type casting.
							Type:     schema.TypeInt,
							Required: true,
						},
						"city": { // TForm's schema automatically enforces validation and type casting.
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
				Required: true,
			},
		},
	}
}

//resourceCreate models the method or resource creation
func resourceCreate(d *schema.ResourceData, m interface{}) error {
	fields := d.Get("fields").(*schema.Set)
	fieldsList := fields.List()
	fieldsData := make([]Information, 0, 0)
	for _, field := range fieldsList {
		fieldMap := field.(map[string]interface{})
		fieldsData = append(fieldsData, Information{
			Name:    fieldMap["name"].(string),
			Address: fieldMap["address"].(string),
			Age:     fieldMap["age"].(int),
			City:    fieldMap["city"].(string),
		})
	}
	db, err1 := sql.Open("mssql", "server="+server1+";Initial Catalog=terrform;user id="+user1+";password="+password1+";encrypt=disable;")
	if err1 != nil {
		fmt.Println("yahan")
		fmt.Println(err1)
	}
	defer db.Close()
	resp := Information{}
	for _, req := range fieldsData {
		fmt.Println("There")
		err := db.QueryRow("EXEC terrform.dbo.terraform_insert @id=?, @name=?, @e_address=?, @age=?, @city=?", req.ID, req.Name, req.Address, req.Age, req.City).Scan(&id, &name, &address, &age, &city)
		resp = Information{
			ID:      id,
			Name:    name,
			Address: address,
			Age:     age,
			City:    city,
		}
		if err != nil {
			log.Fatal(err)
			return err
		}
	}
	d.SetId(strconv.Itoa(resp.ID))
	d.Set("fields", fieldsData)
	return nil
}

//resourceRead models the metho for reading the resource
func resourceRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

//resourceUpdate models the method for updating the resource
func resourceUpdate(d *schema.ResourceData, m interface{}) error {
	fields := d.Get("fields").(*schema.Set)
	fieldsList := fields.List()
	fieldsData := make([]Information, 0, 0)
	for _, field := range fieldsList {
		fieldMap := field.(map[string]interface{})
		fieldsData = append(fieldsData, Information{
			Name:    fieldMap["name"].(string),
			Address: fieldMap["address"].(string),
			Age:     fieldMap["age"].(int),
			City:    fieldMap["city"].(string),
		})
	}
	db, err1 := sql.Open("mssql", "server="+server1+";Initial Catalog=terrform;user id="+user1+";password="+password1+";encrypt=disable;")
	if err1 != nil {
		fmt.Println("yahan")
		fmt.Println(err1)
	}
	defer db.Close()
	resp := Information{}
	for _, req := range fieldsData {
		fmt.Println("There")
		err := db.QueryRow("EXEC terrform.dbo.terraform_update @id=?, @name=?, @e_address=?, @age=?, @city=?", d.Id(), req.Name, req.Address, req.Age, req.City).Scan(&id, &name, &address, &age, &city)
		resp = Information{
			ID:      id,
			Name:    name,
			Address: address,
			Age:     age,
			City:    city,
		}
		if err != nil {
			log.Fatal(err)
			return err
		}
	}
	d.SetId(strconv.Itoa(resp.ID))
	d.Set("fields", fieldsData)
	return nil
}

//resourceDelete models the method for deleting the resource.
func resourceDelete(d *schema.ResourceData, m interface{}) error {
	
	db, err1 := sql.Open("mssql", "server="+server1+";Initial Catalog=terrform;user id="+user1+";password="+password1+";encrypt=disable;")
	if err1 != nil {
		fmt.Println(err1)
	}
	defer db.Close()	
	// _, err := db.Query("EXEC terrform.dbo.terraform_delete @id=?,", d.Id())
	err := db.QueryRow("EXEC terrform.dbo.terraform_delete @id=?", d.Id()).Scan(&test)
	
	fmt.Println(test)
	if err != nil {
		fmt.Println(err)
	}
	d.SetId("")
	return nil
}
