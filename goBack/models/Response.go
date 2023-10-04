package models

type Response struct {
	Success string      `json:"success"`
	token   interface{} `json:"token"`
	Result  interface{} `json:"result"`
}
