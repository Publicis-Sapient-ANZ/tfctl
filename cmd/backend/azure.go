package backend

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/sirupsen/logrus"

	"github.com/quantize-io/tfctl/cmd/model"
)

func initAzureBacked(config *model.Config) error {

	// Setup the credentials
	switch config.Spec.Backend.Azure.Credentials.UseAzLogin {
	case "fallback":
		var azureClientId string = os.Getenv(config.Spec.Backend.Azure.Credentials.ClientIDEnvName)
		var azureClientSecret string = os.Getenv(config.Spec.Backend.Azure.Credentials.ClientSecretEnvName)

		if azureClientId != "" && azureClientSecret != "" {
			logrus.Info("configuring Azure backend using credentials from local environment")
			// Set env for terraform
			os.Setenv("ARM_CLIENT_ID", azureClientId)
			os.Setenv("ARM_CLIENT_SECRET", azureClientSecret)
			os.Setenv("ARM_TENANT_ID", config.Spec.Backend.Azure.TenantID)

			// Set env for azidentity module
			// TODO: split this out into seperate config area (support for seperate principal for backend vs deployments)
			os.Setenv("AZURE_CLIENT_ID", azureClientId)
			os.Setenv("AZURE_CLIENT_SECRET", azureClientSecret)
			os.Setenv("AZURE_TENANT_ID", config.Spec.Backend.Azure.TenantID)
		} else {
			logrus.Info("nominated azure env variables are not configured, falling back to az cli")
		}

	case "true":
		logrus.Info("assuming that az cli session exists for an Azure User (this will fail if the logged in identity is a service principal")
	default:
		return fmt.Errorf("unknown value for useAzLogin supplied: %v", config.Spec.Backend.Azure.Credentials.UseAzLogin)
	}

	// Create the storage container if not exists
	if config.Spec.Backend.Azure.AutoCreateStorage == "true" {
		err := createBackendStorage(config)
		if err != nil {
			return err
		}
	}

	// Setup the environment for backend config
	options := []string{}
	options = append(options, fmt.Sprintf("-backend-config=subscription_id=%v", config.Spec.Backend.Azure.SubscriptionID))
	options = append(options, fmt.Sprintf("-backend-config=resource_group_name=%v", config.Spec.Backend.Azure.StorageAccountRg))
	options = append(options, fmt.Sprintf("-backend-config=storage_account_name=%v", config.Spec.Backend.Azure.StorageAccountName))
	options = append(options, fmt.Sprintf("-backend-config=container_name=%v", config.Spec.Backend.Azure.StorageAccountContainer))
	options = append(options, fmt.Sprintf("-backend-config=key=%v", config.Spec.Backend.Azure.StateFileName))

	os.Setenv("TF_CLI_ARGS_init", strings.Join(options, " "))

	return nil
}

func createBackendStorage(config *model.Config) error {

	// TODO add ability to create astorage account as well as the container

	url := fmt.Sprintf("https://%v.blob.core.windows.net/", config.Spec.Backend.Azure.StorageAccountName)
	ctx := context.Background()

	credential, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return fmt.Errorf("Invalid credentials with error: " + err.Error())
	}

	serviceClient, err := azblob.NewServiceClient(url, credential, nil)
	if err != nil {
		return fmt.Errorf("Invalid credentials with error: " + err.Error())
	}

	// Check to see if the contianer exists
	containerClient, err := serviceClient.NewContainerClient(config.Spec.Backend.Azure.StorageAccountContainer)
	if err != nil {
		return err
	}

	var ops azblob.ListContainersOptions
	ops.Prefix = &config.Spec.Backend.Azure.StorageAccountContainer
	ops.Include.Metadata = true

	pager := serviceClient.ListContainers(&ops)
	if pager.Err() != nil {
		return errors.New(pager.Err().Error())
	}
	pager.NextPage(context.TODO())
	resp := pager.PageResponse()

	if len(resp.ContainerItems) != 1 {
		logrus.Infof("creating storage account container: %v", config.Spec.Backend.Azure.StorageAccountContainer)
		_, err = containerClient.Create(ctx, nil)
		if err != nil {
			return err
		}
	} else {
		logrus.Infof("storage account container exists: %v", config.Spec.Backend.Azure.StorageAccountContainer)
	}

	return nil

}
