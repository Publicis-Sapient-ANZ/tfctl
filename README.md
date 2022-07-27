# Terraform Control

## Purpose

The tfctl cli is used plan and apply a terraform configuration is a consistent repeatable way either locally for development or in a CI pipeline. 

Its main features are:

- common configuration across N number of environments
- ability to have additional terraform configuration per environment
- support for the following backends:
    - AzureRM storage account

## Installation

To install the ftctl binary run to following:

TODO: create the below script....

```
curl -sS https://github.com/Publicis-Sapient-ANZ/tfctl/blob/main/install.sh | sh
```

or optionally use the provided Dockerfile to create a container with the cli baked in.

## Usage

### Inititalise a new project configuration set

TODO: This is not currently implemented

```
tfctl init
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

## Configuration File

tfctl configuration is stored per environment. The file must be named \<env-name\>.tfctl.yaml

Sample configuration is provided below. Refer to the schema for information on each element.

``` yaml
apiVersion: tfctl/v1alpha1
kind: Config
metadata:
  name: dev-config
  environment: dev
  description: Dev terraform configuration
spec:
  backend:
    azure:
      autoCreateStorage: true
      storageAccountRg: RD_DEV_AUS
      storageAccountName: stgdevtfstateaus
      storageAccountContainer: dev-infra-state
      stateFileName: dev.tfstate
      subscriptionId: 00000000-0000-0000-0000-000000000000
      tenantId: 00000000-0000-0000-0000-000000000000
      credentials:
        useAzLogin: fallback
        clientIdEnvName: servicePrincipalId
        clientSecretEnvName: servicePrincipalKey
