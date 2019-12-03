package dto

type Form struct {
	Sites []Site `form:"sites[]"`
}

type Site struct {
	Name string `form:"name"`
	Status string `form:"status"`
}
