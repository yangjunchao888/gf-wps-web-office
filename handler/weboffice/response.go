package weboffice

import (
	weboffice2 "gf-wps-web-office/weboffice"
	"net/http"
)

func JsonConstruct(data any, err error) (statusCode int, wpsData any) {
	if err != nil {
		var respErr *weboffice2.Error
		if e, ok := err.(*weboffice2.Error); ok {
			respErr = e
		} else {
			respErr = weboffice2.ErrInternalError.WithMessage(err.Error())
		}

		return respErr.StatusCode(), &weboffice2.Reply{Code: respErr.Code(), Message: respErr.Message()}
	} else {
		return http.StatusOK, &weboffice2.Reply{Code: weboffice2.OK, Data: data}
	}
}
