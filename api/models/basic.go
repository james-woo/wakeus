package models

/*
	Example
	{
		"color": {
			"r": 0,
			"g": -50,
			"b": -120
		},
		"intensity": 1
	}
*/
type Basic struct {
	Color Color `json:"color"`
	Intensity float32 `json:"intensity"`
}