package helpers

import "example/dto"

func GenerateErrorResponse(m string, d interface{}) dto.ErrorResponse {
	return dto.ErrorResponse{m, d}
}
