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

provider "helm" {
  kubernetes {
    config_path = "/home/admin1/.kube/config"
  }
}

provider "docker" {
  registry_auth {
    address = azurerm_container_registry.nativefier.login_server
    username = azurerm_container_registry.nativefier.admin_username
    password = azurerm_container_registry.nativefier.admin_password
  }
}
provider "kubernetes" {
  config_path = "/home/admin1/.kube/config"
}

provider "azurerm" {
  features {
    
  }
  use_msi = true
}