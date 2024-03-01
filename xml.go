package decimal128

import (
	"encoding/xml"
	"errors"
)

func (d Decimal) MarshalXML(xe *xml.Encoder, start xml.StartElement) error {
	if d.isSpecial() || d.IsNaN() || d.IsZero() {
		return errors.New("invalid value to marshal")
	}
	v := d.String()
	return xe.EncodeElement(v, start)
}

func (d Decimal) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if d.isSpecial() || d.IsNaN() || d.IsZero() {
		return xml.Attr{}, errors.New("invalid value to marshal")
	}
	v := d.String()
	attr := xml.Attr{
		Name:  name,
		Value: v,
	}
	return attr, nil
}

func (d *Decimal) UnmarshalXML(xd *xml.Decoder, start xml.StartElement) error {
	var v string
	neg := false

	err := xd.DecodeElement(&v, &start)
	if err != nil {
		return err
	}
	parse, err := parseNumber(v, neg, false)
	if err != nil {
		return err
	}
	*d = parse
	return nil

}

func (d *Decimal) UnmarshalXMLAttr(attr xml.Attr) error {
	neg := false
	parse, err := parseNumber(attr.Value, neg, false)
	if err != nil {
		return err
	}
	*d = parse
	return nil
}
