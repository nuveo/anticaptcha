package anticaptcha

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

var (
	baseURL = &url.URL{Host: "api.anti-captcha.com", Scheme: "https", Path: "/"}
)

type Client struct {
	ApiKey       string
	ProxyType    string
	ProxyAddress string
	ProxyPort    int
}

// Function to create the task to process the recaptcha, returns the task_id
func (c *Client) CreateTask(websiteURL string, recaptchaKey string) float64 {
	// Mount the data to be sent
	body := map[string]interface{}{
		"clientKey": c.ApiKey,
		"task": map[string]interface{}{
			"type":          "NoCaptchaTask",
			"websiteURL":    websiteURL,
			"websiteKey":    recaptchaKey,
			"proxyType":     c.ProxyType,
			"proxyAddress":  c.ProxyAddress,
			"proxyPort":     c.ProxyPort,
			"proxyLogin":    "",
			"proxyPassword": "",
			"userAgent":     "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/52.0.2743.116 Safari/537.36",
		},
	}

	b, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}

	// Make the request
	u := baseURL.ResolveReference(&url.URL{Path: "/createTask"})
	resp, err := http.Post(u.String(), "application/json", bytes.NewBuffer(b))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Decode response
	responseBody := make(map[string]interface{})
	json.NewDecoder(resp.Body).Decode(&responseBody)
	// TODO treat api errors and handle them properly
	return responseBody["taskId"].(float64)
}