```

## Configuration Schema

**_Config_**

 - ## Config
 - Type: `object`
 - <i id="/definitions/Config">path: #/definitions/Config</i>
 - This schema <u>does not</u> accept additional properties.
 - **_Properties_**
	 - <b id="#/definitions/Config/properties/apiVersion">apiVersion</b> `required`
		 - Type: `string`
		 - <i id="/definitions/Config/properties/apiVersion">path: #/definitions/Config/properties/apiVersion</i>
	 - <b id="#/definitions/Config/properties/kind">kind</b> `required`
		 - Type: `string`
		 - <i id="/definitions/Config/properties/kind">path: #/definitions/Config/properties/kind</i>
	 - <b id="#/definitions/Config/properties/metadata">metadata</b> `required`
		 - <i id="/definitions/Config/properties/metadata">path: #/definitions/Config/properties/metadata</i>
		 - &#36;ref: [#/definitions/Metadata](#/definitions/Metadata)
	 - <b id="#/definitions/Config/properties/spec">spec</b> `required`
		 - <i id="/definitions/Config/properties/spec">path: #/definitions/Config/properties/spec</i>
		 - &#36;ref: [#/definitions/Spec](#/definitions/Spec)


**_Metadata_**

 - ## Metadata
 - Type: `object`
 - <i id="/definitions/Metadata">path: #/definitions/Metadata</i>
 - This schema <u>does not</u> accept additional properties.
 - **_Properties_**
	 - <b id="#/definitions/Metadata/properties/name">name</b> `required`
		 - Type: `string`
		 - <i id="/definitions/Metadata/properties/name">path: #/definitions/Metadata/properties/name</i>
	 - <b id="#/definitions/Metadata/properties/environment">environment</b> `required`
		 - Type: `string`
		 - <i id="/definitions/Metadata/properties/environment">path: #/definitions/Metadata/properties/environment</i>
	 - <b id="#/definitions/Metadata/properties/description">description</b> `required`
		 - Type: `string`
		 - <i id="/definitions/Metadata/properties/description">path: #/definitions/Metadata/properties/description</i>


**_Spec_**

 - ## Spec
 - Type: `object`
 - <i id="/definitions/Spec">path: #/definitions/Spec</i>
 - This schema <u>does not</u> accept additional properties.
 - **_Properties_**
	 - <b id="#/definitions/Spec/properties/backend">backend</b> `required`
		 - <i id="/definitions/Spec/properties/backend">path: #/definitions/Spec/properties/backend</i>
		 - &#36;ref: [#/definitions/Backend](#/definitions/Backend)


**_Backend_**

 - ## Backend
 - Type: `object`
 - <i id="/definitions/Backend">path: #/definitions/Backend</i>
 - This schema <u>does not</u> accept additional properties.
 - **_Properties_**
	 - <b id="#/definitions/Backend/properties/azure">azure</b>
		 - <i id="/definitions/Backend/properties/azure">path: #/definitions/Backend/properties/azure</i>
		 - &#36;ref: [#/definitions/Azure](#/definitions/Azure)


**_Azure_**

 - ## Azure
 - Type: `object`
 - <i id="/definitions/Azure">path: #/definitions/Azure</i>
 - This schema <u>does not</u> accept additional properties.
 - **_Properties_**
	 - <b id="#/definitions/Azure/properties/autoCreateStorage">autoCreateStorage</b> `required`
		 - Type: `boolean`
		 - <i id="/definitions/Azure/properties/autoCreateStorage">path: #/definitions/Azure/properties/autoCreateStorage</i>
	 - <b id="#/definitions/Azure/properties/storageAccountRg">storageAccountRg</b> `required`
		 - Type: `string`
		 - <i id="/definitions/Azure/properties/storageAccountRg">path: #/definitions/Azure/properties/storageAccountRg</i>
	 - <b id="#/definitions/Azure/properties/storageAccountName">storageAccountName</b> `required`
		 - Type: `string`
		 - <i id="/definitions/Azure/properties/storageAccountName">path: #/definitions/Azure/properties/storageAccountName</i>
	 - <b id="#/definitions/Azure/properties/storageAccountContainer">storageAccountContainer</b> `required`
		 - Type: `string`
		 - <i id="/definitions/Azure/properties/storageAccountContainer">path: #/definitions/Azure/properties/storageAccountContainer</i>
	 - <b id="#/definitions/Azure/properties/stateFileName">stateFileName</b> `required`
		 - Type: `string`
		 - <i id="/definitions/Azure/properties/stateFileName">path: #/definitions/Azure/properties/stateFileName</i>
	 - <b id="#/definitions/Azure/properties/subscriptionId">subscriptionId</b> `required`
		 - Type: `string`
		 - <i id="/definitions/Azure/properties/subscriptionId">path: #/definitions/Azure/properties/subscriptionId</i>
		 - String format must be a "uuid"
	 - <b id="#/definitions/Azure/properties/tenantId">tenantId</b> `required`
		 - Type: `string`
		 - <i id="/definitions/Azure/properties/tenantId">path: #/definitions/Azure/properties/tenantId</i>
		 - String format must be a "uuid"
	 - <b id="#/definitions/Azure/properties/credentials">credentials</b> `required`
		 - <i id="/definitions/Azure/properties/credentials">path: #/definitions/Azure/properties/credentials</i>
		 - &#36;ref: [#/definitions/Credentials](#/definitions/Credentials)


**_Credentials_**

 - ## Credentials
 - Type: `object`
 - <i id="/definitions/Credentials">path: #/definitions/Credentials</i>
 - This schema <u>does not</u> accept additional properties.
 - **_Properties_**
	 - <b id="#/definitions/Credentials/properties/useAzLogin">useAzLogin</b> `required`
		 - Type: `string`
		 - <i id="/definitions/Credentials/properties/useAzLogin">path: #/definitions/Credentials/properties/useAzLogin</i>
	 - <b id="#/definitions/Credentials/properties/clientIdEnvName">clientIdEnvName</b>
		 - Type: `string`
		 - <i id="/definitions/Credentials/properties/clientIdEnvName">path: #/definitions/Credentials/properties/clientIdEnvName</i>
	 - <b id="#/definitions/Credentials/properties/clientSecretEnvName">clientSecretEnvName</b>
		 - Type: `string`
		 - <i id="/definitions/Credentials/properties/clientSecretEnvName">path: #/definitions/Credentials/properties/clientSecretEnvName</i>


## Local Development

### Build

To build the cli run:

```
go build -o tfctl
```

### Develop

Please use the .devcontainer to start a new local development environment. It will install and configure all dependencies for you. Check the [README.md](./.devcontainer/README.md) for details.
