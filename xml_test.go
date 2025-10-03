package decimal128

import (
	"bytes"
	"encoding/xml"
	"strings"
	"testing"
)

func TestDecimalMarshalXML(t *testing.T) {
	t.Parallel()

	initDecimalValues()

	for _, val := range decimalValues {
		if val.form != regularForm {
			continue
		}

		decval := val.Decimal()
		buf := &bytes.Buffer{}
		xe := xml.NewEncoder(buf)
		start := xml.StartElement{Name: xml.Name{Local: "test"}}

		err := decval.MarshalXML(xe, start)

		if err != nil {
			t.Errorf("%v.MarshalXML() = (%s, %v), want (%s, <nil>", val, buf.String(), err, buf.String())
		}

		var resval Decimal
		xd := xml.NewDecoder(strings.NewReader(buf.String()))
		tok, err := xd.Token()
		if err != nil {
			t.Errorf("xml.Decoder.Token() = error %v, want <nil>", err)
			continue
		}

		start2, ok := tok.(xml.StartElement)
		if !ok {
			t.Errorf("Expected xml.StartElement, got %T", tok)
			continue
		}

		err = resval.UnmarshalXML(xd, start2)

		if !resval.Equal(decval) || err != nil {
			t.Errorf("Decimal.UnmarshalXML(%s) = (%v, %v), want (%v, <nil>)", buf.String(), resval, err, decval)
		}
	}
}

func TestDecimalMarshalXMLAttr(t *testing.T) {
	t.Parallel()

	initDecimalValues()

	for _, val := range decimalValues {
		if val.form != regularForm {
			continue
		}

		decval := val.Decimal()
		attr, err := decval.MarshalXMLAttr(xml.Name{Local: "value"})

		if err != nil {
			t.Errorf("%v.MarshalXMLAttr() = (%s, %v), want (%s, <nil>", val, attr.Value, err, attr.Value)
		}

		var resval Decimal
		err = resval.UnmarshalXMLAttr(attr)

		if !resval.Equal(decval) || err != nil {
			t.Errorf("Decimal.UnmarshalXMLAttr(%s) = (%v, %v), want (%v, <nil>)", attr.Value, resval, err, decval)
		}
	}
}

func TestDecimalUnmarshalXML(t *testing.T) {
	t.Parallel()

	for val, num := range textValues {
		xmlFragment := "<test>" + val + "</test>"
		var res Decimal
		xd := xml.NewDecoder(strings.NewReader(xmlFragment))
		tok, err := xd.Token()
		if err != nil {
			t.Errorf("xml.Decoder.Token() for %s = error %v", val, err)
			continue
		}

		start, ok := tok.(xml.StartElement)
		if !ok {
			t.Errorf("Expected xml.StartElement for %s, got %T", val, tok)
			continue
		}

		err = res.UnmarshalXML(xd, start)

		if num.isInf() || num.IsNaN() || strings.Contains(val, "_") {
			if err == nil {
				t.Errorf("Decimal.UnmarshalXML(%s) = (0, <nil>), want (%v, cannot unmarshal)", val, res)
			}
		} else if !res.Equal(num) || err != nil {
			t.Errorf("Decimal.UnmarshalXML(%s) = (%v, %v), want (%v, <nil>)", val, res, err, num)
		}
	}

	num := New(123, -1)
	res := num
	xmlFragment := "<test></test>"
	xd := xml.NewDecoder(strings.NewReader(xmlFragment))
	tok, _ := xd.Token()
	start, _ := tok.(xml.StartElement)
	err := res.UnmarshalXML(xd, start)
	if !res.Equal(num) || err != nil {
		t.Errorf("Decimal.UnmarshalXML(<empty>) = (%v, %v), want (%v, <nil>)", res, err, num)
	}

	res = num
	xmlFragment = "<test>   </test>"
	xd = xml.NewDecoder(strings.NewReader(xmlFragment))
	tok, _ = xd.Token()
	start, _ = tok.(xml.StartElement)
	err = res.UnmarshalXML(xd, start)
	if !res.Equal(num) || err != nil {
		t.Errorf("Decimal.UnmarshalXML(<whitespace>) = (%v, %v), want (%v, <nil>)", res, err, num)
	}
}

func TestDecimalUnmarshalXMLAttr(t *testing.T) {
	t.Parallel()

	for val, num := range textValues {
		var res Decimal
		attr := xml.Attr{
			Name:  xml.Name{Local: "value"},
			Value: val,
		}

		err := res.UnmarshalXMLAttr(attr)

		if num.isInf() || num.IsNaN() || strings.Contains(val, "_") {
			if err == nil {
				t.Errorf("Decimal.UnmarshalXMLAttr(%s) = (0, <nil>), want (%v, cannot unmarshal)", val, res)
			}
		} else if !res.Equal(num) || err != nil {
			t.Errorf("Decimal.UnmarshalXMLAttr(%s) = (%v, %v), want (%v, <nil>)", val, res, err, num)
		}
	}

	num := New(123, -1)
	res := num
	attr := xml.Attr{
		Name:  xml.Name{Local: "value"},
		Value: "",
	}
	err := res.UnmarshalXMLAttr(attr)
	if !res.Equal(num) || err != nil {
		t.Errorf("Decimal.UnmarshalXMLAttr(<empty>) = (%v, %v), want (%v, <nil>)", res, err, num)
	}

	res = num
	attr = xml.Attr{
		Name:  xml.Name{Local: "value"},
		Value: "   ",
	}
	err = res.UnmarshalXMLAttr(attr)
	if !res.Equal(num) || err != nil {
		t.Errorf("Decimal.UnmarshalXMLAttr(<whitespace>) = (%v, %v), want (%v, <nil>)", res, err, num)
	}
}

func FuzzDecimalUnmarshalXML(f *testing.F) {
	f.Add("123456.789e10")

	f.Fuzz(func(t *testing.T, data string) {
		t.Parallel()

		xmlFragment := "<test>" + data + "</test>"
		var dec Decimal
		xd := xml.NewDecoder(strings.NewReader(xmlFragment))
		tok, err := xd.Token()
		if err != nil {
			return
		}

		start, ok := tok.(xml.StartElement)
		if !ok {
			return
		}

		dec.UnmarshalXML(xd, start)
	})
}

func FuzzDecimalUnmarshalXMLAttr(f *testing.F) {
	f.Add("123456.789e10")

	f.Fuzz(func(t *testing.T, data string) {
		t.Parallel()

		var dec Decimal
		attr := xml.Attr{
			Name:  xml.Name{Local: "value"},
			Value: data,
		}
		dec.UnmarshalXMLAttr(attr)
	})
}
