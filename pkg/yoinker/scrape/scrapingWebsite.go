package scrape

//ScrapingWebsite declares supported scraping websites
type ScrapingWebsite string

const (
	//CRIMSON scrapes from crimsonmagic
	CRIMSON ScrapingWebsite = "crimsonmagic"
	//WUXIA scrapes from wuxia
	WUXIA ScrapingWebsite = "wuxia"
)