package repo

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func create(c context.Context, d *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := i.(*github.Client)

	user, response, err := client.Users.Get(c, "")

	tflog.Info(c, fmt.Sprintf("Found authenticated user %s", *user.Login))
	if err != nil || response.StatusCode != 200 {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Detail:   err.Error(),
			Summary:  "Failed to fetch authenticated user info from Github",
		})
		return diags
	}

	repoName := d.Get("name").(string)
	description := d.Get("description").(string)
	homepage := d.Get("homepage").(string)

	repo := &github.Repository{
		Name:        &repoName,
		Description: &description,
		Homepage:    &homepage,
	}

	repo, response, err = client.Repositories.Create(c, "", repo)

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Detail:   err.Error(),
			Summary:  "Github server returned an error response",
		})
		return diags
	}

	if response.StatusCode != 201 {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Detail:   fmt.Sprintf("Request Failed with http code %d", response.StatusCode),
			Summary:  "Failed creating github repo",
		})
		return diags
	}

	d.Set("owner", *user.Login)
	d.SetId(*repo.Name)
	d.Set("url", repo.URL)
	return nil
}
