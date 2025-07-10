package model

type TFState struct {
	Version           int               `json:"version"`
	Terraform_version string            `json:"terraform_version"`
	Serial            int               `json:"serial"`
	Lineage           string            `json:"lineage"`
	Outputs           map[string]string `json:"outputs"`
	Resources         []any             `json:"resources"`
}
