package aws

import (
	"github.com/hashicorp/vault/logical"
	"github.com/hashicorp/vault/logical/framework"
)

func pathRoot() *framework.Path {
	return &framework.Path{
		Pattern: "root",
		Fields: map[string]*framework.FieldSchema{
			"access_key": &framework.FieldSchema{
				Type:        framework.TypeString,
				Description: "Access key with permission to create new keys.",
			},

			"secret_key": &framework.FieldSchema{
				Type:        framework.TypeString,
				Description: "Secret key with permission to create new keys.",
			},

			"region": &framework.FieldSchema{
				Type:        framework.TypeString,
				Description: "Region for API calls.",
			},
		},

		Callbacks: map[logical.Operation]framework.OperationFunc{
			logical.WriteOperation: pathRootWrite,
		},

		HelpSynopsis:    pathRootHelpSyn,
		HelpDescription: pathRootHelpDesc,
	}
}

func pathRootWrite(
	req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
	entry, err := logical.StorageEntryJSON("root", rootConfig{
		AccessKey: data.Get("access_key").(string),
		SecretKey: data.Get("secret_key").(string),
		Region:    data.Get("region").(string),
	})
	if err != nil {
		return nil, err
	}

	if err := req.Storage.Put(entry); err != nil {
		return nil, err
	}

	return nil, nil
}

type rootConfig struct {
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
	Region    string `json:"region"`
}

const pathRootHelpSyn = `
Configure the root credentials that are used to manage IAM.
`

const pathRootHelpDesc = `
Before doing anything, the AWS backend needs credentials that are able
to manage IAM policies, users, access keys, etc. This endpoint is used
to configure those credentials. They don't necessarilly need to be root
keys as long as they have permission to manage IAM.
`
