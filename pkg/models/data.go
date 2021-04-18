package models

type ViewData struct {
	Site ViewSiteData
	Permalink string
	Title string
	Content string
	Layout string
}

type ViewSiteData struct {
	Title string
	BaseURL string
	LanguageCode string
}