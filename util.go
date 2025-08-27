package lybic

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func tryToGetDto[T any](resp *http.Response, dto *T) error {
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		if dto != nil {
			return json.NewDecoder(resp.Body).Decode(dto)
		}
		return nil
	}

	var apiErr Error
	if err := json.NewDecoder(resp.Body).Decode(&apiErr); err != nil {
		return Error{
			Code:    strconv.Itoa(resp.StatusCode),
			Message: "request failed with status " + resp.Status + ", and could not decode error response body: " + err.Error(),
		}
	}
	return apiErr
}
