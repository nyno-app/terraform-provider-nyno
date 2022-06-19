package nyno

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Config struct {
	api_endpoint string
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
		},
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	config := Config{
		api_endpoint: d.Get("api_endpoint").(string),
	}

	return config, nil
}
