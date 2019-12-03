package controllers

import (
	"github.com/gin-gonic/gin"
	"site-health-check/modules/site-healthy/dto"
	"site-health-check/modules/site-healthy/services"
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

func (ctrl *SiteHealthyController) Index(c *gin.Context) {
	sites := ctrl.siteService.GetActiveSites()
	c.HTML(200, "form.html", gin.H{
		"site": sites,
	})
}

func (ctrl *SiteHealthyController) Post(c *gin.Context) {
	var site dto.Site
	err := c.Bind(&site)

	if err != nil {
		c.HTML(200, "form.html", gin.H{
			"site": dto.Site{},
		})
		return
	}

	sites, err := ctrl.siteService.PostSite(site)
	if err != nil {
		sites = ctrl.siteService.GetActiveSites()
		c.HTML(200, "form.html", gin.H{
			"site": sites,
			"error": "URL is not valid",
		})
		return
	}
	c.HTML(200, "form.html", gin.H{
		"site": sites,
	})
}
