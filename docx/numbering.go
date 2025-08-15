package docx

import (
	"github.com/nbio/xml"

	"github.com/autobutler-ai/godocx/wml/stypes"
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
	"xmlns:cr":     "http://schemas.microsoft.com/office/comments/2020/reactions",
	"xmlns=":       "http://schemas.microsoft.com/office/tasks/2019/documenttasks",
}

// This element specifies the contents of a main document part in a WordprocessingML document.
type Numbering struct {
	XMLName          xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main w:numbering"`
	Namespace_mc     string   `xml:"xmlns:mc,attr"`
	Namespace_o      string   `xml:"xmlns:o,attr"`
	Namespace_r      string   `xml:"xmlns:r,attr"`
	Namespace_m      string   `xml:"xmlns:m,attr"`
	Namespace_v      string   `xml:"xmlns:v,attr"`
	Namespace_wp     string   `xml:"xmlns:wp,attr"`
	Namespace_w10    string   `xml:"xmlns:w10,attr"`
	Namespace_w      string   `xml:"xmlns:w,attr"`
	Namespace_wne    string   `xml:"xmlns:wne,attr"`
	Namespace_sl     string   `xml:"xmlns:sl,attr"`
	Namespace_a      string   `xml:"xmlns:a,attr"`
	Namespace_pic    string   `xml:"xmlns:pic,attr"`
	Namespace_c      string   `xml:"xmlns:c,attr"`
	Namespace_lc     string   `xml:"xmlns:lc,attr"`
	Namespace_dgm    string   `xml:"xmlns:dgm,attr"`
	Namespace_wps    string   `xml:"xmlns:wps,attr"`
	Namespace_wpg    string   `xml:"xmlns:wpg,attr"`
	Namespace_w14    string   `xml:"xmlns:w14,attr"`
	Namespace_w15    string   `xml:"xmlns:w15,attr"`
	Namespace_w16    string   `xml:"xmlns:w16,attr"`
	Namespace_w16cex string   `xml:"xmlns:w16cex,attr"`
	Namespace_w16cid string   `xml:"xmlns:w16cid,attr"`
	Namespace_cr     string   `xml:"xmlns:cr,attr"`
	AbstractNum      *AbstractNum
	Num              *Num
	RelativePath     string `xml:"-"` // RelativePath is the path to the numbering file within the document package.
}

func NewNumbering(abstractNumId int, isOrdered bool) *Numbering {
	return &Numbering{
		XMLName:          xml.Name{Local: "numbering", Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main"},
		AbstractNum:      NewAbstractNum(abstractNumId, isOrdered),
		Namespace_mc:     "http://schemas.openxmlformats.org/markup-compatibility/2006",
		Namespace_o:      "urn:schemas-microsoft-com:office:office",
		Namespace_r:      "http://schemas.openxmlformats.org/officeDocument/2006/relationships",
		Namespace_m:      "http://schemas.openxmlformats.org/officeDocument/2006/math",
		Namespace_v:      "urn:schemas-microsoft-com:vml",
		Namespace_wp:     "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing",
		Namespace_w10:    "urn:schemas-microsoft-com:office:word",
		Namespace_w:      "http://schemas.openxmlformats.org/wordprocessingml/2006/main",
		Namespace_wne:    "http://schemas.microsoft.com/office/word/2006/wordml",
		Namespace_sl:     "http://schemas.openxmlformats.org/schemaLibrary/2006/main",
		Namespace_a:      "http://schemas.openxmlformats.org/drawingml/2006/main",
		Namespace_pic:    "http://schemas.openxmlformats.org/drawingml/2006/picture",
		Namespace_c:      "http://schemas.openxmlformats.org/drawingml/2006/chart",
		Namespace_lc:     "http://schemas.openxmlformats.org/drawingml/2006/lockedCanvas",
		Namespace_dgm:    "http://schemas.openxmlformats.org/drawingml/2006/diagram",
		Namespace_wps:    "http://schemas.microsoft.com/office/word/2010/wordprocessingShape",
		Namespace_wpg:    "http://schemas.microsoft.com/office/word/2010/wordprocessingGroup",
		Namespace_w14:    "http://schemas.microsoft.com/office/word/2010/wordml",
		Namespace_w15:    "http://schemas.microsoft.com/office/word/2012/wordml",
		Namespace_w16:    "http://schemas.microsoft.com/office/word/2018/wordml",
		Namespace_w16cex: "http://schemas.microsoft.com/office/word/2018/wordml/cex",
		Namespace_w16cid: "http://schemas.microsoft.com/office/word/2016/wordml/cid",
		Namespace_cr:     "http://schemas.microsoft.com/office/comments/2020/reactions",
	}
}

// This element specifies the contents of a main document part in a WordprocessingML document.
type AbstractNum struct {
	XMLName       xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main w:abstractNum"`
	AbstractNumId int      `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main w:abstractNumId,attr"`
	Levels        []Level  `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main w:lvl"`
}

func NewAbstractNum(abstractNumId int, isOrdered bool) *AbstractNum {
	if isOrdered {
		return &AbstractNum{
			AbstractNumId: abstractNumId,
			Levels:        orderedLevels,
		}
	} else {
		return &AbstractNum{
			AbstractNumId: abstractNumId,
			Levels:        unorderedLevels,
		}
	}
}

type Num struct {
	XMLName       xml.Name      `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main w:num"`
	NumId         string        `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main w:numId,attr"`
	AbstractNumId AbstractNumId `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main w:abstractNumId"`
}

type AbstractNumId struct {
	XMLName       xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main w:abstractNumId"`
	AbstractNumId int      `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main w:val,attr"`
}

type Level struct {
	XMLName   xml.Name  `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main w:lvl"`
	Level     int       `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main w:ilvl,attr"`
	Start     Start     `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main w:start"`
	NumFmt    NumFmt    `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main w:numFmt"`
	LevelText LevelText `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main w:lvlText"`
}

type Start struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main w:start"`
	Val     int      `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main w:val,attr"`
}

