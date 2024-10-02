package decimal128

import (
	"bytes"
	"encoding/xml"
	"strings"
	"testing"
)

func TestDecimalMarshalXML(t *testing.T) {
	values := []struct {
		value int64
		scale int
	}{
		{471, 0},
		{0, 0},
		{238, -2},
		{-85414437, -8},
		{1123456789, -9},
		// add more values as needed
	}

	for _, v := range values {
		d := New(v.value, v.scale)
		buf := &bytes.Buffer{}
		xe := xml.NewEncoder(buf)
		start := xml.StartElement{Name: xml.Name{Local: "test"}}

		err := d.MarshalXML(xe, start)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
			continue
		}

		// Unmarshal the data back to its original form
		d2 := &Decimal{}
		xd := xml.NewDecoder(strings.NewReader(buf.String()))
		tok, err := xd.Token()
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
			continue
		}

		start2, ok := tok.(xml.StartElement)
		if !ok {
			t.Errorf("Expected a start element, got %T", tok)
			continue
		}

		err = d2.UnmarshalXML(xd, start2)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
			continue
		}

		// Compare the original and unmarshalled data
		if d.Cmp(*d2) != 0 {
			t.Errorf("Expected %v, got %v", d, *d2)
		}
	}
}

func TestDecimalMarshalXMLAttr(t *testing.T) {
	tests := []struct {
		value    Decimal
		expected string
	}{
		{New(471, 0), `<test value="471"/>`},
		{New(0, 0), `<test value="0"/>`},
		{New(238, -2), `<test value="2.38"/>`},
		{New(-85414437, -7), `<test value="-8.5414437"/>`},
		{New(1123456789, -9), `<test value="1.123456789"/>`},
		// add more tests as needed
	}

	for _, tt := range tests {
		// Marshal the Decimal into an XML attribute
		attr := xml.Attr{
			Name:  xml.Name{Local: "value"},
			Value: tt.value.String(),
		}

		// Create an XML string with the marshalled attribute
		xmlString := `<test ` + attr.Name.Local + `="` + attr.Value + `"/>`

		// Compare the XML string with the expected XML string
		if xmlString != tt.expected {
			t.Errorf("Expected %s, got %s", tt.expected, xmlString)
			continue
		}

		// Unmarshal the XML attribute back to a Decimal
		d := &Decimal{}
		err := d.UnmarshalXMLAttr(attr)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
			continue
		}

		// Compare the unmarshalled Decimal with the original Decimal
		if *d != tt.value {
			t.Errorf("Expected %v, got %v", tt.value, *d)
		}
	}
}

func TestDecimalMarshalUnmarshalXML(t *testing.T) {
	tests := []struct {
		xmlFragment string
		expected    Decimal
	}{
		{`<test>471</test>`, New(471, 0)},
		{`<test>0</test>`, New(0, 0)},
		{`<test>2.38</test>`, New(238, -2)},
		{`<test>-8.5414437</test>`, New(-85414437, -7)},
		{`<test>1.123456789</test>`, New(1123456789, -9)},
		// add more tests as needed
	}

	for _, tt := range tests {
		// Unmarshal the XML fragment into a Decimal
		d := &Decimal{}
		xd := xml.NewDecoder(strings.NewReader(tt.xmlFragment))
		tok, err := xd.Token()
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
			continue
		}

		start, ok := tok.(xml.StartElement)
		if !ok {
			t.Errorf("Expected a start element, got %T", tok)
			continue
		}

		err = d.UnmarshalXML(xd, start)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
			continue
		}

		// Compare the unmarshalled Decimal with the expected Decimal
		if *d != tt.expected {
			t.Errorf("Expected %v, got %v", tt.expected, *d)
			continue
		}

		// Marshal the Decimal back to an XML fragment
		buf := &bytes.Buffer{}
		xe := xml.NewEncoder(buf)
		err = d.MarshalXML(xe, start)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
			continue
		}

		// Compare the marshalled XML fragment with the original XML fragment
		if buf.String() != tt.xmlFragment {
			t.Errorf("Expected %s, got %s", tt.xmlFragment, buf.String())
		}
	}
}

func TestDecimalUnmarshalXMLAttr(t *testing.T) {
	tests := []struct {
		xmlFragment string
		expected    Decimal
	}{
		{`<test value="471" />`, New(471, 0)},
		{`<test value="0" />`, New(0, 0)},
		{`<test value="2.38" />`, New(238, -2)},
		{`<test value="-8.5414437" />`, New(-85414437, -7)},
		{`<test value="1.123456789" />`, New(1123456789, -9)},
		// add more tests as needed
	}

	for _, tt := range tests {
		// Unmarshal the XML fragment into a Decimal
		d := &Decimal{}
		xd := xml.NewDecoder(strings.NewReader(tt.xmlFragment))
		tok, err := xd.Token()
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
			continue
		}

		start, ok := tok.(xml.StartElement)
		if !ok {
			t.Errorf("Expected a start element, got %T", tok)
			continue
		}

		for _, attr := range start.Attr {
			if attr.Name.Local == "value" {
				err = d.UnmarshalXMLAttr(attr)
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
					continue
				}
			}
		}

		// Compare the unmarshalled Decimal with the expected Decimal
		if *d != tt.expected {
			t.Errorf("Expected %v, got %v", tt.expected, *d)
		}
	}
}
