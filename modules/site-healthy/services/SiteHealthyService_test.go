package services

import (
	"github.com/stretchr/testify/assert"
	"site-health-check/modules/site-healthy/dto"
	"testing"
)

func TestSiteHealthyService_PostSite(t *testing.T) {
	var expected dto.Form
	site := dto.Site{
		Name:   "https://www.github.com",
		Status: "HEALTHY",
		Prefix: "githubcom",
	}
	expected.Sites = append(expected.Sites, site)

	handler := SiteHealthyServiceHandler()
	res, err := handler.PostSite(site)

	assert.Equal(t, expected, res)
	assert.NotNil(t, expected)
	assert.Equal(t, nil, err)
}

func TestSiteHealthyService_GetActiveSites(t *testing.T) {
	handler := SiteHealthyServiceHandler()
	res := handler.GetActiveSites()

	assert.Equal(t, dto.Form{}, res)
}
