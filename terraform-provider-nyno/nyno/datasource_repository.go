package nyno

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Repository struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Url      string `json:"url"`
	IsActive bool   `json:"isActive"`
}

func dataSourceRepository() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceRepositoryRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"url": {
				Type:     schema.TypeString,
				Required: true,
			},
			"is_active": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func resourceRepositoryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}

	log.Print("INREPOREAD")
	log.Print(fmt.Sprintf("%[1]s/repositories/%[2]s", m.(Config).api_endpoint, url.QueryEscape(d.Id())))
	req, err := http.NewRequest("GET", fmt.Sprintf("%[1]s/repositories/%[2]s", m.(Config).api_endpoint, url.QueryEscape(d.Get("url").(string))), nil)
	if err != nil {
		return diag.FromErr(err)
	}

	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}

	if r.StatusCode != 200 {
		var response *ResponseError
		err = json.NewDecoder(r.Body).Decode(&response)
		if response == nil {
			return diag.Errorf("Unable to read repository. Status Code: %v", r.StatusCode)
		}
		return diag.Errorf("Unable to read repository. Status Code: %v. Message: %s", r.StatusCode, response.Error)
	}

	var response Repository
	err = json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("name", response.Name)
	d.Set("url", response.Url)
	d.Set("is_active", response.IsActive)
	d.Set("id", response.ID)
	d.SetId(response.ID)

	return nil
}

func dataSourceRepositoryRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	err := resourceRepositoryRead(ctx, d, meta)
	if err != nil {
		return err
	}

	return nil
}
