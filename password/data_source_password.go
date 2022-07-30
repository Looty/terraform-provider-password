package password

import (
	"context"
	"fmt"
	"github.com/sethvargo/go-password/password"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	//"github.com/sethvargo/go-password/password"
)

func dataSourcePassword() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePasswordRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"recipe": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"length": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Default:  6,
							ForceNew: true,
							ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
								v := val.(int)
								if v < 0 || v > 32 {
									errs = append(errs, fmt.Errorf("%q must be between 0 and 32 inclusive, got: %d", key, v))
								}
								return
							},
						},
						"num_digits": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Default:  0,
							ForceNew: true,
						},
						"num_symbols": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Default:  0,
							ForceNew: true,
						},
						"allow_upper": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
							ForceNew: true,
						},
						"allow_repeat": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
							ForceNew: true,
						},
					},
				},
			},
			"value": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourcePasswordRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type

	var diags diag.Diagnostics

	length := 0
	numDigits := 0
	numSymbols := 0
	allowUpper := false
	allowRepeat := false

	if recipe, ok := d.GetOk("recipe"); ok {
		var recipeParams = recipe.(*schema.Set)
		for _, f := range recipeParams.List() {
			length = f.(map[string]interface{})["length"].(int)
			numDigits = f.(map[string]interface{})["num_digits"].(int)
			numSymbols = f.(map[string]interface{})["num_symbols"].(int)
			allowUpper = f.(map[string]interface{})["allow_upper"].(bool)
			allowRepeat = f.(map[string]interface{})["allow_repeat"].(bool)
		}
	}

	res, err := password.Generate(length, numDigits, numSymbols, allowUpper, allowRepeat)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("value", res); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
