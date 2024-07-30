package messagecode

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func messageMapToAnyMap(messageMap map[string]messageCode) (map[string]any, error) {
	byteMap := make(map[string]any, len(messageMap))

	for key, val := range messageMap {
		blob, err := json.Marshal(val)
		if err != nil {
			return nil, err
		}

		byteMap[key] = blob
	}

	return byteMap, nil
}

func byteMapToMessageCodeMap(byteMap map[string]string) (map[string]messageCode, error) {
	messsageCodeMap := make(map[string]messageCode, len(byteMap))
	for key, val := range byteMap {
		var messCode messageCode
		err := json.Unmarshal([]byte(val), &messCode)
		if err != nil {
			return nil, err
		}

		messsageCodeMap[key] = messCode
	}

	return messsageCodeMap, nil
}

func makeHashKey(messageGroup int) string {
	return fmt.Sprintf("messagegroup:%d", messageGroup)
}

func makeFieldKey(locale string, messageCode int) string {
	return fmt.Sprintf("%s:%d", locale, messageCode)
}

func fallbackMessageCodeToHTTPCode(code int) int {
	messCodeStr := fmt.Sprintf("%d", code)

	if len(messCodeStr) != 6 {
		return http.StatusInternalServerError
	}

	switch messCodeStr[2] {
	case '2':
		return http.StatusOK
	case '4':
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
