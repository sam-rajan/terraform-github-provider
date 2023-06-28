package repo

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func read(c context.Context, d *schema.ResourceData, i interface{}) diag.Diagnostics {
	client := i.(*github.Client)

	var diags diag.Diagnostics

	user, response, err := client.Users.Get(c, "")

	if err != nil || response.StatusCode != 200 {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Detail:   err.Error(),
			Summary:  "Failed to find authenticated User",
		})
		return diags
	}

	repoName := d.Get("name").(string)
	if d.Id() != "" {
		repoName = d.Id()
	}
	repo, response, err := client.Repositories.Get(c, *user.Login, repoName)

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Detail:   err.Error(),
			Summary:  "Github server returned an error response",
		})
		return diags
	}

	if response.StatusCode != 200 && response.StatusCode != 404 {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Detail:   fmt.Sprintf("Request Failed with http code %d", response.StatusCode),
			Summary:  "Failed fetch github repo information",
		})
		return diags
	}

	if response.StatusCode == 404 {
		d.SetId("")
	}

	d.Set("name", repo.Name)
	d.Set("description", repo.Description)
	d.Set("homepage", repo.Homepage)
	d.Set("owner", *user.Login)
	d.Set("url", repo.URL)

	d.SetId(*repo.Name)
	return nil
}
