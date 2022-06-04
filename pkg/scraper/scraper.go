package scraper

import (
	"fmt"

  "github.com/gocolly/colly"
)

var (
  baseUrl string = "https://g1.globo.com/rss/g1/brasil/"
  baseDir string = "./data"
  finalFileName string = fmt.Sprintf("%s/news-g1.xml", baseDir)
)

func DownloadRSSNews() {
  c := colly.NewCollector()
  
  if c != nil {
    fmt.Printf("Collector created with success")
  }

  c.OnResponse(func(r *colly.Response) {
    error := r.Save(finalFileName)  
    if error != nil {
      fmt.Printf("An error occurred while downloading news: %s", error)
    } 
  })

  c.Visit(baseUrl)
}
