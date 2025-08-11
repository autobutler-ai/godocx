package docx

import (
	"encoding/xml"
)

// This element specifies the contents of a main document part in a WordprocessingML document.
type AbstractNum struct {
	Levels []*Level
}

func NewAbstractNum() *AbstractNum {
	return &AbstractNum{
		Levels: make([]*Level, 0),
	}
}

// MarshalXML implements the xml.Marshaler interface for the Document type.
func (a AbstractNum) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	start.Name.Local = "w:abstractNum"

	err = e.EncodeToken(start)
	if err != nil {
		return err
	}

	for _, level := range a.Levels {
		levelElement := xml.StartElement{Name: xml.Name{Local: "w:lvl"}}
		if err = level.MarshalXML(e, levelElement); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (a *AbstractNum) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) (err error) {
	for {
		currentToken, err := decoder.Token()
		if err != nil {
			return err
		}

		a.Levels = make([]*Level, 0)
		switch elem := currentToken.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "lvl":
				level := NewLevel()
				if err := decoder.DecodeElement(level, &elem); err != nil {
					return err
				}
				a.Levels = append(a.Levels, level)
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
