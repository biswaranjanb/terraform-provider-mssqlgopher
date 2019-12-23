package main

import(
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
	"github.com/terraform-provider-gopherbiswa/gopherbiswa"
)

func main(){
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider{
			return gopherbiswa.Provider()
		},
	})
}