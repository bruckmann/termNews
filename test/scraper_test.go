package test

import (
	"testing"

	s "github.com/bruckmann/termNews/scraper"
)

func TestScraper(t *testing.T){
  got := s.Extract() 
  expected := "© Copyright 2000-2022 Globo Comunicação e Participações S.A."
  if got != expected {
    t.Errorf("expected: %s but got %s", expected, got)
  }
}

