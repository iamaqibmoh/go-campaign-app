package helper

import (
	"bwa-campaign-app/model/web"
)

func APIResponse(message string, code int, status string, data interface{}) web.Response {
	return web.Response{
		Meta: web.Meta{
			Message: message,
			Code:    code,
			Status:  status,
		},
		Data: data,
	}
}
