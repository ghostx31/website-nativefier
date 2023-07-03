terraform {
  required_version = ">=1.0"

  required_providers {
    azapi = {
      source  = "azure/azapi"
      version = "~> 1.5"
    }
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "~> 3.0"
    }
    random = {
      source  = "hashicorp/random"
      version = "~> 3.0"
    }
    time = {
      source  = "hashicorp/time"
      version = "~> 0.9.1"
    }
    docker = {
      source = "kreuzwerker/docker"
      version = "3.0.2"
    }
  }
}

data "azurerm_kubernetes_cluster" "example" {
  name = azurerm_kubernetes_cluster.k8s.name
  resource_group_name = azurerm_resource_group.rg.name
}

locals {
  kubeconfig_path = "kubeconfig"
}

resource "local_file" "kubeconfig" {
  filename = local.kubeconfig_path
  content = data.azurerm_kubernetes_cluster.example.kube_config_raw
  file_permission = "0664"
}

provider "helm" {
  kubernetes {
    config_path = "kubeconfig"
  }
}

provider "azurerm" {
  features {
    
  }
  use_msi = true
}