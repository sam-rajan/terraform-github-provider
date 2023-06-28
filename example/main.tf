
terraform {
  required_providers {
    mygithub = {
      source  = "hashicorp/mygithub"
      version = "1.0.0"
    }
  }
}

variable "github_token" {
  type = string
}

provider "mygithub" {
  auth_token = var.github_token
}

resource "mygithub_repo" "myrepo" {
  name        = "sample"
  description = "This is a sample respository"
  homepage    = "sample.com"
}

data "mygithub_repo" "myrepo_data" {
  name = "abc"
}


output "myrepo_url" {
  value = mygithub_repo.myrepo.url
}

output "myrepo_data_name" {
  value = data.mygithub_repo.myrepo_data.url
}
