package repo

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func update(c context.Context, d *schema.ResourceData, client interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	if d.HasChange("name") || d.HasChange("description") || d.HasChange("homepage") {
		err := updateRepo(c, d, client)

		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Detail:   err.Error(),
				Summary:  "Github server returned an error response while trying to update",
			})
		}
	}
	return diags
}

func updateRepo(c context.Context, d *schema.ResourceData, i interface{}) error {
	client := i.(*github.Client)

	tflog.Info(c, "Getting Repo configuration from remote server")
	_, response, err := client.Repositories.Get(c, d.Get("owner").(string), d.Get("name").(string))

	if err != nil {
		return err
	}

	if response.StatusCode != 200 && response.StatusCode != 404 {
		return fmt.Errorf("repo fetch request failed with http code %d", response.StatusCode)
	}

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	homepage := d.Get("homepage").(string)

	repositories := &github.Repository{
		Name:        &name,
		Description: &description,
		Homepage:    &homepage,
	}

	tflog.Info(c, "Updating Repo with new configurations")

	_, response, err = client.Repositories.Edit(c, d.Get("owner").(string), d.Get("name").(string), repositories)

	if err != nil {
		return err
	}

	if response.StatusCode != 200 && response.StatusCode != 404 {
		return fmt.Errorf("update request failed with http code %d", response.StatusCode)
	}

	return nil
}
