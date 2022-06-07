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

resource "template" "new" {
  provider    = nyno.something
  name        = "My User"
  description = "hello.com"
}
