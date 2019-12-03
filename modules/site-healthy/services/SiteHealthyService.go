package services

import (
	"site-health-check/modules/site-healthy/dto"
	helpers "site-health-check/common/helpers"
)

type SiteHealthyService struct {}

var form dto.Form

func SiteHealthyServiceHandler() SiteHealthyInterface  {
	svc := &SiteHealthyService{}

	return svc
}

func (service *SiteHealthyService) GetActiveSites() dto.Form {
	return form
}

func (service *SiteHealthyService) PostSite(site dto.Site) (dto.Form, error)  {
	err := helpers.ValidateURL(site.Name)
	if err != nil {
		return dto.Form{}, err
	}

	_, status, err := helpers.Checker(site.Name)
	if err != nil {
		return dto.Form{}, err
	}

	site.Status = status
	form.Sites = append(form.Sites, site)

	return form, nil
}