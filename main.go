package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/kislerdm/aws-amplify-deploy-zip/amplifydeployzip"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{ProviderFunc: amplifydeployzip.Provider})
}
