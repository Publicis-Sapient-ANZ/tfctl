package backend

import (
	"fmt"

	"github.com/quantize-io/tfctl/cmd/model"
)

func InitBackend(config *model.Config) error {

	switch config.Spec.Backend.Type {
	case "azure":
		return initAzureBacked(config)
	default:
		return fmt.Errorf("unsupported backend: %v", config.Spec.Backend.Type)
	}

}
