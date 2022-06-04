package test

import (
	"os"
	"testing"

	"github.com/bruckmann/termNews/pkg/scraper"
)

func TestScraper(t *testing.T) {
  scraper.DownloadRSSNews()
  _, got := os.Stat("./data/news-g1.xml")

  if got != nil {
    t.Errorf("DonwloadRSSNews: error: %s", got)
  }
}
