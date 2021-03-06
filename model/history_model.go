package model

import "time"

type History struct {
	StatusCode   string          `json:"s"`
	ErrorMessage string          `json:"errmsg,omitempty"`
	BarTime      []time.Duration `json:"t"`
	ClosingPrice []float64       `json:"c"`
	OpeningPrice []float64       `json:"o"`
	HighPrice    []float64       `json:"h"`
	LowPrice     []float64       `json:"l"`
	Volume       []int           `json:"v"`
	NextTime     string          `json:"nextTime,omitempty"`
}

type DexTrade struct {
	TimeInterval  TimeInterval `json:"timeInterval"`
	High          float64      `json:"high"`
	Low           float64      `json:"low"`
	Open          string       `json:"open"`
	Close         string       `json:"close"`
	BaseCurrency  Currency     `json:"baseCurrency"`
	QuoteCurrency Currency     `json:"quoteCurrency"`
	Date          Date         `json:"date"`
	Trades        int          `json:"trades"`
}

type Date struct {
	Date string `json:"date"`
}

type Currency struct {
	Name string `json:"name"`
}

type TimeInterval struct {
	Minute string `json:"minute"`
}

const (
	STATUS_OK      = "ok"
	STATUS_ERROR   = "error"
	STATUS_NO_DATA = "no_data"
)
