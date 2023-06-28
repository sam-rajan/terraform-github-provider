package repo

import (
	"terraform-provider-mygithub/terraform/validator"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Resource() *schema.Resource {
	return &schema.Resource{
		CreateContext: create,
		UpdateContext: update,
		DeleteContext: delete,
		ReadContext:   read,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(1 * time.Minute),
			Update: schema.DefaultTimeout(1 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validator.ValidateName,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"homepage": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"has_issues": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"has_projects": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"has_wiki": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"has_discussions": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"auto_init": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"gitignore_template": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"license_template": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"allow_squash_merge": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"allow_merge_commit": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"allow_rebase_merge": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"allow_auto_merge": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"delete_branch_on_merge": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"has_downloads": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"is_template": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"owner": {
				Type:     schema.TypeString,
				Required: false,
				Optional: false,
				Computed: true,
			},
			"url": {
				Type:     schema.TypeString,
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func Data() *schema.Resource {
	return &schema.Resource{
		ReadContext: read,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"homepage": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"owner": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"url": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
