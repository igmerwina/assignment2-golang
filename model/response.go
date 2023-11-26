package model

type DeleteStruct struct {
	Status  int    `json:"status" form:"status"`
	Message string `json:"message" form:"message"`
	Success bool   `json:"success" form:"success"`
}

type DataResponse struct {
	Data []Order `json:"data"`
}
