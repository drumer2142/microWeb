package models

type Response struct {
	Data  Data  `json:"data"`
}

type Data struct {
	SuccessMsg  string `json:"successMsg"`
}