package dash

import "encoding/xml"

type PanelXML struct {
	XMLName xml.Name `xml:"panel"`
	Charts  *[]Chart `xml:"chart,omitempty"`
}

type Chart struct {
	XMLName xml.Name  `xml:"chart"`
	Title   string    `xml:"title,omitempty"`
	Options *[]Option `xml:"option,omitempty"`
}

type Option struct {
	XMLName xml.Name `xml:"option"`
	Name    string   `xml:"name,attr,omitempty"`
	Value   string   `xml:",chardata"`
}
