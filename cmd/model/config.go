package model

type Config struct {
	APIVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	Metadata   Metadata `yaml:"metadata"`
	Spec       Spec     `yaml:"spec"`
}
type Metadata struct {
	Name        string `yaml:"name"`
	Environment string `yaml:"environment"`
	Description string `yaml:"description"`
}
type Credentials struct {
	UseAzLogin          string `yaml:"useAzLogin"`
	FromEnvironment     string `yaml:"fromEnvironment"`
	ClientIDEnvName     string `yaml:"clientIdEnvName"`
	ClientSecretEnvName string `yaml:"clientSecretEnvName"`
}
type AzureBackend struct {
	AutoCreateStorage       string      `yaml:"autoCreateStorage"`
	StorageAccountRg        string      `yaml:"storageAccountRg"`
	StorageAccountName      string      `yaml:"storageAccountName"`
	StorageAccountContainer string      `yaml:"storageAccountContainer"`
	StateFileName           string      `yaml:"stateFileName"`
	SubscriptionID          string      `yaml:"subscriptionId"`
	TenantID                string      `yaml:"tenantId"`
	Credentials             Credentials `yaml:"credentials"`
}
type Backend struct {
	Azure AzureBackend `yaml:"azure"`
}
type Spec struct {
	Backend Backend `yaml:"backend"`
}
