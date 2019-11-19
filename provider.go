package main

import 	"github.com/hashicorp/terraform/helper/schema"


//Provider returns the schema Provider
func Provider() *schema.Provider {
	return &schema.Provider{
			ResourcesMap: map[string]*schema.Resource{
					"csv_server": resourceServer(),
			},
	}
}