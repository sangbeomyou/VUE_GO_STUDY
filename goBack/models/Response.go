package models

type Response struct {
	Success string      `json:"success"`
	Result  interface{} `json:"result"`
}
