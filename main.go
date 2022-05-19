//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs
package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/snakeice/terraform-provider-hash-sum/hashsum"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return hashsum.Provider()
		},
	})
}
