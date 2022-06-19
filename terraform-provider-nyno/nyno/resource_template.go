package nyno

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

type Action struct {
	ID           string `json:"id,omitempty"`
	Type         string `json:"type"`
	Path         string `json:"path"`
	SourceBranch string `json:"sourceBranch"`
	TargetBranch string `json:"targetBranch"`
	TemplateCode string `json:"templateCode"`
	PullRequest  bool   `json:"pullRequest"`
	RepositoryId string `json:"repositoryId"`
}

type Variable struct {
	ID           string `json:"id,omitempty"`
	Title        string `json:"title"`
	Variable     string `json:"variable"`
	Description  string `json:"description"`
	Type         string `json:"type"`
	DefaultValue string `json:"defaultValue"`
}

type Permissions struct {
	ID          string `json:"id,omitempty"`
	AccessLevel string `json:"accessLevel"`
	RoleId      string `json:"roleId"`
}

type Template struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Actions     []*Action      `json:"actions"`
	Variables   []*Variable    `json:"variables"`
	Permissions []*Permissions `json:"permissions"`
}

type TemplateCreateRequest struct {
	Template *Template `json:"template"`
}

type ResponseError struct {
	Error string `json:"error"`
}

func expandActions(config []interface{}) []*Action {
	actions := make([]*Action, 0, len(config))

	for _, rawAction := range config {
		actionConfig := rawAction.(map[string]interface{})

		action := &Action{
			Type:         actionConfig["type"].(string),
			Path:         actionConfig["path"].(string),
			SourceBranch: actionConfig["source_branch"].(string),
			TargetBranch: actionConfig["target_branch"].(string),
			TemplateCode: actionConfig["template_code"].(string),
			PullRequest:  actionConfig["pull_request"].(bool),
			RepositoryId: actionConfig["repository_id"].(string),
		}

		actions = append(actions, action)
	}

	return actions
}

func expandVariables(config []interface{}) []*Variable {
	variables := make([]*Variable, 0, len(config))

	for _, rawVariable := range config {
		variableConfig := rawVariable.(map[string]interface{})

		variable := &Variable{
			Title:        variableConfig["title"].(string),
			Variable:     variableConfig["variable"].(string),
			Description:  variableConfig["description"].(string),
			Type:         variableConfig["type"].(string),
			DefaultValue: variableConfig["default_value"].(string),
		}

		variables = append(variables, variable)
	}

	return variables
}

func expandPermissions(config []interface{}) []*Permissions {
	permissions := make([]*Permissions, 0, len(config))

	for _, rawPermissions := range config {
		permissionsConfig := rawPermissions.(map[string]interface{})

		permission := &Permissions{
			AccessLevel: permissionsConfig["access_level"].(string),
			RoleId:      permissionsConfig["role_id"].(string),
		}

		permissions = append(permissions, permission)
	}

	return permissions
}

