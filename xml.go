package decimal128

import (
	"encoding/xml"
	"errors"
	"strings"
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

	value := strings.TrimSpace(v)
	l := len(value)

	if l == 0 {
		return nil
	}

	i := 0
	if value[0] == '+' {
		i = 1
	} else if value[0] == '-' {
		neg = true
		i = 1
	}

	parse, err := parseNumber(value[i:], neg, false)
	if err != nil {
		return err
	}
	*d = parse
	return nil

}

func (d *Decimal) UnmarshalXMLAttr(attr xml.Attr) error {
	neg := false

	value := strings.TrimSpace(attr.Value)
	l := len(value)

	if l == 0 {
		return nil
	}

	i := 0
	if value[0] == '+' {
		i = 1
	} else if value[0] == '-' {
		neg = true
		i = 1
	}

	parse, err := parseNumber(value[i:], neg, false)
	if err != nil {
		return err
	}

	*d = parse
	return nil
}
