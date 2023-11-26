package model

type DeleteStruct struct {
	Status  int    `json:"status" form:"status"`
	Message string `json:"message" form:"message"`
}

type DataResponse struct {
	Data []Order `json:"data"`
}
