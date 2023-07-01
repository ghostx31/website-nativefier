resource "random_pet" "rg_name" {
  prefix = var.resource_group_location
}

resource "azurerm_resource_group" "rg" {
  location = var.resource_group_location
  name = random_pet.rg_name.id
}

resource "azurerm_container_registry" "nativefier" {
  name = "nativefier"
  resource_group_name = azurerm_resource_group.rg.name
  location = azurerm_resource_group.rg.location
  sku = "Basic"
  admin_enabled = true
}

resource "azurerm_user_assigned_identity" "containerapp" {
  location = azurerm_resource_group.rg.location
  name = "containerami"
  resource_group_name = azurerm_resource_group.rg.name
}

resource "azurerm_role_assignment" "containerRoleassignment" {
  scope = azurerm_container_registry.nativefier.id
  role_definition_name = "Contributor"
  principal_id = azurerm_user_assigned_identity.containerapp.principal_id
  depends_on = [
    azurerm_user_assigned_identity.containerapp
   ]
  skip_service_principal_aad_check = true
}

resource "null_resource" "docker_push" {
  provisioner "local-exec" {
    command = <<-EOT
    docker login ${azurerm_container_registry.nativefier.login_server} -u ${azurerm_container_registry.nativefier.admin_username} -p ${azurerm_container_registry.nativefier.admin_password}
    docker push ${azurerm_container_registry.nativefier.login_server}/nativefier:latest
  EOT
  }
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

# data "azurerm_user_assigned_identity" "agentpool_identity" {
#   name = "${azurerm_kubernetes_cluster.k8s.name}-agentpool"
#   resource_group_name = azurerm_resource_group.rg.name
#   depends_on = [ azurerm_kubernetes_cluster.k8s, azurerm_container_registry.nativefier ]
# }

# resource "azurerm_role_assignment" "kubernetes_acs_access" {
#   role_definition_name = "AcrPull"
#   scope = "${azurerm_container_registry.nativefier.id}"
#   skip_service_principal_aad_check = true
#   principal_id = data.azurerm_user_assigned_identity.agentpool_identity.principal_id
# }

# resource "azurerm_role_assignment" "acr2aks" {
#   principal_id = azurerm_kubernetes_cluster.k8s.kubelet_identity[0].object_id
#   role_definition_name = "AcrPull"
#   scope = azurerm_container_registry.nativefier.id
#   skip_service_principal_aad_check = true
#   depends_on = [ 
#     azurerm_kubernetes_cluster.k8s,
#     azurermazurerm_container_registry.nativefier
#   ]
# }

resource "helm_release" "nativefier" {
  name = "nativefier"
  repository = "${path.module}../nativefier-helm-chart"
  chart = "nativeifier"
  # depends_on = [ azurerm_kubernetes_cluster.k8s, azurerm_role_assignment.kubernetes_acs_access ]
}

