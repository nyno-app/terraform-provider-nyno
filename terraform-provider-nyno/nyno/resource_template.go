package nyno

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	// "strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const baseURL = "http://localhost:3000"

// Create a struct to handler template object.

// type Action struct {
// 	ID           string `json:"id"`
// 	id           string `json:"id"`
// 	Type         string `json:"type"`
// 	Path         string `json:"path"`
// 	SourceBranch string `json:"sourceBranch"`
// 	TargetBranch string `json:"targetBranch"`
// 	TemplateCode string `json:"templateCode"`
// 	PullRequest  string `json:"pullRequest"`
// 	RepositoryId string `json:"repositoryId"`
// }

type Template struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Function definition to create the resource.

func resourceTemplateCreate(d *schema.ResourceData, m interface{}) error {
	client := &http.Client{Timeout: 10 * time.Second}

	// Get fields from terraform resource data.

	name := d.Get("name").(string)
	description := d.Get("description").(string)

	// Build Template object

	template := &Template{
		Name:        name,
		Description: description,
	}

	requestBody, err := json.Marshal(template)
	if err != nil {
		return err
	}

	body := bytes.NewBuffer(requestBody)

	// Send POST to /templates with Template informations
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/templates", baseURL), body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	r, err := client.Do(req)
	if err != nil {
		return err
	}

	// Get response from API
	var response Template
	err = json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		return err
	}

	// Set response on terraform state
	// Here you can create validation to check if all fields are valid.

	d.Set("name", response.Name)
	d.Set("description", response.Description)
	// Set state id for resource
	d.SetId(response.ID)

	return nil
}

// Function definition to read the resource.

func resourceTemplateRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

// Function definition to update the resource.

func resourceTemplateUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

// Function definition to delete the resource.

func resourceTemplateDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}

// Resource schema definition

func resourceTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceTemplateCreate,
		Read:   resourceTemplateRead,
		Update: resourceTemplateUpdate,
		Delete: resourceTemplateDelete,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString, // Field type
				Computed: true,              // This flag means that the fields will be created after some processing
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true, // Field is required
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
