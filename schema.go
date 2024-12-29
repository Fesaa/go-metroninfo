package metroninfo

import "time"

var xmlHeader = []byte(`<?xml version="1.0" encoding="UTF-8"?>`)

type MetronInfo struct {
	IDS             IDS        `xml:"IDS,omitempty"`
	Publisher       *Publisher `xml:"Publisher,omitempty"`
	Series          Series     `xml:"Series"`
	MangaVolume     string     `xml:"MangaVolume,omitempty"`
	CollectionTitle string     `xml:"CollectionTitle,omitempty"`
	Number          string     `xml:"Number,omitempty"`
	Stories         Stories    `xml:"Stories,omitempty"`
	Summary         string     `xml:"Summary,omitempty"`
	Notes           string     `xml:"Notes,omitempty"`
	Prices          Prices     `xml:"Prices,omitempty"`
	CoverDate       *Date      `xml:"CoverDate,omitempty"`
	StoreDate       *Date      `xml:"StoreDate,omitempty"`
	PageCount       uint       `xml:"PageCount,omitempty"`
	Genres          Genres     `xml:"Genres,omitempty"`
	Tags            Tags       `xml:"Tags,omitempty"`
	Arcs            Arcs       `xml:"Arcs,omitempty"`
	Characters      Characters `xml:"Characters,omitempty"`
	Teams           Teams      `xml:"Teams,omitempty"`
	Universes       Universes  `xml:"Universes,omitempty"`
	Locations       Locations  `xml:"Locations,omitempty"`
	GTIN            *GTIN      `xml:"GTIN,omitempty"`
	AgeRating       AgeRating  `xml:"AgeRating,omitempty"`
	Reprints        Reprints   `xml:"Reprints,omitempty"`
	URLs            URLs       `xml:"URLs,omitempty"`
	Credits         Credits    `xml:"Credits,omitempty"`
	LastModified    *time.Time `xml:"LastModified,omitempty"`

	// Internal
	//XmlnsXsd string `xml:"xmlns:xsd,attr"`
	XmlNsXsi          string `xml:"xmlns:xsi,attr"`
	XmlXsiNoNameSpace string `xml:"xsi:noNamespaceSchemaLocation,attr"`
}

type ID struct {
	Source  string `xml:"source,attr,omitempty"`
	Primary bool   `xml:"primary,attr,omitempty"`
	Value   string `xml:",chardata"`
}

type Resource struct {
	ID    string `xml:"id,attr,omitempty"`
	Value string `xml:",chardata"`
}

type Publisher struct {
	ID      string   `xml:"id,attr,omitempty"`
	Name    string   `xml:"Name"`
	Imprint Resource `xml:"Imprint,omitempty"`
}

type AlternativeName struct {
	ID    string       `xml:"id,attr,omitempty"`
	Lang  LanguageCode `xml:"lang,attr,omitempty"`
	Value string       `xml:",chardata"`
}

type Series struct {
	ID   string       `xml:"id,attr,omitempty"`
	Lang LanguageCode `xml:"lang,attr,omitempty"`

	Name             string           `xml:"Name"`
	SortName         string           `xml:"SortName,omitempty"`
	Volume           uint             `xml:"Volume,omitempty"`
	Format           Format           `xml:"Format,omitempty"`
	StartYear        int              `xml:"StartYear,omitempty"`
	IssueCount       uint             `xml:"IssueCount,omitempty"`
	VolumeCount      uint             `xml:"VolumeCount,omitempty"`
	AlternativeNames AlternativeNames `xml:"AlternativeNames,omitempty"`
}

type Price struct {
	Country CountryCode `xml:"country,attr"`
	Value   float64     `xml:",chardata"`
}

type Genre struct {
	ID    string `xml:"id,attr,omitempty"`
	Value string `xml:",chardata"`
}

type Universe struct {
	ID          string `xml:"id,attr,omitempty"`
	Name        string `xml:"Name"`
	Designation string `xml:"Designation,omitempty"`
}

type Arc struct {
	ID     string `xml:"id,attr,omitempty"`
	Name   string `xml:"Name"`
	Number uint   `xml:"Number,omitempty"`
}

type GTIN struct {
	ISBN string `xml:"ISBN"`
	UPC  string `xml:"UPC"`
}

type URL struct {
	Primary bool   `xml:"primary,attr,omitempty"`
	Value   string `xml:",chardata"`
}

type Credit struct {
	Creator Resource `xml:"Creator"`
	Roles   Roles    `xml:"Roles"`
}

type Role struct {
	ID    string    `xml:"id,attr,omitempty"`
	Value RoleValue `xml:",chardata"`
}

type LanguageCode string

type CountryCode string

func (mi *MetronInfo) SetXMLAttributes() {
	//mi.XmlnsXsd = "http://www.w3.org/2001/XMLSchema"
	mi.XmlNsXsi = "http://www.w3.org/2001/XMLSchema-instance"
	mi.XmlXsiNoNameSpace = "MetronInfo.xsd"
}

func NewMetronInfo() *MetronInfo {
	mi := &MetronInfo{}
	mi.SetXMLAttributes()
	return mi
}
