package utils

import "koriebruh/management/dto"

func SuccessRes(code int, status string, data interface{}) dto.WebResponse {
	return dto.WebResponse{
		Code:   code,
		Status: status,
		Data:   data,
	}
}
