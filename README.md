# weather-can

## Introduction
weather-can is a simple api written in go for accessing weather information published by the [Meteorological Service of Canada](https://eccc-msc.github.io/open-data/readme_en/).

## Usage
weather-can can be installed using go modules:

```bash
go get -u github.com/shauncampbell/weather-can
```

### Retrieving the current conditions in a particular city

Firstly, create a new weather object. There are no arguments required to do this.
```go
w, err := weather.New()
if err != nil {
   // do something.
   return
}
```

The weather data is split into sites/cities. The API allows matching on the site/city by either its english or its french name. For convenience the API contains constants for each of the provinces and territories as well as codes for the english and french languages.

The example below retrieves the site for burlington, ontario matching on its english name (in this case the english and french are the same anyway). The capitalisation of the city doesn't matter either. 
```go
f, err := w.GetSite("Burlington", weather.Ontario, weather.English)
if err != nil {
    // do something.
    return
}
```

Finally, once there is a site then the data for that city/site can be retrieved. The function called `GetSiteData` returns data about the site including current conditions and expected forecast conditions. 
```go
c, err := f.GetSiteData(weather.English)
if err != nil {
    // do something.
    return
}
```

With the site data retrieved the information can be put to some use. The following example prints out the current temperature for Burlington, ON.
```go
fmt.Printf("in %s, %s it is currently %.2f %s\n", f.EnglishName, f.Province, c.CurrentConditions.Temperature.Value, c.CurrentConditions.Temperature.Units)
```
 
