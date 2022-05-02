package backend

import (
	"fmt"

	"github.com/quantize-io/tfctl/cmd/model"
)

func InitBackend(config *model.Config) error {

	if (model.AzureBackend{}) != config.Spec.Backend.Azure {
		return initAzureBacked(config)
	}

	return fmt.Errorf("unsupported or no backend type provided")

}
