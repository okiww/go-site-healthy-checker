package services

import (
	"errors"
	"runtime"
	helpers "site-health-check/common/helpers"
	"site-health-check/modules/site-healthy/dto"
)

type SiteHealthyService struct{}

var form dto.Form

func SiteHealthyServiceHandler() SiteHealthyInterface {
	svc := &SiteHealthyService{}

	return svc
}

// Service GetActiveSites for get all active site
func (service *SiteHealthyService) GetActiveSites() dto.Form {
	return form
}

// Service PostSite for post site
func (service *SiteHealthyService) PostSite(site dto.Site) (dto.Form, error) {
	err := helpers.ValidateURL(site.Name)
	if err != nil {
		return dto.Form{}, err
	}

	already := service.checkSiteAlreadyActive(site.Name)
	if already {
		err := errors.New("already")
		return dto.Form{}, err
	}

	_, status, err := helpers.Checker(site.Name)
	if err != nil {
		return dto.Form{}, err
	}

	site.Status = status
	site.Prefix = site.Prefix
	form.Sites = append(form.Sites, site)

	return form, nil
}

// Service CheckURLEvery5Minutes for running checker siter every 5 minutes
func (service *SiteHealthyService) CheckURLEvery5Minutes(URL, prefix string) {
	runtime.GOMAXPROCS(2)
	go helpers.CheckerSite(URL, prefix)
}

func (service *SiteHealthyService) checkSiteAlreadyActive(domain string) bool {
	sites := service.GetActiveSites()

	if len(sites.Sites) == 0 {
		return false
	}

	for i, _ := range sites.Sites {
		if sites.Sites[i].Name == domain {
			return true
		}
	}
	return false
}
