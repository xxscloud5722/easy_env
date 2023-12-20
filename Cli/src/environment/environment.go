package environment

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/pkg/errors"
	"github.com/xxscloud5722/easy_env/cli/src/common"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var client = &http.Client{}

var (
	serverUrl   string
	serverToken string
)

// SetArgs Global Environment Variables
func SetArgs(url string, token string) {
	serverUrl = url
	serverToken = token
}

// Print All Environment.
func Print(prefix string) error {
	color.Green("Get All ...")
	var response []byte
	var err error
	if prefix == "" {
		response, err = get("/pair/list")
	} else {
		response, err = get("/pair/list/" + prefix)
	}
	if err != nil {
		return err
	}
	var result struct {
		Success bool `json:"success"`
		Data    []struct {
			Key         string `json:"key,omitempty"`
			Value       string `json:"value,omitempty"`
			Description string `json:"description,omitempty"`
		} `json:"data"`
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return err
	}
	var table [][]string
	for index, item := range result.Data {
		table = append(table, []string{strconv.Itoa(index + 1), item.Key, valueParse(item.Value), item.Description})
	}
	common.PrintTable([]string{"序号", "Key", "Value", "描述"}, table)
	return nil
}

// PrintByKey Key Environment.
func PrintByKey(key string) error {
	color.Green(fmt.Sprintf("Get Key: %s ...", key))
	response, err := get("/pair/" + key)
	if err != nil {
		return err
	}
	var result struct {
		Success bool   `json:"success"`
		Data    string `json:"data"`
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return err
	}
	if !result.Success || result.Data == "" {
		return errors.New("Key: " + key + " Not Value")
	}
	// 打印输出
	color.Blue(result.Data)
	return nil
}

// Push Key.
func Push(key, value, description string) error {
	if strings.HasPrefix(value, "#file://") {
		file, err := os.ReadFile(value[8:])
		if err != nil {
			return err
		}
		value = string(file)
	}
	data := map[string]interface{}{
		"key":         key,
		"value":       value,
		"description": description,
	}
	response, err := post("/pair/save", data)
	if err != nil {
		return err
	}
	var result struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return err
	}
	if !result.Success {
		return errors.New(result.Message)
	}
	return nil
}

// Remove Key.
func Remove(key string) error {
	data := map[string]interface{}{
		"key": key,
	}
	response, err := post("/pair/remove", data)
	if err != nil {
		return err
	}
	var result struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return err
	}
	if !result.Success {
		return errors.New(result.Message)
	}
	return nil
}

// valueParse Value Parse.
func valueParse(value string) string {
	value = strings.ReplaceAll(value, "\r\n", "\n")
	var rows = strings.Split(value, "\n")
	value = rows[0]
	value = strings.ReplaceAll(value, " ", "　")
	value = strings.ReplaceAll(value, "|", "")
	value = strings.ReplaceAll(value, "-", "")
	if len(value) > 100 {
		return value[0:100] + " ..."
	} else if len(rows) > 1 {
		return value + " ..."
	}
	return value
}

// getServer Server Info
func getServer() (*string, *string, error) {
	return &serverUrl, &serverToken, nil
}

// post
func post(path string, data map[string]interface{}) ([]byte, error) {
	url, accessToken, err := getServer()
	if err != nil {
		return nil, err
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", fmt.Sprintf("%s%s?access-token=%s", *url, path, *accessToken), bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)
	return io.ReadAll(response.Body)
}

// get
func get(path string) ([]byte, error) {
	url, accessToken, err := getServer()
	if err != nil {
		return nil, err
	}
	response, err := http.Get(fmt.Sprintf("%s%s?access-token=%s", *url, path, *accessToken))
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)

	return io.ReadAll(response.Body)
}
