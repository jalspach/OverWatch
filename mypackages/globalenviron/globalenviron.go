package globalenviron

import "github.com/garethpaul/purpleair-go"

// Series of functions to read from Purple Air. These only pull the first data point. Need to fix that most likely.
// https://github.com/garethpaul/purpleair-go/blob/f06144ce241d44212feb26d73136e9f89232051b/results.go

//Returns the The "name" that appears on the map for this sensor, based on the sensor ID provided
func GetLabel(sensor string) string {
	client := purpleair.NewClient()
	s := client.Sensor(sensor)
	return s.Results[0].Label
}

//Returns the Location Type, based on the sensor ID provided
func GetLocationType(sensor string) string {
	client := purpleair.NewClient()
	s := client.Sensor(sensor)
	return s.Results[0].DEVICELOCATIONTYPE
}

//Returns the Latitude position info, based on the sensor ID provided
func GetLat(sensor string) float64 {
	client := purpleair.NewClient()
	s := client.Sensor(sensor)
	return s.Results[0].Lat
}

//Returns the Longitude position info, based on the sensor ID provided
func GetLon(sensor string) float64 {
	client := purpleair.NewClient()
	s := client.Sensor(sensor)
	return s.Results[0].Lon
}

//Returns the raw PM2.5 (ug/m^3) value, based on the sensor ID provided
func GetPM25Value(sensor string) string {
	client := purpleair.NewClient()
	s := client.Sensor(sensor)
	return s.Results[0].Label
}

//Returns the Last seen data time stamp in UTC, based on the sensor ID provided
func GetLastSeen(sensor string) int {
	client := purpleair.NewClient()
	s := client.Sensor(sensor)
	return s.Results[0].LastSeen
}

//Returns the Sensor type (PMS5003, PMS1003, BME280 etc), based on the sensor ID provided
func GetType(sensor string) string {
	client := purpleair.NewClient()
	s := client.Sensor(sensor)
	return s.Results[0].Type
}

//Returns the Hidden (Hide from public view on map: true/false), based on the sensor ID provided
func GetHidden(sensor string) string {
	client := purpleair.NewClient()
	s := client.Sensor(sensor)
	return s.Results[0].Hidden
}

//Returns the Current version of sensor firmware, based on the sensor ID provided
func GetVersion(sensor string) string {
	client := purpleair.NewClient()
	s := client.Sensor(sensor)
	return s.Results[0].Version
}

//Returns the Last update checked at time stamp in UTC, based on the sensor ID provided
func GetLastUpdateCheck(sensor string) int {
	client := purpleair.NewClient()
	s := client.Sensor(sensor)
	return s.Results[0].LastUpdateCheck
}

//Returns the Created based on the sensor ID provided
func GetCreated(sensor string) int {
	client := purpleair.NewClient()
	s := client.Sensor(sensor)
	return s.Results[0].Created
}

//Returns the Sensor uptime in seconds, based on the sensor ID provided
func GetUptime(sensor string) string {
	client := purpleair.NewClient()
	s := client.Sensor(sensor)
	return s.Results[0].Uptime
}

//Returns the Sensor's WiFi signal strength (RSSI) in dBm, based on the sensor ID provided
func GetRSSI(sensor string) string {
	client := purpleair.NewClient()
	s := client.Sensor(sensor)
	return s.Results[0].RSSI
}

//Returns the ADC based on the sensor ID provided
func GetAdc(sensor string) string {
	client := purpleair.NewClient()
	s := client.Sensor(sensor)
	return s.Results[0].Adc
}

//Returns the Paricles >=.3um (particles/dl) based on the sensor ID provided
func GetP03um(sensor string) string {
	client := purpleair.NewClient()
	s := client.Sensor(sensor)
	return s.Results[0].P03Um
}

//Returns the Paricles >=.5um (particles/dl) based on the sensor ID provided
func GetP05um(sensor string) string {
	client := purpleair.NewClient()
	s := client.Sensor(sensor)
	return s.Results[0].P05Um
}

//Returns the Paricles >=1um (particles/dl) based on the sensor ID provided
func GetP10um(sensor string) string {
	client := purpleair.NewClient()
	s := client.Sensor(sensor)
	return s.Results[0].P10Um
}

//Returns the Paricles >=2.5um (particles/dl) based on the sensor ID provided
func GetP25um(sensor string) string {
	client := purpleair.NewClient()
	s := client.Sensor(sensor)
	return s.Results[0].P25Um
}

//Returns the Paricles >=5.0um (particles/dl) based on the sensor ID provided
func GetP50um(sensor string) string {
	client := purpleair.NewClient()
	s := client.Sensor(sensor)
	return s.Results[0].P50Um
}

//Returns the Paricles >=10.0um (particles/dl) based on the sensor ID provided
func GetP100um(sensor string) string {
	client := purpleair.NewClient()
	s := client.Sensor(sensor)
	return s.Results[0].P100Um
}

//Returns the mass concentration calculated from count data for particle sizes 0.3um to ~1.0um for “standard” particles based on the sensor ID provided
func GetPm10Cf1(sensor string) string {
	client := purpleair.NewClient()
	s := client.Sensor(sensor)
	return s.Results[0].Pm10Cf1
}

//Returns the mass concentration calculated from count data for particle sizes ~0.3um to ~2.5um for “standard” particles based on the sensor ID provided
func GetPm25Cf1(sensor string) string {
	client := purpleair.NewClient()
	s := client.Sensor(sensor)
	return s.Results[0].Pm25Cf1
}

//Returns the mass concentration calculated from count data for particle sizes ~0.3um to ~10um for “standard” particles based on the sensor ID provided
func GetPm100Cf1(sensor string) string {
	client := purpleair.NewClient()
	s := client.Sensor(sensor)
	return s.Results[0].Pm100Cf1
}

//Returns the label based on the sensor ID provided
func GetPm10Atm(sensor string) string {
	client := purpleair.NewClient()
	s := client.Sensor(sensor)
	return s.Results[0].Pm100Atm
}

//Returns the label based on the sensor ID provided
func GetPm100Atm(sensor string) string {
	client := purpleair.NewClient()
	s := client.Sensor(sensor)
	return s.Results[0].Pm100Atm
}

//Returns the raw Humidity (%), based on the sensor ID provided
func GetHumidity(sensor string) string {
	client := purpleair.NewClient()
	s := client.Sensor(sensor)
	return s.Results[0].Humidity
}

//Returns the raw Temp (deg F), based on the sensor ID provided
func GetTempF(sensor string) string {
	client := purpleair.NewClient()
	s := client.Sensor(sensor)
	return s.Results[0].TempF
}

//Returns the Current pressure in Millibars,  based on the sensor ID provided
func GetPressure(sensor string) string {
	client := purpleair.NewClient()
	s := client.Sensor(sensor)
	return s.Results[0].Pressure
}

//Returns the Sensor data age (when data was last received) in minutes, based on the sensor ID provided
func GetAge(sensor string) int {
	client := purpleair.NewClient()
	s := client.Sensor(sensor)
	return s.Results[0].AGE
}

//Returns the Statistics for PM2.5 (https://www2.purpleair.com/community/faq#hc-json-object-fields), based on the sensor ID provided
func GetStats(sensor string) string {
	client := purpleair.NewClient()
	s := client.Sensor(sensor)
	return s.Results[0].Stats
}

