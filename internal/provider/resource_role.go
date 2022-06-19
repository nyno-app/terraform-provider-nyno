package provider

import (
	"bytes"
	"context"
	"encoding/json"

	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Role struct {
	ID string `json:"id,omitempty"`

	Name                          string `json:"name"`
	CreateCredentials             bool   `json:"createCredentials"`
	GetCredentials                bool   `json:"getCredentials"`
	UpdateCredentials             bool   `json:"updateCredentials"`
	DeleteCredentials             bool   `json:"deleteCredentials"`
	CreateRepository              bool   `json:"createRepository"`
	GetRepository                 bool   `json:"getRepository"`
	UpdateRepository              bool   `json:"updateRepository"`
	DeleteRepository              bool   `json:"deleteRepository"`
	GetUser                       bool   `json:"getUser"`
	UpdateUser                    bool   `json:"updateUser"`
	GetRole                       bool   `json:"getRole"`
	UpdateRole                    bool   `json:"updateRole"`
	CreateRole                    bool   `json:"createRole"`
	DeleteRole                    bool   `json:"deleteRole"`
	GetAllTemplates               bool   `json:"getAllTemplates"`
	UpdateAllTemplates            bool   `json:"updateAllTemplates"`
	CreateTemplate                bool   `json:"createTemplate"`
	DeleteAllTemplates            bool   `json:"deleteAllTemplates"`
	GetAllDeployments             bool   `json:"getAllDeployments"`
	UpdateAllDeployments          bool   `json:"updateAllDeployments"`
	CreateDeploymentsAllTemplates bool   `json:"createDeploymentsAllTemplates"`
	DeleteAllDeployment           bool   `json:"deleteAllDeployment"`
}

// Resource schema definition
func resourceRole() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceRoleCreate,
		ReadContext:   resourceRoleRead,
		UpdateContext: resourceRoleUpdate,
		DeleteContext: resourceRoleDelete,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString, // Field type
				Computed: true,              // This flag means that the fields will be created after some processing
			},
			"name":                             {Type: schema.TypeString, Required: true},
			"create_credentials":               {Type: schema.TypeBool, Required: true},
			"get_credentials":                  {Type: schema.TypeBool, Required: true},
			"update_credentials":               {Type: schema.TypeBool, Required: true},
			"delete_credentials":               {Type: schema.TypeBool, Required: true},
			"create_repository":                {Type: schema.TypeBool, Required: true},
			"get_repository":                   {Type: schema.TypeBool, Required: true},
			"update_repository":                {Type: schema.TypeBool, Required: true},
			"delete_repository":                {Type: schema.TypeBool, Required: true},
			"get_user":                         {Type: schema.TypeBool, Required: true},
			"update_user":                      {Type: schema.TypeBool, Required: true},
			"get_role":                         {Type: schema.TypeBool, Required: true},
			"update_role":                      {Type: schema.TypeBool, Required: true},
			"create_role":                      {Type: schema.TypeBool, Required: true},
			"delete_role":                      {Type: schema.TypeBool, Required: true},
			"get_all_templates":                {Type: schema.TypeBool, Required: true},
			"update_all_templates":             {Type: schema.TypeBool, Required: true},
			"create_templates":                 {Type: schema.TypeBool, Required: true},
			"delete_all_templates":             {Type: schema.TypeBool, Required: true},
			"get_all_deployments":              {Type: schema.TypeBool, Required: true},
			"update_all_deployments":           {Type: schema.TypeBool, Required: true},
			"create_deployments_all_templates": {Type: schema.TypeBool, Required: true},
			"delete_all_deployments":           {Type: schema.TypeBool, Required: true},
		},
	}
}

func resourceRoleCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}

	role := &Role{
		Name:                          d.Get("name").(string),
		CreateCredentials:             d.Get("create_credentials").(bool),
		GetCredentials:                d.Get("get_credentials").(bool),
		UpdateCredentials:             d.Get("update_credentials").(bool),
		DeleteCredentials:             d.Get("delete_credentials").(bool),
		CreateRepository:              d.Get("create_repository").(bool),
		GetRepository:                 d.Get("get_repository").(bool),
		UpdateRepository:              d.Get("update_repository").(bool),
		DeleteRepository:              d.Get("delete_repository").(bool),
		GetUser:                       d.Get("get_user").(bool),
		UpdateUser:                    d.Get("update_user").(bool),
		GetRole:                       d.Get("get_role").(bool),
		UpdateRole:                    d.Get("update_role").(bool),
		CreateRole:                    d.Get("create_role").(bool),
		DeleteRole:                    d.Get("delete_role").(bool),
		GetAllTemplates:               d.Get("get_all_templates").(bool),
		UpdateAllTemplates:            d.Get("update_all_templates").(bool),
		CreateTemplate:                d.Get("create_templates").(bool),
		DeleteAllTemplates:            d.Get("delete_all_templates").(bool),
		GetAllDeployments:             d.Get("get_all_deployments").(bool),
		UpdateAllDeployments:          d.Get("update_all_deployments").(bool),
		CreateDeploymentsAllTemplates: d.Get("create_deployments_all_templates").(bool),
		DeleteAllDeployment:           d.Get("delete_all_deployments").(bool),
	}

	requestBody, err := json.Marshal(role)

	if err != nil {
		return diag.FromErr(err)
	}

	body := bytes.NewBuffer(requestBody)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/roles", m.(Config).api_endpoint), body)
	if err != nil {
		return diag.FromErr(err)
	}

	req.Header.Set("Content-Type", "application/json")

	log.Println("Sending http request")
	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}

	if r.StatusCode != 200 {
		var response *ResponseError
		err = json.NewDecoder(r.Body).Decode(&response)

		if response == nil {
			return diag.Errorf("Unable to create role. Status Code: %v", r.StatusCode)
		}
		return diag.Errorf("Unable to create role. Status Code: %v. Message: %s", r.StatusCode, response.Error)
	}

	var response Role
	err = json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("name", response.Name)
	d.Set("create_credentials", response.CreateCredentials)
	d.Set("get_credentials", response.GetCredentials)
	d.Set("update_credentials", response.UpdateCredentials)
	d.Set("delete_credentials", response.DeleteCredentials)
	d.Set("create_repository", response.CreateRepository)
	d.Set("get_repository", response.GetRepository)
	d.Set("update_repository", response.UpdateRepository)
	d.Set("delete_repository", response.DeleteRepository)
	d.Set("get_user", response.GetUser)
	d.Set("update_user", response.UpdateUser)
	d.Set("get_role", response.GetRole)
	d.Set("update_role", response.UpdateRole)
	d.Set("create_role", response.CreateRole)
	d.Set("delete_role", response.DeleteRole)
	d.Set("get_all_templates", response.GetAllTemplates)
	d.Set("update_all_templates", response.UpdateAllTemplates)
	d.Set("create_templates", response.CreateTemplate)
	d.Set("delete_all_templates", response.DeleteAllTemplates)
	d.Set("get_all_deployments", response.GetAllDeployments)
	d.Set("update_all_deployments", response.UpdateAllDeployments)
	d.Set("create_deployments_all_templates", response.CreateDeploymentsAllTemplates)
	d.Set("delete_all_deployments", response.DeleteAllDeployment)
	d.SetId(response.ID)

	return nil
}

func resourceRoleRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest("GET", fmt.Sprintf("%[1]s/roles/%[2]s", m.(Config).api_endpoint, d.Id()), nil)
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
			return diag.Errorf("Unable to read role. Status Code: %v", r.StatusCode)
		}
		return diag.Errorf("Unable to read role. Status Code: %v. Message: %s", r.StatusCode, response.Error)
	}

	var response Role
	err = json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("name", response.Name)
	d.Set("create_credentials", response.CreateCredentials)
	d.Set("get_credentials", response.GetCredentials)
	d.Set("update_credentials", response.UpdateCredentials)
	d.Set("delete_credentials", response.DeleteCredentials)
	d.Set("create_repository", response.CreateRepository)
	d.Set("get_repository", response.GetRepository)
	d.Set("update_repository", response.UpdateRepository)
	d.Set("delete_repository", response.DeleteRepository)
	d.Set("get_user", response.GetUser)
	d.Set("update_user", response.UpdateUser)
	d.Set("get_role", response.GetRole)
	d.Set("update_role", response.UpdateRole)
	d.Set("create_role", response.CreateRole)
	d.Set("delete_role", response.DeleteRole)
	d.Set("get_all_templates", response.GetAllTemplates)
	d.Set("update_all_templates", response.UpdateAllTemplates)
	d.Set("create_templates", response.CreateTemplate)
	d.Set("delete_all_templates", response.DeleteAllTemplates)
	d.Set("get_all_deployments", response.GetAllDeployments)
	d.Set("update_all_deployments", response.UpdateAllDeployments)
	d.Set("create_deployments_all_templates", response.CreateDeploymentsAllTemplates)
	d.Set("delete_all_deployments", response.DeleteAllDeployment)

	return nil
}

func resourceRoleUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}

	role := &Role{
		ID:                            d.Id(),
		Name:                          d.Get("name").(string),
		CreateCredentials:             d.Get("create_credentials").(bool),
		GetCredentials:                d.Get("get_credentials").(bool),
		UpdateCredentials:             d.Get("update_credentials").(bool),
		DeleteCredentials:             d.Get("delete_credentials").(bool),
		CreateRepository:              d.Get("create_repository").(bool),
		GetRepository:                 d.Get("get_repository").(bool),
		UpdateRepository:              d.Get("update_repository").(bool),
		DeleteRepository:              d.Get("delete_repository").(bool),
		GetUser:                       d.Get("get_user").(bool),
		UpdateUser:                    d.Get("update_user").(bool),
		GetRole:                       d.Get("get_role").(bool),
		UpdateRole:                    d.Get("update_role").(bool),
		CreateRole:                    d.Get("create_role").(bool),
		DeleteRole:                    d.Get("delete_role").(bool),
		GetAllTemplates:               d.Get("get_all_templates").(bool),
		UpdateAllTemplates:            d.Get("update_all_templates").(bool),
		CreateTemplate:                d.Get("create_templates").(bool),
		DeleteAllTemplates:            d.Get("delete_all_templates").(bool),
		GetAllDeployments:             d.Get("get_all_deployments").(bool),
		UpdateAllDeployments:          d.Get("update_all_deployments").(bool),
		CreateDeploymentsAllTemplates: d.Get("create_deployments_all_templates").(bool),
		DeleteAllDeployment:           d.Get("delete_all_deployments").(bool),
	}

	requestBody, err := json.Marshal(role)

	if err != nil {
		return diag.FromErr(err)
	}

	body := bytes.NewBuffer(requestBody)

	req, err := http.NewRequest("PUT", fmt.Sprintf("%[1]s/roles/%[2]s", m.(Config).api_endpoint, d.Id()), body)
	if err != nil {
		return diag.FromErr(err)
	}

	req.Header.Set("Content-Type", "application/json")
	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}

	if r.StatusCode != 200 {
		var response *ResponseError
		err = json.NewDecoder(r.Body).Decode(&response)

		if response == nil {
			return diag.Errorf("Unable to update role. Status Code: %v", r.StatusCode)
		}

		return diag.Errorf("Unable to update role. Status Code: %v. Message: %s", r.StatusCode, response.Error)
	}

	var response Role
	err = json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("name", response.Name)
	d.Set("create_credentials", response.CreateCredentials)
	d.Set("get_credentials", response.GetCredentials)
	d.Set("update_credentials", response.UpdateCredentials)
	d.Set("delete_credentials", response.DeleteCredentials)
	d.Set("create_repository", response.CreateRepository)
	d.Set("get_repository", response.GetRepository)
	d.Set("update_repository", response.UpdateRepository)
	d.Set("delete_repository", response.DeleteRepository)
	d.Set("get_user", response.GetUser)
	d.Set("update_user", response.UpdateUser)
	d.Set("get_role", response.GetRole)
	d.Set("update_role", response.UpdateRole)
	d.Set("create_role", response.CreateRole)
	d.Set("delete_role", response.DeleteRole)
	d.Set("get_all_templates", response.GetAllTemplates)
	d.Set("update_all_templates", response.UpdateAllTemplates)
	d.Set("create_templates", response.CreateTemplate)
	d.Set("delete_all_templates", response.DeleteAllTemplates)
	d.Set("get_all_deployments", response.GetAllDeployments)
	d.Set("update_all_deployments", response.UpdateAllDeployments)
	d.Set("create_deployments_all_templates", response.CreateDeploymentsAllTemplates)
	d.Set("delete_all_deployments", response.DeleteAllDeployment)

	return nil

}

func resourceRoleDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%[1]s/roles/%[2]s", m.(Config).api_endpoint, d.Id()), nil)
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
			return diag.Errorf("Unable to delete role. Status Code: %v", r.StatusCode)
		}
		return diag.Errorf("Unable to delete role. Status Code: %v. Message: %s", r.StatusCode, response.Error)
	}

	defer r.Body.Close()

	var response *Role
	err = json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}
