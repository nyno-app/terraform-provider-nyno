package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Config struct {
	api_endpoint  string
	session_token string
}

func Provider() *schema.Provider {
	return &schema.Provider{
		ConfigureContextFunc: providerConfigure,
		ResourcesMap: map[string]*schema.Resource{
			"nyno_template": resourceTemplate(),
			"nyno_role":     resourceRole(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"nyno_template":   dataSourceTemplate(),
			"nyno_role":       dataSourceRole(),
			"nyno_repository": dataSourceRepository(),
		},
		Schema: map[string]*schema.Schema{
			"api_endpoint": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("NYNO_API_ENDPOINT", "https://nyno.io/api"),
				Description: "The URL to use for Nyno API",
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("NYNO_USERNAME", nil),
				Description: "Username to use for Nyno API",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("NYNO_PASSWORD", nil),
				Description: "Password to use for Nyno API",
			},
			"organization": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("NYNO_ORGANIZATION", nil),
				Description: "Organization to use for Nyno API",
			},
		},
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	username := d.Get("username").(string)
	password := d.Get("password").(string)
	organization := d.Get("organization").(string)
	api_endpoint := d.Get("api_endpoint").(string)

	session_token := getSessionToken(api_endpoint, username, password, organization)

	config := Config{
		api_endpoint:  d.Get("api_endpoint").(string),
		session_token: session_token,
	}

	return config, nil
}
