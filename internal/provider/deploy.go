package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/kislerdm/terraform-provider-aws_amplify_deploy_zip/client"
)

var c *client.Client

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"region": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "AWS region",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"aws_amplify_deploy_zip": resourceDeploy(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	c = client.New(d.Get("region").(string))
	return c, diag.Diagnostics{}
}

func resourceDeploy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: create,
		Schema: map[string]*schema.Schema{
			"app_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Amplify app ID",
			},
			"branch": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Amplify deployment branch",
			},
			"path": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Path to zip archive with the app build",
			},
		},
	}
}

func create(ctx context.Context, d *schema.ResourceData, v interface{}) diag.Diagnostics {
	diags := diag.Diagnostics{}
	if err := c.DeployZip(ctx,
		d.Get("app_id").(string),
		d.Get("branch").(string),
		d.Get("path").(string),
	); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Deploy error",
			Detail:   fmt.Sprintf("Deploy error: %s", err),
		})
	}
	return diags
}
