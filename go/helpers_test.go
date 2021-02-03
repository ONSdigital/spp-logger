package spp_logger_test

import (
	"encoding/json"
	"strings"
)

func parseLogLines(logs string) ([]map[string]string, error) {
	var log_lines []map[string]string
	split_logs := strings.Split(logs, "\n")
	for _, log := range split_logs {
		if log == "" {
			continue
		}
		logMessage := make(map[string]string)
		err := json.Unmarshal([]byte(log), &logMessage)
		if err != nil {
			return nil, err
		}
		log_lines = append(log_lines, logMessage)
	}
	return log_lines, nil
}
