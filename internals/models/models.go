package models

import "net/http"

var Error = map[int]string{
	http.StatusBadRequest:          "Bad Request",
	http.StatusNotFound:            "Page not found",
	http.StatusMethodNotAllowed:    "Method not allowed",
	http.StatusInternalServerError: "Internal Server Error",
	http.StatusUnauthorized:        "Unauthorized",
	http.StatusForbidden:           "Forbidden",
}
