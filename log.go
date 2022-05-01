package log

import (
	"bufio"
	"encoding/json"
	"io"
)

// UnMarshalJSONFormat decodes a json formatted log file into a Go object.
func UnMarshalJSONFormat(data io.Reader) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		var line map[string]interface{}
		if err := json.Unmarshal(scanner.Bytes(), &line); err != nil {
			return nil, err
		}
		result = append(result, line)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return result, nil
}
