# Terraform Control

## Purpose

The tfctl cli is used plan and apply a terraform configuration is a consistent repeatable way either locally for development or in a CI pipeline. 

Its main features are:

- common configuration across N number of environments
- ability to have additional tf configuration per environment
- support for the following backends:
    - AzureRM storage account

## Installation

To install the ftctl binary run to following:

TODO: create the below script....

```
curl -sS https://github.com/Publicis-Sapient-ANZ/tfctl/blob/main/install.sh | sh
```

or optionally use the provided Dockerfile to create a cointainer with the cli baked in.

## Usage

### Inititaliase a new project configuration set

TODO: This is not currently implemented

```
tfctl project init
```

### Perform a Terraform Validate

```
tfctl validate --config-path ./example --env-name dev
```

### Perform a Terraform Plan

```
tfctl plan --config-path ./example --env-name dev
```

### Perform a Terraform Apply

```
tfctl apply --config-path ./example --env-name dev
```

### Perform a Terraform Destroy

```
tfctl destroy --config-path ./example --env-name dev
```

## Configuration

tfctl configuration is stored per environment. The file must be named <env-name>.tfctl.yaml

Sample configuration is provided below. Refer to the table for information on each element.

``` yaml
apiVersion: tfctl/v1alpha1
kind: Config
metadata:
  name: dev-config
  environment: dev
  description: Dev terraform configuration
spec:
  backend:
    type: azure
    autoCreateStorage: true
    storageAccountRg: RD_DEV_AUS
    storageAccountName: stgdevtfstateaus
    storageAccountContainer: dev-infra-state
    stateFileName: dev.tfstate
    subscriptionId: 00000000-0000-0000-0000-000000000000
    tenantId: 00000000-0000-0000-0000-000000000000
    credentials:
      useAzLogin: fallback
      fromEnvironment: ifPresent
      clientIdEnvName: servicePrincipalId
      clientSecretEnvName: servicePrincipalKey
```

TODO: table showing config params

## Local Development

### Build

To build the cli run:

```
go build
```

### Develop

Please use the .devcontainer to start a new local environment. It will install and configure all dependencies for you.
