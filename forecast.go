package weather

import "time"

// DateTime represents a point in time specified in the response
type DateTime struct {
	Name      string  `xml:"name,attr"`      // Name is the title of what the DateTime object represents.
	Zone      string  `xml:"zone,attr"`      // Zone is the name of the timezone that the DateTime object is relative to.
	UTCOffset float64 `xml:"UTCOffset,attr"` // UTCOffset is the number of hours that the specified timezone differs from UTC.
	Year      int     `xml:"year"`           // Year is the year in which the event occurred.
	Month     struct {
		Number int    `xml:",chardata"` // Number is the month represented as an integer, e.g. february = 2
		Name   string `xml:"name,attr"` // Name is the month represented as a text string in english or french.
	} `xml:"month"` // Month is month in which the event occurred.
	Day struct {
		Number int    `xml:",chardata"` // Number is the day of the month represented as an integer, e.g. 25
		Name   string `xml:"name,attr"` // Name is the day of the month represented as a text string in english or french, e.g. Friday
	} `xml:"day"` // Day is the day of the month on which the event occurred.
	Hour      int    `xml:"hour"`        // Hour is the hour of the day when the event occurred.
	Minute    int    `xml:"minute"`      // Minute is the minute of the day when the event occurred.
	Timestamp string `xml:"timeStamp"`   // TimeStamp is a timestamp in the form of yyyymmddhhiiss where all times are UTC.
	Summary   string `xml:"textSummary"` // Summary is a text summary of the DateTime e.g. Thursday June 11, 2020 at 19:30 UTC
}

// ToTime creates a new time object from the current object.
func (d DateTime) ToTime() time.Time {
	return time.Date(d.Year, time.Month(d.Month.Number), d.Day.Number, d.Hour, d.Minute, 0, 0, time.FixedZone(d.Zone,int(d.UTCOffset)))
}

// Location represents a specific place.
type Location struct {
	Continent string            `xml:"continent"` // Continent is the name of the continent that the site is located.
	Country   NameCodedItem     `xml:"country"`   // Country is the name of the country that the site is located.
	Province  NameCodedItem     `xml:"province"`  // Province is the name of the province/state that the site is located.
	Site      NameCodedLocation `xml:"name"`      // Site provides information about the site such as site code, name and lat long coordinates.
	Region    string            `xml:"region"`    // Region provides a textual description of the area in which the site is located.
}

// CurrentConditions represent the most recently recorded weather conditions.
type CurrentConditions struct {
	Station          NameCodedLocation      `xml:"station"`          // Station provides information about where the current conditions were recorded.
	Dates            []DateTime             `xml:"dateTime"`         // Dates provides an array of times about when the current conditions were recorded.
	IconCode         IconCode               `xml:"iconCode"`         // IconCode
	Temperature      MetricWithUnitsAndType `xml:"temperature"`      // Temperature provides information on the current temperature.
	DewPoint         MetricWithUnitsAndType `xml:"dewpoint"`         // DewPoint provides information on the current dew point.
	Humidex          MetricWithUnitsAndType `xml:"humidex"`          // Humidex provides information on the current humidex.
	AirPressure      AirPressure            `xml:"pressure"`         // AirPressure provides information on air pressure.
	Visibility       MetricWithUnitsAndType `xml:"visibility"`       // Visibility provides information on current visibility.
	RelativeHumidity MetricWithUnits        `xml:"relativeHumidity"` // RelativeHumidity provides information on the current humidity.
	Wind             Wind                   `xml:"wind"`             // Wind provides current wind speeds and gusts.
}

// NamedCodedLocation represents a location with a name, code and coordinates.
type NameCodedLocation struct {
	Code      string `xml:"code,attr"` // Code is a code to describe the location.
	Latitude  string `xml:"lat,attr"`  // Latitude is the latitude coordinates of the location.
	Longitude string `xml:"lon,attr"`  // Longitude is the longitude coordinates of the location.
	Name      string `xml:",chardata"` // Name is a name given to the location.
}

// IconCode
type IconCode struct {
	Format string `xml:"format,attr"` // Format describes the format of the IconCode.
}

//	Wind provides information about wind speeds.
type Wind struct {
	Speed     MetricWithUnitsAndType `xml:"speed"`     // Speed describes the current sustained wind speed.
	Gust      MetricWithUnitsAndType `xml:"gust"`      // Gust describes the speed of periodic wind gusts.
	Direction string                 `xml:"direction"` //	Direction describes the direction that the wind is coming from.
	Bearing   MetricWithUnits        `xml:"bearing"`   // Bearing provides a compass bearing for the wind direction.
}

// AirPressure provides information about air pressure.
type AirPressure struct {
	UnitType string  `xml:"unitType,attr"` //	UnitType the type of units associated with the value (e.g. Metric or imperial).
	Units    string  `xml:"units,attr"`    // Units is the actual unit of measurement of the value (e.g. kPa)
	Change   float64 `xml:"change,attr"`   //	Change is a measure of how much the value has changed since the last reading.
	Tendency string  `xml:"tendency,attr"` // Tendency is the trend of the measurement (is it going up or going down).
	Value    float64 `xml:",chardata"`     // Value is the actual value of the measurement.
}

//	NameCodedItem describes an item which is identified by both a human friendly name and also a code.
type NameCodedItem struct {
	Code string `xml:"code,attr"` // Code is the code which represents this item.
	Name string `xml:",chardata"` // Name is the friendly name which represents this item.
}

// MetricWithUnitsAndType describes an item which has a value, but also contains its units and unit type.
type MetricWithUnitsAndType struct {
	UnitType string  `xml:"unitType,attr"` // UnitType is the type of the unit being used (e.g. Metric or imperial)
	Units    string  `xml:"units,attr"`    // Units is the actual units being used (km/h)
	Value    float64 `xml:",chardata"`     // Value is the value of the reading.
}

// MetricWithUnitsAndType describes an item which has a value, but also contains its units and assumed to be metric type.
type MetricWithUnits struct {
	Units string  `xml:"units,attr"` // Units is the actual units being used (km/h)
	Value float64 `xml:",chardata"`  // Value is the value of the reading.
}

type RiseSet struct {
	Disclaimer string     `xml:"disclaimer"` // Legal disclaimer included with the sunrise/sunset data.
	DateTime   []DateTime `xml:"dateTime"`   // List of date and times associated with the rising and setting of the sun.
}

// SiteData represents all the information about a site.
// TODO: Pull in forecasts.
type SiteData struct {
	License           string            `xml:"license"`           //	License contains licensing information
	Dates             []DateTime        `xml:"dateTime"`          // Dates describes some important dates and times associated with the data.
	Location          Location          `xml:"location"`          // Location describes the location that this data relates to.
	CurrentConditions CurrentConditions `xml:"currentConditions"` // CurrentConditions describes the current weather conditions at the location.
	RiseSet           RiseSet           `xml:"riseSet"`           // RiseSet describes the sunrise and sunset times for the location
}
