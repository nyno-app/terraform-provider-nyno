package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceRole() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceRoleRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name":                             {Type: schema.TypeString, Computed: true},
			"create_credentials":               {Type: schema.TypeBool, Computed: true},
			"get_credentials":                  {Type: schema.TypeBool, Computed: true},
			"update_credentials":               {Type: schema.TypeBool, Computed: true},
			"delete_credentials":               {Type: schema.TypeBool, Computed: true},
			"create_repository":                {Type: schema.TypeBool, Computed: true},
			"get_repository":                   {Type: schema.TypeBool, Computed: true},
			"update_repository":                {Type: schema.TypeBool, Computed: true},
			"delete_repository":                {Type: schema.TypeBool, Computed: true},
			"get_user":                         {Type: schema.TypeBool, Computed: true},
			"update_user":                      {Type: schema.TypeBool, Computed: true},
			"create_user":                      {Type: schema.TypeBool, Computed: true},
			"get_role":                         {Type: schema.TypeBool, Computed: true},
			"update_role":                      {Type: schema.TypeBool, Computed: true},
			"create_role":                      {Type: schema.TypeBool, Computed: true},
			"delete_role":                      {Type: schema.TypeBool, Computed: true},
			"get_all_templates":                {Type: schema.TypeBool, Computed: true},
			"update_all_templates":             {Type: schema.TypeBool, Computed: true},
			"create_templates":                 {Type: schema.TypeBool, Computed: true},
			"delete_all_templates":             {Type: schema.TypeBool, Computed: true},
			"get_all_deployments":              {Type: schema.TypeBool, Computed: true},
			"update_all_deployments":           {Type: schema.TypeBool, Computed: true},
			"create_deployments_all_templates": {Type: schema.TypeBool, Computed: true},
			"delete_all_deployments":           {Type: schema.TypeBool, Computed: true},
			"get_global_settings":              {Type: schema.TypeBool, Computed: true},
			"update_global_settings":           {Type: schema.TypeBool, Computed: true},
		},
	}
}

func dataSourceRoleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	d.SetId(d.Get("id").(string))

	err := resourceRoleRead(ctx, d, meta)
	if err != nil {
		return err
	}

	return nil
}