// Resource schema definition
func resourceTemplate() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceTemplateCreate,
		ReadContext:   resourceTemplateRead,
		UpdateContext: resourceTemplateUpdate,
		DeleteContext: resourceTemplateDelete,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString, // Field type
				Computed: true,              // This flag means that the fields will be created after some processing
			},
			"name": {
				Type:     schema.TypeString,
				Required: true, // Field is required
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"action": {
				Type:     schema.TypeList,
				MinItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"source_branch": {
							Type:     schema.TypeString,
							Required: true,
						},
						"target_branch": {
							Type:     schema.TypeString,
							Required: true,
						},
						"path": {
							Type:     schema.TypeString,
							Required: true,
						},
						"template_code": {
							Type:     schema.TypeString,
							Required: true,
						},
						"pull_request": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"repository_id": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"variable": {
				Type:     schema.TypeList,
				MinItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"title": {
							Type:     schema.TypeString,
							Required: true,
						},
						"variable": {
							Type:     schema.TypeString,
							Required: true,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"default_value": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"permissions": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"access_level": {
							Type:     schema.TypeString,
							Required: true,
						},
						"role_id": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourceTemplateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}

	// Getting from terraform
	name := d.Get("name").(string)
	description := d.Get("description").(string)
	actions := d.Get("action").([]interface{})
	variables := d.Get("variable").([]interface{})
	permissions := d.Get("permissions").([]interface{})

	// Build Template object
	template := &Template{
		Name:        name,
		Description: description,
		Actions:     expandActions(actions),
		Variables:   expandVariables(variables),
		Permissions: expandPermissions(permissions),
	}

	requestBody, err := json.Marshal(template)

	if err != nil {
		return diag.FromErr(err)
	}

	body := bytes.NewBuffer(requestBody)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/templates", m.(Config).api_endpoint), body)
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
			return diag.Errorf("Unable to create template. Status Code: %v", r.StatusCode)
		}
		return diag.Errorf("Unable to create template. Status Code: %v. Message: %s", r.StatusCode, response.Error)
	}

	var response Template
	err = json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("name", response.Name)
	d.Set("description", response.Description)
	d.Set("actions", response.Actions)
	d.Set("variables", response.Variables)
	d.Set("permissions", response.Permissions)
	d.SetId(response.ID)

	return nil
}

func resourceTemplateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}

	log.Print(fmt.Sprintf("Sending request to: %[1]s/templates/%[2]s", m.(Config).api_endpoint, d.Id()))
	log.Printf("ID:", d.Id())
	req, err := http.NewRequest("GET", fmt.Sprintf("%[1]s/templates/%[2]s", m.(Config).api_endpoint, d.Id()), nil)
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
		log.Printf("MyResponse: %s", response)
		if response == nil {
			return diag.Errorf("Unable to read template. Status Code: %v", r.StatusCode)
		}
		return diag.Errorf("Unable to read template. Status Code: %v. Message: %s", r.StatusCode, response.Error)
	}

	var response Template
	err = json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("name", response.Name)
	d.Set("description", response.Description)
	d.Set("actions", response.Actions)
	d.Set("variables", response.Variables)
	d.Set("permissions", response.Permissions)

	return nil
}

func resourceTemplateUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}

	// Getting from terraform
	id := d.Id()
	name := d.Get("name").(string)
	description := d.Get("description").(string)
	actions := d.Get("action").([]interface{})
	variables := d.Get("variable").([]interface{})
	permissions := d.Get("permissions").([]interface{})

	// Build Template object
	template := &Template{
		ID:          id,
		Name:        name,
		Description: description,
		Actions:     expandActions(actions),
		Variables:   expandVariables(variables),
		Permissions: expandPermissions(permissions),
	}

	requestBody, err := json.Marshal(template)

	if err != nil {
		return diag.FromErr(err)
	}

	body := bytes.NewBuffer(requestBody)

	req, err := http.NewRequest("PUT", fmt.Sprintf("%[1]s/templates/%[2]s", m.(Config).api_endpoint, d.Id()), body)
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
			return diag.Errorf("Unable to update template. Status Code: %v", r.StatusCode)
		}

		return diag.Errorf("Unable to update template. Status Code: %v. Message: %s", r.StatusCode, response.Error)
	}

	var response Template
	err = json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("name", response.Name)
	d.Set("description", response.Description)
	d.Set("actions", response.Actions)
	d.Set("variables", response.Variables)
	d.Set("permissions", response.Permissions)

	return nil

}

func resourceTemplateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%[1]s/templates/%[2]s", m.(Config).api_endpoint, d.Id()), nil)
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
			return diag.Errorf("Unable to delete template. Status Code: %v", r.StatusCode)
		}
		return diag.Errorf("Unable to delete template. Status Code: %v. Message: %s", r.StatusCode, response.Error)
	}

	defer r.Body.Close()

	var response *Template
	err = json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}
