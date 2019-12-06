package helpers

import (
	"net/http"
	"net/url"
	"site-health-check/common/infra/socket"
	"time"
)

var client = http.Client{}
var HEALTH_TIME_NORMAL int64 = 800 // time normal in millisecond
const healthy  = "HEALTHY"
const unhealthy  = "UNHEALTHY"

func ValidateURL(URL string) error {
	_, err := url.ParseRequestURI(URL)
	if err != nil {
		return err
	}

	return nil
}

func Checker(domain string) (int, string, error) {
	url := domain
	status := unhealthy

	time_start := time.Now()
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return 0, "", err
	}
	resp, err := client.Do(req)
	if err != nil {
		return 0, "", err
	}
	resp.Body.Close()
	totalTime := time.Since(time_start).Milliseconds()
	if totalTime < HEALTH_TIME_NORMAL {
		status = healthy
	}
	return resp.StatusCode, status, nil
}

func CheckerSite(domain string, prefix string) {
	for  {
		siteCode, stat, _ := Checker(domain)
		currentConn := socket.GetCurrentConnection()
		socket.BrodacastMessage(currentConn, siteCode, stat, domain, prefix)
		time.Sleep(5 * time.Second)
	}
}
