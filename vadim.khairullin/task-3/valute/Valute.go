package valute

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

type Valute struct {
	NumCode  int     `xml:"NumCode" json:"num_code"`
	CharCode string  `xml:"CharCode" json:"char_code"`
	Value    float64 `xml:"-" json:"value"`
}

func (v *Valute) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type tempValute struct {
		NumCode  int    `xml:"NumCode"`
		CharCode string `xml:"CharCode"`
		Value    string `xml:"Value"`
	}

	var temp tempValute

	if err := d.DecodeElement(&temp, &start); err != nil {
		return fmt.Errorf("failed to decode Valute element: %w", err)
	}

	v.NumCode = temp.NumCode
	v.CharCode = temp.CharCode

	valueStr := strings.ReplaceAll(temp.Value, ",", ".")

	val, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return fmt.Errorf("failed to parse Valute value %q: %w", temp.Value, err)
	}

	v.Value = val

	return nil
}
