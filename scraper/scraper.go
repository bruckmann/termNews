package scraper

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

type Response struct {
  text string
}

func Extract() string {
  
  r := Response{}
  
  c := colly.NewCollector()
  
  c.OnRequest(func(r *colly.Request) {
    fmt.Println("Visiting: ", r.URL)
  })

  c.OnError(func(_ *colly.Response, err error){
    log.Println("Something went wrong", err)
  })

  c.OnHTML("body > div.glb-grid > footer > div > div > div > div.footer__service > span", func(e *colly.HTMLElement){
    fmt.Println("First column of a table row:", e.Text)
    r.text = e.Text
  })
  
  c.Visit("https://g1.globo.com/")
  return r.text
}


