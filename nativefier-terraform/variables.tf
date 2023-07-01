variable "resource_group_location" {
  type        = string
  default     = "centralindia"
  description = "Location of the resource group."
}

variable "resource_group_name_prefix" {
  type        = string
  default     = "rg"
  description = "Prefix of the resource group name that's combined with a random ID so name is unique in your Azure subscription."
}

variable "node_count" {
  type = number
  description = "Quantity of nodes in pool initially"
  default = 1
}

variable "msi_id" {
  type = string
  description = "ID of the managed identity to use for the cluster"
  default = null
}