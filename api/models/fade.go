package models

/*
	Example
	{
		"start_color": {
			"r": 0,
			"g": -50,
			"b": -120
		},
		"end_color": {
			"r": 255,
			"g": 130,
			"b": 40
		},
		"start_intensity": 0,
		"end_intensity": 1,
		"duration": 123
	}
 */
type Fade struct {
	StartColor Color `json:"start_color"`
	EndColor Color `json:"end_color"`
	StartIntensity float64 `json:"start_intensity"`
	EndIntensity float64 `json:"end_intensity"`
	Duration int `json:"duration"`
}