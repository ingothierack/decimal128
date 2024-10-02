package decimal128

import (
	"encoding/xml"
	"errors"
	"strings"
)

var (
	errInvalidMarshal     = errors.New("invalid value to marshal")
	errInvalidMarshalAttr = errors.New("invalid value to marshal attribute")
)

func (d Decimal) MarshalXML(xe *xml.Encoder, start xml.StartElement) error {
	if d.isSpecial() {
		return errInvalidMarshal
	}
	// Write the start element
	if err := xe.EncodeToken(start); err != nil {
		return err
	}
	// v := d.formatDecimal()
	v, err := d.MarshalText()
	if err != nil {
		return err
	}

	if err := xe.EncodeToken(xml.CharData(v)); err != nil {
		return err
	}

	if err := xe.EncodeToken(start.End()); err != nil {
		return err
	}
	return xe.Flush()
}

func (d Decimal) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if d.isSpecial() {
		return xml.Attr{}, errInvalidMarshalAttr
	}

	// v := d.formatDecimal()
	v, err := d.MarshalText()
	if err != nil {
		return xml.Attr{}, err
	}

	return xml.Attr{
		Name:  name,
		Value: string(v),
	}, nil
}

func (d *Decimal) UnmarshalXML(xd *xml.Decoder, start xml.StartElement) error {
	var v string
	if err := xd.DecodeElement(&v, &start); err != nil {
		return err
	}
	return d.unmarshalString(v)
}

func (d *Decimal) UnmarshalXMLAttr(attr xml.Attr) error {
	return d.unmarshalString(attr.Value)
}

func (d *Decimal) unmarshalString(value string) error {
	value = strings.TrimSpace(value)

	if value == "" {
		return nil
	}

	neg := false
	i := 0

	switch value[0] {
	case '+':
		i = 1
	case '-':
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

func (d Decimal) formatDecimal() []byte {
	var digs digits
	d.digits(&digs)

	prec := 0
	if digs.ndig != 0 {
		prec = digs.ndig - 1
	}

	exp := digs.exp + prec

	var v []byte
	if exp < -6 || exp >= 20 {
		v = digs.fmtE(nil, prec, 0, false, false, false, false, false, false, 'e')
	} else {
		prec = 0
		if digs.exp < 0 {
			prec = -digs.exp
		}
		v = digs.fmtF(nil, prec, 0, false, false, false, false, false)
	}
	return v
}
