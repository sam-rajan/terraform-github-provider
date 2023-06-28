package validator

import (
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

func ValidateName(val interface{}, path cty.Path) diag.Diagnostics {
	name := val.(string)

	var diags diag.Diagnostics

	if len(name) > 20 {
		diags = append(diags, diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Validation Error",
			Detail:        "The 'name' attribute must have max length of 20 characters.",
			AttributePath: path,
		})
	}

	return diags
}
