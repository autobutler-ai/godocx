package docx

import (
	"encoding/xml"
)

// This element specifies the contents of a main document part in a WordprocessingML document.
type Level struct {
	// @jamesaorson TODO: Fill it out
}

func NewLevel() *Level {
	return &Level{}
}

// MarshalXML implements the xml.Marshaler interface for the Document type.
func (l Level) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	start.Name.Local = "w:lvl"

	err = e.EncodeToken(start)
	if err != nil {
		return err
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (l *Level) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) (err error) {
	for {
		currentToken, err := decoder.Token()
		if err != nil {
			return err
		}

		switch elem := currentToken.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "start":
				// @jamesaorson TODO: Implement start element
				fallthrough
			case "numFmt":
				// @jamesaorson TODO: Implement numFmt element
				fallthrough
			case "lvlText":
				// @jamesaorson TODO: Implement lvlText element
				fallthrough
			case "lvlJc":
				// @jamesaorson TODO: Implement lvlJc element
				fallthrough
			case "pPr":
				// @jamesaorson TODO: Implement pPr element
				fallthrough
			case "rPr":
				// @jamesaorson TODO: Implement rPr element
				fallthrough
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
