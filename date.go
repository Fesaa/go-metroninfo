package metroninfo

import (
	"encoding/xml"
	"errors"
	"time"
)

var parseLayout = "2006-01-02"

type Date time.Time

func (t Date) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	formattedTime := time.Time(t).Format(parseLayout)
	return e.EncodeElement(formattedTime, start)
}

func (t *Date) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var content string
	if err := d.DecodeElement(&content, &start); err != nil {
		return err
	}

	parsedTime, err := time.Parse(parseLayout, content)
	if err != nil {
		return errors.New("invalid date format, expected YYYY-MM-DD")
	}

	*t = Date(parsedTime)
	return nil
}
