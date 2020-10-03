package dash

import (
	"context"
	"encoding/xml"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePanel() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePanelCreate,
		ReadContext:   resourcePanelRead,
		UpdateContext: resourcePanelUpdate,
		DeleteContext: resourcePanelDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"panel_xml": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"chart": chartSchema(),
		},
	}
}

func resourcePanelCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var input_xml string
	log.Println("In Panel Create")

	panel_name := d.Get("name").(string)

	input_xml, _ = GeneratePanelXML(d)

	d.Set("panel_xml", input_xml)
	d.SetId(panel_name)

	return nil

}

func resourcePanelRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func resourcePanelUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var input_xml string
	log.Println("In Panel Update")
	log.Println("panel resource data", d)

	input_xml, _ = GeneratePanelXML(d)

	d.Set("panel_xml", input_xml)

	return nil
}

func resourcePanelDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	d.SetId("")

	return diags
}

func GeneratePanelXML(d *schema.ResourceData) (string, error) {
	var panelXml PanelXML
	charts := d.Get("chart").(*schema.Set)
	log.Println("Charts", charts)
	log.Println("Charts Interface", d.Get("chart"))
	panelXml = PanelXML{
		Charts: GetChartList(charts),
	}

	genxml, err := xml.MarshalIndent(panelXml, "", "  ")
	if err != nil {
		return "", err
	}

	log.Println("Generated panel xml string")
	log.Println(string(genxml))

	return string(genxml), nil
}
