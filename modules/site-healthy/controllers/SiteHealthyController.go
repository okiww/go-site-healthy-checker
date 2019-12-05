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
		panic(err)
	}

	site.Prefix = strings.Replace(u.Host, ".", "", -1)

	if err != nil {
		c.HTML(200, "form.html", gin.H{
			"site": dto.Site{},
		})
		return
	}

	// post to service
	sites, err := ctrl.siteService.PostSite(site)
	if err != nil {
		sites = ctrl.siteService.GetActiveSites()
		c.HTML(200, "form.html", gin.H{
			"site": sites,
			"error": "URL is not valid",
		})
		return
	}

	// running go routine check URL every 5 minutes
	go ctrl.siteService.CheckURLEvery5Minutes(site.Name, site.Prefix)

	c.HTML(200, "form.html", gin.H{
		"site": sites,
	})
}