terraform {
  required_version = ">= 1.1.9, < 2.0.0"

  required_providers {
    azurerm = {
      source = "hashicorp/azurerm"
      version = "3.3.0"
    }
  }  
}

provider "azurerm" {
    subscription_id = var.subscription_id
    features {}
}