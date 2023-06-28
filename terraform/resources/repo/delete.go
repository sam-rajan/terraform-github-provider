package repo

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func delete(c context.Context, d *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := i.(*github.Client)
	response, err := client.Repositories.Delete(c,
		d.Get("owner").(string),
		d.Get("name").(string))

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Detail:   err.Error(),
			Summary:  "Github server returned an error response",
		})
		return diags
	}

	if response.StatusCode != 204 {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Detail:   fmt.Sprintf("Request Failed with http code %d", response.StatusCode),
			Summary:  "Failed to delete github repo",
		})
		return diags
	}

	d.SetId("")
	return nil
}
