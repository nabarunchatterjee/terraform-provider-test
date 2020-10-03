package dash

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func chartSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"title": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"options": {
					Type:     schema.TypeMap,
					Optional: true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
			},
		},
	}
}

func GetChartList(s *schema.Set) *[]Chart {
	var charts []Chart

	log.Println("number of charts: ", s.Len())

	for _, chart := range s.List() {
		chart_elem := chart.(map[string]interface{})

		log.Println("Chart elements")
		for key, value := range chart_elem {
			log.Println(key, value)
		}

		chart_struct_item := Chart{
			Title:   chart_elem["title"].(string),
			Options: GetOptions(chart_elem["options"].(map[string]interface{})),
		}
		charts = append(charts, chart_struct_item)
	}

	log.Println("charts", charts)

	return &charts
}

func GetOptions(ops map[string]interface{}) *[]Option {
	var options []Option

	for name, value := range ops {
		op := Option{
			Name:  name,
			Value: value.(string),
		}
		options = append(options, op)
	}
	return &options
}
