package decimal128

import (
	"encoding/xml"
)

func (d Decimal) MarshalXML(xe *xml.Encoder, start xml.StartElement) (err error) {
	if (Decimal{} == d) {
		return nil
	}
	return nil
}

func (d Decimal) MarshalXMLAttr(xe *xml.Encoder, start xml.StartElement) (err error) {
	if (Decimal{} == d) {
		return nil
	}
	return nil
}

func (d *Decimal) UnmarshalXML(xd *xml.Decoder, start xml.StartElement) error {
	var v string
	neg := false

	err := xd.DecodeElement(&v, &start)
	if err != nil {
		return err
	}
	parse, _ := parseNumber(v, neg, false)
	*d = parse
	return nil

}

func (d *Decimal) UnmarshalXMLAttr(attr xml.Attr) error {
	neg := false
	parse, _ := parseNumber(attr.Value, neg, false)
	*d = parse
	return nil
}
