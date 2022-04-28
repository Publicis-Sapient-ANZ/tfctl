package model

type Config struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name        string `yaml:"name"`
		Description string `yaml:"description"`
		Environment string `yaml:"environment"`
	} `yaml:"metadata"`
	Spec struct {
		Backend struct {
			Type                    string `yaml:"type"`
			AutoCreateStorage       string `yaml:"autoCreateStorage"`
			StorageAccountRG        string `yaml:"storageAccountRg"`
			StorageAccountName      string `yaml:"storageAccountName"`
			StorageAccountContainer string `yaml:"storageAccountContainer"`
			StateFileName           string `yaml:"stateFileName"`
			SubscriptionID          string `yaml:"subscriptionId"`
			TenantID                string `yaml:"tenantId"`
			Credentials             struct {
				UseAzLogin          string `yaml:"useAzLogin"`
				FromEnvironment     string `yaml:"fromEnvironment"`
				ClientIDEnvName     string `yaml:"clientIdEnvName"`
				ClientSecretEnvName string `yaml:"clientSecretEnvName"`
			} `yaml:"credentials"`
		} `yaml:"backend"`
	} `yaml:"spec"`
}
