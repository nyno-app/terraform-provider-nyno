terraform {
  required_providers {
    nyno = {
      version = "1.1.1"
      source  = "nyno.io/nyno/nyno"
    }
  }
}

provider "nyno" {
  api_endpoint = "http://localhost:3000/api"
  organization = "customer"
  username     = "someuser"
  password     = "123456789123456789"
}

resource "nyno_template" "new" {
  name        = "My User2"
  description = "hello.com"

  variable {
    title    = "dewooj"
    variable = "ddoj"
    type     = "string"
  }

  variable {
    title    = "kpo"
    variable = "dpkdoj"
    type     = "string"
  }


  action {
    type          = "createFile2"
    source_branch = "yoyoy2"
    target_branch = "efeojf2"
    path          = "string"
    template_code = filebase64("templatefile.yaml")
    pull_request  = true
    repository_id = data.nyno_repository.my_rpository.id
  }
}

resource "nyno_role" "new" {
  name                             = "my-awesome-rojojole"
  create_credentials               = true
  get_credentials                  = true
  update_credentials               = true
  delete_credentials               = true
  create_repository                = true
  get_repository                   = true
  update_repository                = true
  delete_repository                = true
  get_user                         = true
  update_user                      = true
  create_user                      = true
  get_role                         = false
  update_role                      = true
  create_role                      = true
  delete_role                      = true
  get_all_templates                = true
  update_all_templates             = true
  create_templates                 = false
  delete_all_templates             = true
  get_all_deployments              = true
  update_all_deployments           = true
  create_deployments_all_templates = true
  delete_all_deployments           = true
  get_global_settings              = true
  update_global_settings           = true
}

data "nyno_template" "template" {
  id = "45a73a32-7ac9-4639-b0a7-33c1ea840244"
}

data "nyno_role" "my_role" {
  id = "9bf21461-eb62-4289-a4be-e446a495d20f"
}

data "nyno_repository" "my_rpository" {
  url = "https://github.com/Eitan1112/chat-app"
}
