terraform {
  required_providers {
    helm = {
      source  = "hashicorp/helm"
      version = "2.14.1"
    }

    kubectl = {
      source  = "gavinbunney/kubectl"
      version = "1.14.0"
    }
  }

  required_version = ">= 1.0"
}

provider "helm" {
  kubernetes {
    config_path = "~/.kube/config"
  }
}

provider "kubectl" {
  config_path = "~/.kube/config"
}
