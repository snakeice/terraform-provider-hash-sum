package hashsum

import (
	"context"
	"math"
	"strconv"
	"time"

	"github.com/cespare/xxhash/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{},
		DataSourcesMap: map[string]*schema.Resource{
			"hashsum": &schema.Resource{
				ReadContext: hashGet,
				Schema: map[string]*schema.Schema{
					"sum": {
						Type:     schema.TypeFloat,
						Optional: true,
					},
					"data": {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
	}
}

func hashGet(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	data := d.Get("data").(string)
	sum := xxhash.Sum64String(data)

	sumConversion := sum % math.MaxInt32

	var diags diag.Diagnostics

	if err := d.Set("sum", sumConversion); err != nil {
		return diag.Errorf("failed to set sum: %s", err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}
