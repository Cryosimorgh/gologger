package lg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func LogToService(level, message string) error {
	client := &http.Client{}
	url := fmt.Sprintf("http://%s:%s/%s", ServerIP, Port, strings.ToLower(level))

	reqBody := struct {
		Message string `json:"message"`
	}{
		Message: message,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("X-Service-Name", ServiceName)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

var (
	INFO        = "INFO"
	WARN        = "WARN"
	ERROR       = "ERROR"
	ServiceName = "OverArching Server"
	ServerIP    = ""
	Port        = "8383"
)
