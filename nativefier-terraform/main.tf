resource "random_pet" "rg_name" {
  prefix = var.resource_group_location
}

resource "azurerm_resource_group" "rg" {
  location = var.resource_group_location
  name = random_pet.rg_name.id
}

resource "random_pet" "azurerm_kubernetes_cluster_name" {
  prefix = "cluster"
}
  
resource "random_pet" "azurerm_kubernetes_cluster_dns_prefix" {
  prefix = "dns"
}

resource "azurerm_kubernetes_cluster" "k8s" {
  location = azurerm_resource_group.rg.location
  name = random_pet.azurerm_kubernetes_cluster_name.id
  dns_prefix = random_pet.azurerm_kubernetes_cluster_dns_prefix.id
  resource_group_name = azurerm_resource_group.rg.name

  identity {
    type = "SystemAssigned"
  }

  default_node_pool {
    name = "agentpool"
    node_count = var.node_count
    vm_size = "Standard_A2_v2"
  }
  role_based_access_control_enabled = true
  linux_profile {
    admin_username = "azureuser"
    ssh_key {
      key_data = jsondecode(azapi_resource_action.ssh_public_key_gen.output).publicKey
    }
  }

  network_profile {
    network_plugin = "kubenet"
    load_balancer_sku = "basic"
  }
}

resource "time_sleep" "helm-wait" {
  create_duration = "10s"
  depends_on = [ helm_release.nativefier ]
}

resource "helm_release" "nativefier" {
  name = "nativefier"
  chart = "../nativefier-helm-chart"
  depends_on = [ azurerm_kubernetes_cluster.k8s ]
  set {
    name = "metrics.enabled"
    value = "true"
  }
}

