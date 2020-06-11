package weather

import (
	"encoding/xml"
	"fmt"
	"strings"
	"testing"
)

func TestDateTime(t *testing.T) {
	data := `<dateTime name="xmlCreation" zone="HAE" UTCOffset="-4">
<year>2020</year>
<month name="juin">06</month>
<day name="jeudi">11</day>
<hour>13</hour>
<minute>46</minute>
<timeStamp>20200611134600</timeStamp>
<textSummary>11 juin 2020 13h46 HAE</textSummary>
</dateTime>`

	var dt DateTime
	reader := strings.NewReader(data)
	err := xml.NewDecoder(reader).Decode(&dt)
	fmt.Println(err)
	fmt.Println(dt)
}

func TestLocation(t *testing.T) {
	data := `<location>
<continent>Amérique du Nord</continent>
<country code="ca">Canada</country>
<province code="on">Ontario</province>
<name code="s0000368" lat="43.37N" lon="79.81O">Burlington</name>
<region>Burlington - Oakville</region>
</location>`

	var dt Location
	reader := strings.NewReader(data)
	err := xml.NewDecoder(reader).Decode(&dt)
	fmt.Println(err)
	fmt.Println(dt)
}

func TestCurrentConditions(t *testing.T) {
	data := `<currentConditions>
<station code="wwb" lat="43.3N" lon="79.8W">Burlington Lift Bridge</station>
<dateTime name="observation" zone="UTC" UTCOffset="0">
<year>2020</year>
<month name="June">06</month>
<day name="Thursday">11</day>
<hour>16</hour>
<minute>00</minute>
<timeStamp>20200611160000</timeStamp>
<textSummary>Thursday June 11, 2020 at 16:00 UTC</textSummary>
</dateTime>
<dateTime name="observation" zone="EDT" UTCOffset="-4">
<year>2020</year>
<month name="June">06</month>
<day name="Thursday">11</day>
<hour>12</hour>
<minute>00</minute>
<timeStamp>20200611120000</timeStamp>
<textSummary>Thursday June 11, 2020 at 12:00 EDT</textSummary>
</dateTime>
<condition/>
<iconCode format="gif"/>
<temperature unitType="metric" units="C">21.6</temperature>
<dewpoint unitType="metric" units="C">12.5</dewpoint>
<humidex unitType="metric">24</humidex>
<pressure unitType="metric" units="kPa" change="0.11" tendency="rising">101.5</pressure>
<visibility unitType="metric" units="km"/>
<relativeHumidity units="%">56</relativeHumidity>
<wind>
<speed unitType="metric" units="km/h">14</speed>
<gust unitType="metric" units="km/h"/>
<direction>SW</direction>
<bearing units="degrees">228.6</bearing>
</wind>
</currentConditions>`

	var dt CurrentConditions
	reader := strings.NewReader(data)
	err := xml.NewDecoder(reader).Decode(&dt)
	fmt.Println(err)
	fmt.Println(dt)

}