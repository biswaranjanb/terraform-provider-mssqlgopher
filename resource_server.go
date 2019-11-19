package main

import (
	// "encoding/csv"
	"github.com/hashicorp/terraform/helper/schema"
	"os"
	"strconv"
	"sort"
	// "strings"
	//"bytes"
)

//Information models the information struct
type Information struct {
	ID   string `json:"line_id"`
	Name string `json:"name"`
	Zip  string `json:"zip"`
}

func resourceServer() *schema.Resource {
	return &schema.Resource{ // This defines the data schema and CRUD operations of the resource. This is the only thing we require to create a resource.
		Create: resourceCreate,
		Read:   resourceRead,
		Update: resourceUpdate,
		Delete: resourceDelete,
		Schema: map[string]*schema.Schema{
			"file": &schema.Schema{ // TForm's schema automatically enforces validation and type casting.
				Type:     schema.TypeString,
				Required: true,
			},
			"fields": {
				Type: schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"line_id":{ // TForm's schema automatically enforces validation and type csting.
							Type:     schema.TypeString,
							Required: true,
						},
						"name": { // TForm's schema automatically enforces validation and type castng.
							Type:     schema.TypeString,
							Required: true,
						},
						"zip": { // TForm's schema automatically enforces validation and type casting.
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
	fieldsData := make([]Information,0,0)
	for _,field := range fieldsList{
		fieldMap := field.(map[string]interface{})		
		fieldsData = append(fieldsData, Information{
			Name: fieldMap["name"].(string),
			Zip:  fieldMap["zip"].(string),
			ID:   fieldMap["line_id"].(string),	
		})
	}
	fileName := d.Get("file").(string) // put the code to create a csv file.
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}

	sort.Strings(fieldsData)
	
	for _,emp := range fieldsData{		
		if len(emp.ID) == 0 && len(emp.Name) == 0 && len(emp.Zip) == 0{
			_,err = file.WriteString("")
		}else{
			_,err = file.WriteString(emp.ID + ", " + emp.Name + ", " + emp.Zip + "\n")
		}
		
		if err != nil {
			file.Close()
			return err		
		}
	}
	err = file.Close()
	if err != nil {
		d.SetId(strconv.Itoa(d.Get("line_id").(int)))
		d.Set("fields", fieldsData)
		return nil
	}
	return nil
}

//resourceRead models the metho for reading the resource
func resourceRead(d *schema.ResourceData, m interface{}) error {	
	
	 return nil
}

//resourceUpdate models the method for updating the resource
func resourceUpdate(d *schema.ResourceData, m interface{}) error {	
	return resourceCreate(d,m)
}

//resourceDelete models the method for deleting the resource.
func resourceDelete(d *schema.ResourceData, m interface{}) error {	
	return nil
}

