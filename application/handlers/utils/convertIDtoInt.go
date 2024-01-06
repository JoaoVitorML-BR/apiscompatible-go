package handlers

import (
	"login/application/handlers/utils/erruser"
	"net/http"
	"strconv"
)

func ConvertID(w http.ResponseWriter, idParam string) (uint64, error) {
	ID, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		erruser.ErrMessageConvertIDtoInt(w, err)
		return 0, err
	}
	return ID, nil
}
