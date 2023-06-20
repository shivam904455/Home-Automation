package model

import "time"

type CommonParameres struct {
	Name      string `json:"name"`
	Starttime time.Time `json:"start_time"`
	Endtime   time.Time		`json:"end_time"`
}

// ErrorResponse struct
type ErrorResponse struct{
	Massage string `json:"error"` 
}

// SuccessResponse struct
type SuccessResponse struct{
	massage string `json:"massage"`
}