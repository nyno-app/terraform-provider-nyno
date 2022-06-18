terraform {
  required_providers {
    nyno = {
      version = "0.1.0"
      source  = "nyno.io/nyno/nyno"
    }
  }
}

provider "nyno" {
  alias = "something"
}

resource "nyno_template" "new" {
  provider    = nyno.something
  name        = "My User"
  description = "hello.com"
  
  variable {
    title = "Hello"
    variable = "ddd"
    type = "string"
  }
  action {
    type = "createFile"
    source_branch = "yoyoy"
    target_branch = "efef"
    path = "string"
    template_code = "efd"
    pull_request = false
    repository_id = "4fbf3dc3-2a66-41f3-87b7-70c7ec9f9609"
  }

}
