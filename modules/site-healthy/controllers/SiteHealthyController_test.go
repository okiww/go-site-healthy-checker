package controllers

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/url"
	"strings"

	//"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	//mockRepositories "site-health-check/gen/mocks"
	//"site-health-check/modules/site-healthy/dto"
	"testing"
)

// Helper function to create a router during testing
func getRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()
	if withTemplates {
		r.LoadHTMLGlob("../../../views/*")
	}
	return r
}

// Helper function to process a request and test its response
func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {

	// Create a response recorder
	w := httptest.NewRecorder()

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	if !f(w) {
		t.Fail()
	}
}

func TestSiteHealthyController_Index(t *testing.T) {
	t.Run("test get active sites", func(t *testing.T) {
		// Create a response recorder
		w := httptest.NewRecorder()

		// Get a new router
		r := getRouter(true)

		// Set the token cookie to simulate an authenticated user
		http.SetCookie(w, &http.Cookie{Name: "token", Value: "123"})
		handler := SiteHealthyControllerHandler()
		// Define the route similar to its definition in the routes file
		r.GET("/", handler.Index)

		// Create a request to send to the above route
		req, _ := http.NewRequest("GET", "/", nil)

		// Create the service and process the above request.
		r.ServeHTTP(w, req)

		// Test that the http status code is 200
		if w.Code != http.StatusOK {
			t.Fail()
		}

		// Test that the page title is "Home Page"
		// You can carry out a lot more detailed tests using libraries that can
		// parse and process HTML pages
		p, err := ioutil.ReadAll(w.Body)
		if err != nil || strings.Index(string(p), "<title>Okky - Site Healthy Checker</title>") < 0 {
			t.Fail()
		}
	})
}

func TestSiteHealthyController_Post(t *testing.T) {
	t.Run("test post sites", func(t *testing.T) {
		data := url.Values{}
		data.Set("Name", "https://www.github.com")
		// Create a response recorder
		w := httptest.NewRecorder()

		// Get a new router
		r := getRouter(true)

		// Set the token cookie to simulate an authenticated user
		http.SetCookie(w, &http.Cookie{Name: "token", Value: "123"})
		handler := SiteHealthyControllerHandler()
		// Define the route similar to its definition in the routes file
		r.POST("/post", handler.Post)

		// Create a request to send to the above route
		req, err := http.NewRequest("POST", "/post", bytes.NewBufferString(data.Encode()))
		if err != nil {
			fmt.Println(err)
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		// Create the service and process the above request.
		r.ServeHTTP(w, req)

		// Test that the http status code is 200
		if w.Code != http.StatusOK {
			t.Fail()
		}

		// Test that the page title is "Home Page"
		// You can carry out a lot more detailed tests using libraries that can
		// parse and process HTML pages
		//p, err := ioutil.ReadAll(w.Body)
		//if err != nil || strings.Index(string(p), "<title>Okky - Site Healthy Checker</title>") < 0 {
		//	t.Fail()
		//}
	})
}