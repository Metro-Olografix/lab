terraform {
  required_providers {
    coder = {
      source  = "coder/coder"
    }
    docker = {
      source  = "kreuzwerker/docker"
    }
  }
}

locals {
  username = data.coder_workspace_owner.me.name
}

data "coder_provisioner" "me" {
}

provider "docker" {
}

provider "coder" {
}

data "coder_workspace" "me" {
}

data "coder_workspace_owner" "me" {
}

data "coder_parameter" "go_version" {
  name        = "Go version"
  description = "Che versione di Go vuoi usare?"
  icon        = "/icon/go.svg"
  type        = "string"
  default     = "1.24.1"

  option {
    name = "1.24.1"
    value = "1.24.1"
  }
}