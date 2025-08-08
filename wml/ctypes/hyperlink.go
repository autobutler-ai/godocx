package ctypes

import (
	"encoding/xml"
)

type Hyperlink struct {
	XMLName  xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main hyperlink,omitempty"`
	ID       string   `xml:"http://schemas.openxmlformats.org/officeDocument/2006/relationships id,attr"`
	Run      *Run     `xml:"r,omitempty"`
	Children []ParagraphChild
}

func (h Hyperlink) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	start.Name.Local = "w:hyperlink"

	if h.ID != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:id"}, Value: h.ID})
	}

	err = e.EncodeToken(start)
	if err != nil {
		return err
	}

	// 1. Run
	if h.Run != nil {
		propsElement := xml.StartElement{Name: xml.Name{Local: "w:r"}}
		if err = e.EncodeElement(h.Run, propsElement); err != nil {
			return err
		}
	}

	// 2. Children
	for _, cElem := range h.Children {
		if cElem.Run != nil {
			if err = cElem.Run.MarshalXML(e, xml.StartElement{
				Name: xml.Name{Local: "w:r"},
			}); err != nil {
				return err
			}
		}

		if cElem.Link != nil {
			if err = e.EncodeElement(cElem.Link, xml.StartElement{
				Name: xml.Name{Local: "w:hyperlink"},
			}); err != nil {
				return err
			}
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (h *Hyperlink) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	// Decode attributes
	for _, attr := range start.Attr {
		if attr.Name.Local == "id" {
			h.ID = attr.Value
		}
	}

loop:
	for {
		currentToken, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := currentToken.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "r":
				r := NewRun()
				if err = d.DecodeElement(r, &elem); err != nil {
					return err
				}

				h.Run = r
			default:
				if err = d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break loop
		}
	}

	return nil
}
