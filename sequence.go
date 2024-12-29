package metroninfo

import "encoding/xml"

type IDS []ID

func (s *IDS) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var inner struct {
		Items []ID `xml:"ID"`
	}

	if err := d.DecodeElement(&inner, &start); err != nil {
		return err
	}

	*s = inner.Items
	return nil
}

func (s IDS) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "IDS"
	if err := e.EncodeToken(start); err != nil {
		return err
	}
	for _, id := range s {
		if err := e.Encode(id); err != nil {
			return err
		}
	}
	return e.EncodeToken(start.End())

}

type AlternativeNames []AlternativeName

func (a *AlternativeNames) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var inner struct {
		Items []AlternativeName `xml:"AlternativeName"`
	}

	if err := d.DecodeElement(&inner, &start); err != nil {
		return err
	}

	*a = inner.Items
	return nil
}

func (a AlternativeNames) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "AlternativeNames"
	if err := e.EncodeToken(start); err != nil {
		return err
	}
	for _, id := range a {
		if err := e.Encode(id); err != nil {
			return err
		}
	}
	return e.EncodeToken(start.End())

}

type Story Resource
type Stories []Story

func (s *Stories) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var inner struct {
		Items []Story `xml:"Story"`
	}

	if err := d.DecodeElement(&inner, &start); err != nil {
		return err
	}

	*s = inner.Items
	return nil
}

func (s Stories) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "Stories"
	if err := e.EncodeToken(start); err != nil {
		return err
	}
	for _, id := range s {
		if err := e.Encode(id); err != nil {
			return err
		}
	}
	return e.EncodeToken(start.End())

}

type Prices []Price

func (p *Prices) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var inner struct {
		Items []Price `xml:"Price"`
	}

	if err := d.DecodeElement(&inner, &start); err != nil {
		return err
	}

	*p = inner.Items
	return nil
}

func (p Prices) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "Prices"
	if err := e.EncodeToken(start); err != nil {
		return err
	}

	for _, price := range p {
		if err := e.Encode(price); err != nil {
			return err
		}
	}

	return e.EncodeToken(start.End())
}

type Genres []Genre

func (g *Genres) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var inner struct {
		Items []Genre `xml:"Genre"`
	}

	if err := d.DecodeElement(&inner, &start); err != nil {
		return err
	}

	*g = inner.Items
	return nil
}

func (g Genres) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "Genres"
	if err := e.EncodeToken(start); err != nil {
		return err
	}

	for _, genre := range g {
		if err := e.Encode(genre); err != nil {
			return err
		}
	}

	return e.EncodeToken(start.End())
}

type Character Resource
type Characters []Character

func (c *Characters) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var inner struct {
		Items []Character `xml:"Character"`
	}

	if err := d.DecodeElement(&inner, &start); err != nil {
		return err
	}

	*c = inner.Items
	return nil
}

func (c Characters) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "Characters"
	if err := e.EncodeToken(start); err != nil {
		return err
	}
	for _, id := range c {
		if err := e.Encode(id); err != nil {
			return err
		}
	}
	return e.EncodeToken(start.End())
}

type Team Resource
type Teams []Team

func (t *Teams) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var inner struct {
		Items []Team `xml:"Team"`
	}

	if err := d.DecodeElement(&inner, &start); err != nil {
		return err
	}

	*t = inner.Items
	return nil
}

func (t Teams) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "Teams"
	if err := e.EncodeToken(start); err != nil {
		return err
	}
	for _, id := range t {
		if err := e.Encode(id); err != nil {
			return err
		}
	}
	return e.EncodeToken(start.End())
}

type Location Resource
type Locations []Location

func (l *Locations) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var inner struct {
		Items []Location `xml:"Location"`
	}

	if err := d.DecodeElement(&inner, &start); err != nil {
		return err
	}

	*l = inner.Items
	return nil
}

func (l Locations) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "Locations"
	if err := e.EncodeToken(start); err != nil {
		return err
	}
	for _, id := range l {
		if err := e.Encode(id); err != nil {
			return err
		}
	}
	return e.EncodeToken(start.End())
}

type Tag Resource
type Tags []Tag

func (t *Tags) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var inner struct {
		Items []Tag `xml:"Tag"`
	}

	if err := d.DecodeElement(&inner, &start); err != nil {
		return err
	}

	*t = inner.Items
	return nil
}

func (t Tags) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "Tags"
	if err := e.EncodeToken(start); err != nil {
		return err
	}
	for _, id := range t {
		if err := e.Encode(id); err != nil {
			return err
		}
	}
	return e.EncodeToken(start.End())
}

type Universes []Universe

func (u *Universes) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var inner struct {
		Items []Universe `xml:"Universe"`
	}

	if err := d.DecodeElement(&inner, &start); err != nil {
		return err
	}

	*u = inner.Items
	return nil
}

func (u Universes) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "Universes"
	if err := e.EncodeToken(start); err != nil {
		return err
	}
	for _, id := range u {
		if err := e.Encode(id); err != nil {
			return err
		}
	}
	return e.EncodeToken(start.End())
}

type Arcs []Arc

func (a *Arcs) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var inner struct {
		Items []Arc `xml:"Arc"`
	}

	if err := d.DecodeElement(&inner, &start); err != nil {
		return err
	}

	*a = inner.Items
	return nil
}

func (a Arcs) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "Arcs"
	if err := e.EncodeToken(start); err != nil {
		return err
	}

	for _, arc := range a {
		if err := e.Encode(arc); err != nil {
			return err
		}
	}

	return e.EncodeToken(start.End())
}

type Reprint Resource
type Reprints []Reprint

func (r *Reprints) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var inner struct {
		Items []Reprint `xml:"Reprint"`
	}

	if err := d.DecodeElement(&inner, &start); err != nil {
		return err
	}

	*r = inner.Items
	return nil
}

func (r Reprints) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "Reprints"
	if err := e.EncodeToken(start); err != nil {
		return err
	}
	for _, id := range r {
		if err := e.Encode(id); err != nil {
			return err
		}
	}
	return e.EncodeToken(start.End())
}

type URLs []URL

func (r *URLs) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var inner struct {
		Items []URL `xml:"URL"`
	}

	if err := d.DecodeElement(&inner, &start); err != nil {
		return err
	}

	*r = inner.Items
	return nil
}

func (r URLs) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "URLs"
	if err := e.EncodeToken(start); err != nil {
		return err
	}
	for _, id := range r {
		if err := e.Encode(id); err != nil {
			return err
		}
	}
	return e.EncodeToken(start.End())
}

type Roles []Role

func (r *Roles) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var inner struct {
		Items []Role `xml:"Role"`
	}

	if err := d.DecodeElement(&inner, &start); err != nil {
		return err
	}

	*r = inner.Items
	return nil
}

func (r Roles) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "Roles"
	if err := e.EncodeToken(start); err != nil {
		return err
	}

	for _, id := range r {
		if err := e.Encode(id); err != nil {
			return err
		}
	}

	return e.EncodeToken(start.End())
}

type Credits []Credit

func (c *Credits) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var inner struct {
		Items []Credit `xml:"Credit"`
	}

	if err := d.DecodeElement(&inner, &start); err != nil {
		return err
	}

	*c = inner.Items
	return nil
}

func (c Credits) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "Credits"
	if err := e.EncodeToken(start); err != nil {
		return err
	}

	for _, id := range c {
		if err := e.Encode(id); err != nil {
			return err
		}
	}

	return e.EncodeToken(start.End())
}
