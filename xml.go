package decimal128

import (
	"encoding/xml"
	"errors"
	"strings"
)

func (d Decimal) MarshalXML(xe *xml.Encoder, start xml.StartElement) error {
	var v []byte

	if d.isSpecial() {
		return errors.New("invalid value to marshal")
	}

	// Write the start element
	err := xe.EncodeToken(start)
	if err != nil {
		return err
	}

	var digs digits
	d.digits(&digs)

	prec := 0
	if digs.ndig != 0 {
		prec = digs.ndig - 1
	}

	exp := digs.exp + prec

	if exp < -6 || exp >= 20 {
		v = digs.fmtE(nil, prec, 0, false, false, false, false, false, false, 'e')
	} else {
		prec = 0
		if digs.exp < 0 {
			prec = -digs.exp
		}

		v = digs.fmtF(nil, prec, 0, false, false, false, false, false)
	}
	// Write the Decimal string
	err = xe.EncodeToken(xml.CharData(v))
	if err != nil {
		return err
	}

	// Write the end element
	err = xe.EncodeToken(start.End())
	if err != nil {
		return err
	}

	// Flush the encoder
	err = xe.Flush()
	if err != nil {
		return err
	}
	return nil
	// return xe.EncodeElement(v, start)
}

func (d Decimal) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	var v []byte

	if d.isSpecial() {
		return xml.Attr{}, errors.New("invalid value to marshal attribute")
	}

	var digs digits
	d.digits(&digs)

	prec := 0
	if digs.ndig != 0 {
		prec = digs.ndig - 1
	}

	exp := digs.exp + prec

	if exp < -6 || exp >= 20 {
		v = digs.fmtE(nil, prec, 0, false, false, false, false, false, false, 'e')
	} else {
		prec = 0
		if digs.exp < 0 {
			prec = -digs.exp
		}

		v = digs.fmtF(nil, prec, 0, false, false, false, false, false)
	}

	attr := xml.Attr{
		Name:  name,
		Value: string(v),
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

	if len(value) == 0 {
		return nil
	}

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

func (d *Decimal) UnmarshalXMLAttr(attr xml.Attr) error {
	neg := false

	value := strings.TrimSpace(attr.Value)

	if len(value) == 0 {
		return nil
	}

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
