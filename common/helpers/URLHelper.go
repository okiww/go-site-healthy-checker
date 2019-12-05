package helpers

import (
	"fmt"
	"net/http"
	"net/url"
	"site-health-check/common/infra/socket"
	"time"
)

var client = http.Client{}
var HEALTH_TIME_NORMAL int64 = 800 // time normal in millisecond

func ValidateURL(URL string) error {
	_, err := url.ParseRequestURI(URL)
	if err != nil {
		return err
	}

	return nil
}

func Checker(domain string) (int, string, error) {
	url := domain
	status := "UNHEALTHY"

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
		status = "HEALTHY"
	}
	return resp.StatusCode, status, nil
}

func CheckerSite(domain string, prefix string, code chan<- int, status chan<- string, error chan<- error) {
	for  {
		siteCode, stat, _ := Checker(domain)

		fmt.Println("============")
		fmt.Println(prefix)
		fmt.Println(siteCode)
		fmt.Println(stat)
		fmt.Println("============")
		//code <- siteCode
		//status <- stat
		//
		//var c int
		//c = <-code
		//s := <-status

		currenConn := socket.GetCurrentConnection()
		socket.BrodacastMessage(currenConn, siteCode, stat, domain, prefix)
		//status <- stat
		//error <- err
		//Checker(domain)
		time.Sleep(5 * time.Minute)
	}

}
