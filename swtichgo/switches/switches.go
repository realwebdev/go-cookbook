package main

import "fmt"

type GlucoseReading struct {
	UserId          int             `json:"userId" form:"userId" validate:"required" field:"userId"`
	Timestamp       int64           `json:"timestamp" form:"timestamp" validate:"required" field:"timestamp"`
	Reading         float64         `json:"reading" form:"reading" validate:"required" field:"reading"`
	BucketTimestamp int64           `json:"bucketTimestamp" form:"bucketTimestamp"`
	MeasurementType string          `json:"measurementType" form:"measurementType"` // karazHealth or karazScale
	MeasurementName string          `json:"measurementName" form:"measurementName"` // heartRate, bloodPressure, etc.
	DeviceID        string          `json:"deviceID"`
	SensorID        string          `json:"sensorID"`
	TrendID         int             `json:"trendID"`
	ReadingRange    string          `json:"readingRange"`    // inrange, belowrange, aboverange
	ReadingCategory string          `json:"readingCategory"` // low, normal, high
	MeasurmentFlags MeasurementFlag `json:"measurmentFlags"` // flag on current reading
	GlucoseStats    GlucoseStats    `json:"glucoseStats"`
}

type GlucoseStats struct {
	GMI           float64 `json:"gmi"`         // glucose management index
	GRI           float64 `json:"gri"`         // glycemia risk index
	MeanGlucose   float64 `json:"meanGlucose"` // mg/dL average glucose
	A1C           float64 `json:"a1c"`         // percentage
	TIR           float64 `json:"tir"`         // percentage of readings in range
	TotalReadings int     `json:"totalReadings"`

	// ranges on all readings
	TotalInTightRange         float64 `json:"totalInTightRange"`
	TotalInRange              float64 `json:"totalInRange"`
	TotalBelowRange           float64 `json:"totalBelowRange"`
	TotalAboveRange           float64 `json:"totalAboveRange"`
	TotalInRangeNotTightRange float64 `json:"totalInRangeNotTightRange"`
	TotalTarget               float64 `json:"totalTarget"`

	// status on all readings
	PercentageTotalLow      float64 `json:"percentageTotalLow"`
	PercentageTotalVeryLow  float64 `json:"percentageTotalVeryLow"`
	PercentageTotalHigh     float64 `json:"percentageTotalHigh"`
	PercentageTotalVeryHigh float64 `json:"percentageTotalVeryHigh"`

	// target on all readings
	PercentageTotalReadingsInTarget float64 `json:"percentageTotalReadingsInTarget"`

	// Total readings in each category
	TotalLowReadings      float64 `json:"totalLowReadings"`
	TotalVeryLowReadings  float64 `json:"totalVeryLowReadings"`
	TotalHighReadings     float64 `json:"totalHighReadings"`
	TotalVeryHighReadings float64 `json:"totalVeryHighReadings"`
}

type MeasurementFlag struct {
	TightRange           bool `json:"tightRange"`
	InRange              bool `json:"inRange"`
	BelowRange           bool `json:"belowRange"`
	AboveRange           bool `json:"aboveRange"`
	InRangeNotTightRange bool `json:"inRangeNotTightRange"`
	Low                  bool `json:"low"`
	VeryLow              bool `json:"veryLow"`
	High                 bool `json:"high"`
	VeryHigh             bool `json:"veryHigh"`
	Target               bool `json:"target"`
}

func main() {
	newReadings := GlucoseReading{
		MeasurmentFlags: MeasurementFlag{
			High:     true,
			VeryHigh: true,
		},
	}
	lastGlucoseReadingData := GlucoseReading{
		GlucoseStats: GlucoseStats{
			TotalLowReadings:      0,
			TotalVeryLowReadings:  0,
			TotalHighReadings:     0,
			TotalVeryHighReadings: 0,
		},
	}

	switch {
	case newReadings.MeasurmentFlags.Low:
		newReadings.GlucoseStats.TotalLowReadings = lastGlucoseReadingData.GlucoseStats.TotalLowReadings + 1
	case newReadings.MeasurmentFlags.VeryLow:
		newReadings.GlucoseStats.TotalVeryLowReadings = lastGlucoseReadingData.GlucoseStats.TotalVeryLowReadings + 1
	case newReadings.MeasurmentFlags.High:
		newReadings.GlucoseStats.TotalHighReadings = lastGlucoseReadingData.GlucoseStats.TotalHighReadings + 1
	case newReadings.MeasurmentFlags.VeryHigh:
		newReadings.GlucoseStats.TotalVeryHighReadings = lastGlucoseReadingData.GlucoseStats.TotalVeryHighReadings + 1

	default:
		fmt.Println("No reading category found")
	}

	fmt.Printf("newReadings: %+v\n", newReadings)
}
