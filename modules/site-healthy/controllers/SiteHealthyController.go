package controllers

import (
	"github.com/gin-gonic/gin"
	"net/url"
	"site-health-check/modules/site-healthy/dto"
	"site-health-check/modules/site-healthy/services"
	"strings"
)

type SiteHealthyController struct {
	siteService services.SiteHealthyInterface
}

func SiteHealthyControllerHandler() SiteHealthyController {
	handler := SiteHealthyController{
		siteService: services.SiteHealthyServiceHandler(),
	}

	return handler
}

// Get Sites
func (ctrl *SiteHealthyController) Index(c *gin.Context) {
	sites := ctrl.siteService.GetActiveSites()
	c.HTML(200, "form.html", gin.H{
		"site": sites,
	})
}

// Post Sites
// Param dto.Site
func (ctrl *SiteHealthyController) Post(c *gin.Context) {
	var site dto.Site
	err := c.Bind(&site)
	s := site.Name

	// Generate prefix URL
	u, err := url.Parse(s)
	if err != nil {
		c.JSON(500, gin.H{
			"status":  "posted",
			"data": nil,
			"message": err,
		})
		return
	}

	site.Prefix = strings.Replace(u.Host, ".", "", -1)

	// post to service
	_, err = ctrl.siteService.PostSite(site)
	if err != nil {
		c.JSON(500, gin.H{
			"status":  "posted",
			"data": nil,
			"message": err,
		})
		return
	}

	// running go routine check URL every 5 minutes
	ctrl.siteService.CheckURLEvery5Minutes(site.Name, site.Prefix)

	c.JSON(200, gin.H{
		"status":  "posted",
		"data": site,
		"message": "success get data",
	})
}