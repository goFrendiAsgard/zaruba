package monitoring

import "net/http"

func getHTTPCodeByStatus(status bool) (untyped int) {
	if status {
		return http.StatusOK
	}
	return http.StatusInternalServerError
}
