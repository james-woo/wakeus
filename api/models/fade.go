package models

type Fade struct {
	StartColor Color `json:"start_color"`
	EndColor Color `json:"end_color"`
	StartIntensity int32 `json:"start_intensity"`
	EndIntensity int32 `json:"end_intensity"`
	Duration int32 `json:"duration"`
}