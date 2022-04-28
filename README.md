# Terraform Control Command Line Tool

## Purpose

The tfctl cli is used plan and apply a terraform configuration is a consistent repeatable way either locally for development or in a CI pipeline. 

Its main features are:

- common configuration across N number of environments
- ability to have additional tf configuration per environment
- support for the following backends:
    - AzureRM storage account

## Usage


### Perform a Terraform Validate

### Perform a Terraform Plan

### Perform a Terraform Apply


## Local Development

### Build

To build the cli run:

```
go build
```

### Develop

Please use the .devcontainer to start a new local environment. It will install and configure all dependencies for you. 

To run the qntzcli run the below and follow the help output for syntax

```
go run main.go
```
