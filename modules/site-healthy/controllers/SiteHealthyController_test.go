package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"site-health-check/modules/site-healthy/dto"

	"strings"

	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	mockService "site-health-check/gen/mocks"
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
		data := map[string]string{"Name": "https://www.github.com"}
		request, _ := json.Marshal(data)
		// Create a response recorder
		w := httptest.NewRecorder()

		// Get a new router
		r := getRouter(true)

		handler := SiteHealthyControllerHandler()
		// Define the route similar to its definition in the routes file
		r.POST("/post", handler.Post)

		// Create a request to send to the above route
		req, err := http.NewRequest("POST", "/post", bytes.NewBuffer(request))
		if err != nil {
			fmt.Println(err)
		}
		req.Header.Set("Content-Type", "application/json")
		// Create the service and process the above request.
		r.ServeHTTP(w, req)

		// Test that the http status code is 200
		if w.Code != http.StatusOK {
			t.Fail()
		}

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		mock := mockService.NewMockSiteHealthyInterface(mockCtrl)
		//
		var expected dto.Form
		site := dto.Site{
			Name:   "https://www.github.com",
			Status: "HEALTHY",
			Prefix: "githubcom",
		}
		expected.Sites = append(expected.Sites, site)
		errMock := errors.New("Mock Error")
		mock.EXPECT().PostSite(site).Return(expected, errMock).AnyTimes()

		ctrlSite := SiteHealthyController{siteService:mock}
		resService, _ := ctrlSite.siteService.PostSite(site)
		assert.Equal(t, expected, resService)
		assert.NotNil(t, expected)

	})
}