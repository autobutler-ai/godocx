package docx

import (
	"encoding/xml"
)

var numberingAttrs = map[string]string{
	"xmlns:mc":     "http://schemas.openxmlformats.org/markup-compatibility/2006",
	"xmlns:o":      "urn:schemas-microsoft-com:office:office",
	"xmlns:r":      "http://schemas.openxmlformats.org/officeDocument/2006/relationships",
	"xmlns:m":      "http://schemas.openxmlformats.org/officeDocument/2006/math",
	"xmlns:v":      "urn:schemas-microsoft-com:vml",
	"xmlns:wp":     "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing",
	"xmlns:w10":    "urn:schemas-microsoft-com:office:word",
	"xmlns:w":      "http://schemas.openxmlformats.org/wordprocessingml/2006/main",
	"xmlns:wne":    "http://schemas.microsoft.com/office/word/2006/wordml",
	"xmlns:sl":     "http://schemas.openxmlformats.org/schemaLibrary/2006/main",
	"xmlns:a":      "http://schemas.openxmlformats.org/drawingml/2006/main",
	"xmlns:pic":    "http://schemas.openxmlformats.org/drawingml/2006/picture",
	"xmlns:c":      "http://schemas.openxmlformats.org/drawingml/2006/chart",
	"xmlns:lc":     "http://schemas.openxmlformats.org/drawingml/2006/lockedCanvas",
	"xmlns:dgm":    "http://schemas.openxmlformats.org/drawingml/2006/diagram",
	"xmlns:wps":    "http://schemas.microsoft.com/office/word/2010/wordprocessingShape",
	"xmlns:wpg":    "http://schemas.microsoft.com/office/word/2010/wordprocessingGroup",
	"xmlns:w14":    "http://schemas.microsoft.com/office/word/2010/wordml",
	"xmlns:w15":    "http://schemas.microsoft.com/office/word/2012/wordml",
	"xmlns:w16":    "http://schemas.microsoft.com/office/word/2018/wordml",
	"xmlns:w16cex": "http://schemas.microsoft.com/office/word/2018/wordml/cex",
	"xmlns:w16cid": "http://schemas.microsoft.com/office/word/2016/wordml/cid",
	"xmlns=":       "http://schemas.microsoft.com/office/tasks/2019/documenttasks",
	"xmlns:cr":     "http://schemas.microsoft.com/office/comments/2020/reactions",
}

// This element specifies the contents of a main document part in a WordprocessingML document.
type Numbering struct {
	// Reference to the RootDoc
	Root *RootDoc

	// Elements
	AbstractNum *AbstractNum
	// Num         *Num
	RelativePath string // RelativePath is the path to the numbering file within the document package.
}

// MarshalXML implements the xml.Marshaler interface for the Document type.
func (n Numbering) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	start.Name.Local = "w:numbering"

	for key, value := range numberingAttrs {
		attr := xml.Attr{Name: xml.Name{Local: key}, Value: value}
		start.Attr = append(start.Attr, attr)
	}

	err = e.EncodeToken(start)
	if err != nil {
		return err
	}

	if n.AbstractNum != nil {
		abstractNumElement := xml.StartElement{Name: xml.Name{Local: "w:abstractNum"}}
		if err = n.AbstractNum.MarshalXML(e, abstractNumElement); err != nil {
			return err
		}
	}

	// if n.Num != nil {
	// 	numElement := xml.StartElement{Name: xml.Name{Local: "w:num"}}
	// 	if err = n.Num.MarshalXML(e, numElement); err != nil {
	// 		return err
	// 	}
	// }

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (n *Numbering) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) (err error) {
	for {
		currentToken, err := decoder.Token()
		if err != nil {
			return err
		}

		switch elem := currentToken.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "abstractNum":
				abstractNum := NewAbstractNum()
				if err := decoder.DecodeElement(abstractNum, &elem); err != nil {
					return err
				}
				n.AbstractNum = abstractNum
			// case "num":
			// 	num := NewNum()
			// 	if err := decoder.DecodeElement(num, &elem); err != nil {
			// 		return err
			// 	}
			// 	n.Num = num
			default:
				if err = decoder.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			return nil
		}
	}

}
