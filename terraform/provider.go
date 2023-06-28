package terraform

import (
	"context"
	"terraform-provider-mygithub/terraform/resources/repo"

	"github.com/google/go-github/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/oauth2"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"auth_token": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"mygithub_repo": repo.Resource(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"mygithub_repo": repo.Data(),
		},
		ConfigureContextFunc: configureClient,
	}
}

func configureClient(c context.Context, data *schema.ResourceData) (interface{}, diag.Diagnostics) {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: data.Get("auth_token").(string)},
	)
	tc := oauth2.NewClient(c, ts)

	client := github.NewClient(tc)

	return client, nil
}
