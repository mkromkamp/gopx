package main

import "encoding/xml"

// Gpx represents the root of a GPX file
type Gpx struct {
	XMLName      xml.Name  `xml:"gpx"`
	XMLNs        string    `xml:"xmlns,attr"`
	XMLNsXsi     string    `xml:"xmlns:xsi,attr,omitempty"`
	XMLSchemaLoc string    `xml:"xsi:schemaLocation,attr,omitempty"`
	Version      string    `xml:"version,attr"`
	Creator      string    `xml:"creator,attr"`
	Metadata     *Metadata `xml:"metadata,omitempty"`
	Points       Points    `xml:"wpt,omitempty"`
	Routes       []Route   `xml:"rte,omitempty"`
	Tracks       []Track   `xml:"trk"`
}

// Route is a GPX Route
type Route struct {
	XMLName xml.Name `xml:"rte"`
	Name    string   `xml:"name,omitempty"`
	Cmt     string   `xml:"cmt,omitempty"`
	Desc    string   `xml:"desc,omitempty"`
	Src     string   `xml:"src,omitempty"`
	Links   []Link   `xml:"link"`
	Number  int      `xml:"number,omitempty"`
	Type    string   `xml:"type,omitempty"`
	Points  `xml:"rtept"`
}

// Metadata is a GPX metadata tag
type Metadata struct {
	XMLName   xml.Name   `xml:"metadata"`
	Name      string     `xml:"name,omitempty"`
	Desc      string     `xml:"desc,omitempty"`
	Author    *Person    `xml:"author,omitempty"`
	Copyright *Copyright `xml:"copyright,omitempty"`
	Links     []Link     `xml:"link,omitempty"`
	Timestamp string     `xml:"time,omitempty"`
	Keywords  string     `xml:"keywords,omitempty"`
	Bounds    *Bounds    `xml:"bounds"`
}

// Track is a GPX track
type Track struct {
	XMLName  xml.Name       `xml:"trk"`
	Name     string         `xml:"name,omitempty"`
	Cmt      string         `xml:"cmt,omitempty"`
	Desc     string         `xml:"desc,omitempty"`
	Src      string         `xml:"src,omitempty"`
	Links    []Link         `xml:"link"`
	Number   int            `xml:"number,omitempty"`
	Type     string         `xml:"type,omitempty"`
	Segments []TrackSegment `xml:"trkseg"`
}

// TrackSegment is a GPX track segment
type TrackSegment struct {
	XMLName xml.Name `xml:"trkseg"`
	Points  `xml:"trkpt"`
}

// Points is a collection of points whether in a track, a route, or standalone.
type Points []Point

// Point is a GPX Point
type Point struct {
	Lat float64 `xml:"lat,attr"`
	Lon float64 `xml:"lon,attr"`
	// Position info
	Elevation   float64 `xml:"ele,omitempty"`
	Timestamp   string  `xml:"time,omitempty"`
	MagVar      string  `xml:"magvar,omitempty"`
	GeoIDHeight string  `xml:"geoidheight,omitempty"`
	// Description info
	Name  string `xml:"name,omitempty"`
	Cmt   string `xml:"cmt,omitempty"`
	Desc  string `xml:"desc,omitempty"`
	Src   string `xml:"src,omitempty"`
	Links []Link `xml:"link"`
	Sym   string `xml:"sym,omitempty"`
	Type  string `xml:"type,omitempty"`
	// Accuracy info
	Fix          string  `xml:"fix,omitempty"`
	Sat          int     `xml:"sat,omitempty"`
	Hdop         float64 `xml:"hdop,omitempty"`
	Vdop         float64 `xml:"vdop,omitempty"`
	Pdop         float64 `xml:"pdop,omitempty"`
	AgeOfGpsData float64 `xml:"ageofgpsdata,omitempty"`
	DGpsID       int     `xml:"dgpsid,omitempty"`
}

// Link is a GPX link
type Link struct {
	XMLName xml.Name `xml:"link"`
	URL     string   `xml:"href,attr,omitempty"`
	Text    string   `xml:"text,omitempty"`
	Type    string   `xml:"type,omitempty"`
}

// Copyright is a GPX copyright tag
type Copyright struct {
	XMLName xml.Name `xml:"copyright"`
	Author  string   `xml:"author,attr"`
	Year    string   `xml:"year,omitempty"`
	License string   `xml:"license,omitempty"`
}

// Email is a GPX email tag
type Email struct {
	XMLName xml.Name `xml:"email"`
	ID      string   `xml:"id,attr,omitempty"`
	Domain  string   `xml:"domain,attr,omitempty"`
}

// Person is a GPX person tag
type Person struct {
	XMLName xml.Name `xml:"author"`
	Name    string   `xml:"name,omitempty"`
	Email   *Email   `xml:"email,omitempty"`
	Link    *Link    `xml:"link,omitempty"`
}

// Bounds is a GPX bounds tag
type Bounds struct {
	XMLName xml.Name `xml:"bounds"`
	MinLat  float64  `xml:"minlat,attr"`
	MaxLat  float64  `xml:"maxlat,attr"`
	MinLon  float64  `xml:"minlon,attr"`
	MaxLon  float64  `xml:"maxlon,attr"`
}
