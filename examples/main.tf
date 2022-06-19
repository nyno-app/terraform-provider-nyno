terraform {
  required_providers {
    nyno = {
      version = "0.1.0"
      source  = "nyno.io/nyno/nyno"
    }
  }
}

provider "nyno" {
  api_endpoint = "http://localhost:3000/api"
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
    type          = "createFile"
    source_branch = "yoyoy"
    target_branch = "efeojf"
    path          = "string"
    template_code = "efd"
    pull_request  = false
    repository_id = "4fbf3dc3-2a66-41f3-87b7-70c7ec9f9609"
  }

  action {
    type          = "createFile2"
    source_branch = "yoyoy2"
    target_branch = "efeojf2"
    path          = "string"
    template_code = "efd222"
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
}

data "nyno_template" "template" {
  id = "087b10de-d672-4b88-8c84-14432521e856"
}

data "nyno_role" "my_role" {
  id = "5ec2219f-4e65-43d8-aa4f-ba880c1e436e"
}

data "nyno_repository" "my_rpository" {
  url = "https://github.com/Eitan1112/crypto-trader"
}