type LevelText struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main w:lvlText"`
	Val     string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main w:val,attr"`
}

type LevelJustification struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main w:lvlJc"`
	Val     string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main w:val,attr"`
}

type NumFmt struct {
	XMLName xml.Name      `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main w:numFmt"`
	Val     stypes.NumFmt `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main w:val,attr"`
}

var OrderedLevelText = LevelText{Val: "%1."}
var orderedLevels = []Level{
	{
		Level:     0,
		Start:     Start{Val: 1},
		NumFmt:    NumFmt{Val: stypes.NumFmtDecimal},
		LevelText: OrderedLevelText,
	},
	{
		Level:     1,
		Start:     Start{Val: 1},
		NumFmt:    NumFmt{Val: stypes.NumFmtDecimal},
		LevelText: OrderedLevelText,
	},
	{
		Level:     2,
		Start:     Start{Val: 1},
		NumFmt:    NumFmt{Val: stypes.NumFmtDecimal},
		LevelText: OrderedLevelText,
	},
	{
		Level:     3,
		Start:     Start{Val: 1},
		NumFmt:    NumFmt{Val: stypes.NumFmtDecimal},
		LevelText: OrderedLevelText,
	},
	{
		Level:     4,
		Start:     Start{Val: 1},
		NumFmt:    NumFmt{Val: stypes.NumFmtDecimal},
		LevelText: OrderedLevelText,
	},
	{
		Level:     5,
		Start:     Start{Val: 1},
		NumFmt:    NumFmt{Val: stypes.NumFmtDecimal},
		LevelText: OrderedLevelText,
	},
	{
		Level:     6,
		Start:     Start{Val: 1},
		NumFmt:    NumFmt{Val: stypes.NumFmtDecimal},
		LevelText: OrderedLevelText,
	},
	{
		Level:     7,
		Start:     Start{Val: 1},
		NumFmt:    NumFmt{Val: stypes.NumFmtDecimal},
		LevelText: OrderedLevelText,
	},
	{
		Level:     8,
		Start:     Start{Val: 1},
		NumFmt:    NumFmt{Val: stypes.NumFmtDecimal},
		LevelText: OrderedLevelText,
	},
}

var UnorderedLevelText = LevelText{Val: "●"}
var unorderedLevels = []Level{
	{
		Level:     0,
		Start:     Start{Val: 1},
		NumFmt:    NumFmt{Val: stypes.NumFmtBullet},
		LevelText: UnorderedLevelText,
	},
	{
		Level:     1,
		Start:     Start{Val: 1},
		NumFmt:    NumFmt{Val: stypes.NumFmtBullet},
		LevelText: UnorderedLevelText,
	},
	{
		Level:     2,
		Start:     Start{Val: 1},
		NumFmt:    NumFmt{Val: stypes.NumFmtBullet},
		LevelText: UnorderedLevelText,
	},
	{
		Level:     3,
		Start:     Start{Val: 1},
		NumFmt:    NumFmt{Val: stypes.NumFmtBullet},
		LevelText: UnorderedLevelText,
	},
	{
		Level:     4,
		Start:     Start{Val: 1},
		NumFmt:    NumFmt{Val: stypes.NumFmtBullet},
		LevelText: UnorderedLevelText,
	},
	{
		Level:     5,
		Start:     Start{Val: 1},
		NumFmt:    NumFmt{Val: stypes.NumFmtBullet},
		LevelText: UnorderedLevelText,
	},
	{
		Level:     6,
		Start:     Start{Val: 1},
		NumFmt:    NumFmt{Val: stypes.NumFmtBullet},
		LevelText: UnorderedLevelText,
	},
	{
		Level:     7,
		Start:     Start{Val: 1},
		NumFmt:    NumFmt{Val: stypes.NumFmtBullet},
		LevelText: UnorderedLevelText,
	},
	{
		Level:     8,
		Start:     Start{Val: 1},
		NumFmt:    NumFmt{Val: stypes.NumFmtBullet},
		LevelText: UnorderedLevelText,
	},
}
