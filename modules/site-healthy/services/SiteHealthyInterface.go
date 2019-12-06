package services

import "site-health-check/modules/site-healthy/dto"

type SiteHealthyInterface interface {
	GetActiveSites() dto.Form
	PostSite(site dto.Site) (dto.Form, error)
	CheckURLEvery5Minutes(URL string, prefix string)
}
