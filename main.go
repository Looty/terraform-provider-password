package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"terraform-provider-password/password"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return password.Provider()
		},
	})

	//var debug bool
	//
	//flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	//flag.Parse()
	//
	//opts := &plugin.ServeOpts{
	//	ProviderFunc: password.Provider,
	//}
	//
	//if debug {
	//	err := plugin.Debug(context.Background(), "hashicorp.com/edu/password", opts)
	//	if err != nil {
	//		log.Fatal(err.Error())
	//	}
	//}
}
