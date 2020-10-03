terraform {
  required_providers {
    dash = {
      versions = ["0.1"]
      source   = "packages.example.com/terraform/dash"
    }
  }
}

provider "dash" {}
