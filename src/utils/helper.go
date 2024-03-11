package utils

import (
	"net/url"
	"strconv"
)

type LimitOffset struct {
	Limit  int
	Offset int
}

type ResponseDetail struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func ResponseDetailOutput(success bool, code int, message string, data interface{}) ResponseDetail {
	res := ResponseDetail{
		Success: success,
		Code:    code,
		Message: message,
		Data:    data,
	}
	return res
}

func GetLimitOffset(urlValues url.Values) (output LimitOffset, err error) {
	param := new(LimitOffset)

	offset := urlValues.Get("offset")
	if offset != "" {
		offsetInt, err := strconv.Atoi(offset)
		if err != nil {
			return LimitOffset{}, err
		}
		param.Offset = offsetInt
	} else {
		param.Offset = 0
	}
	limit := urlValues.Get("limit")
	if limit != "" {
		limitInt, err := strconv.Atoi(limit)
		if err != nil {
			return LimitOffset{}, err
		}
		param.Limit = limitInt
	} else {
		param.Limit = 20
	}

	return *param, nil
}
