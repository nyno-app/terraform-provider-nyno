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
  name        = "My User"
  description = "hello.com"

  variable {
    title    = "dewooj"
    variable = "ddoj"
    type     = "string"
  }
  action {
    type          = "createFile"
    source_branch = "yoyoy"
    target_branch = "efef"
    path          = "string"
    template_code = "efd"
    pull_request  = false
    repository_id = "4fbf3dc3-2a66-41f3-87b7-70c7ec9f9609"
  }
}

data "nyno_template" "template" {
  id = "087b10de-d672-4b88-8c84-14432521e856"
}
