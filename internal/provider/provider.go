package provider

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/kislerdm/terraform-provider-awsamplify/client"
)

var c *client.Client

func init() {
	schema.DescriptionKind = schema.StringMarkdown
	rand.Seed(time.Now().Unix())
}

func New() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"region": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "AWS region",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"awsamplify_deploy_zip": resourceDeploy(),
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
		Description:   `Resource ` + "`awsamplify_deploy_zip`" + ` deploys the web application into AWS Amplify from zip archive.`,
		SchemaVersion: 1,
		CreateContext: resourceCreate,
		UpdateContext: resourceUpdate,
		ReadContext:   resourceRead,
		DeleteContext: resourceDelete,
		Schema: map[string]*schema.Schema{
			"app_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Amplify app ID",
			},
			"branch_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Amplify deployment branch",
			},
			"path": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Path to zip archive with the app build",
			},
			"triggers": {
				Description: "A map of arbitrary strings that, when changed, will force re-deployment.",
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
			},
			"id": {
				Description: "This is set to a random value at create time.",
				Computed:    true,
				Type:        schema.TypeString,
			},
		},
	}
}

func resourceCreate(ctx context.Context, d *schema.ResourceData, v interface{}) diag.Diagnostics {
	if err := c.DeployZip(ctx,
		d.Get("app_id").(string),
		d.Get("branch_name").(string),
		d.Get("path").(string),
	); err != nil {
		return diag.Errorf("Deploy error: %s", err)
	}
	d.SetId(fmt.Sprintf("%d", rand.Int()))
	return diag.Diagnostics{}
}

func resourceUpdate(ctx context.Context, d *schema.ResourceData, v interface{}) diag.Diagnostics {
	return nil
}

func resourceRead(ctx context.Context, d *schema.ResourceData, v interface{}) diag.Diagnostics {
	return nil
}

func resourceDelete(ctx context.Context, d *schema.ResourceData, v interface{}) diag.Diagnostics {
	d.SetId("")
	return nil
}
