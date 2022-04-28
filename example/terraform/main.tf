resource "azurerm_resource_group" "resource-group" {
  name     = var.rg
  location = var.location
}