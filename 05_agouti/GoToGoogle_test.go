package main

import (
	"testing"

	"github.com/sclevine/agouti"
)

func TestGoToGoogle(t *testing.T) {
	driver := agouti.Selenium()
	if err := driver.Start(); err != nil {
		t.Fatal("Failed to start Selenium:", err)
	}

	page, err := driver.NewPage(agouti.Browser("chrome"))
	if err != nil {
		t.Fatal("Fail to open page: ", err)
	}

	if err := page.Navigate("http://www.google.com"); err != nil {
		t.Fatal("Failed to navigate: ", err)
	}

	if err := driver.Stop(); err != nil {
		t.Fatal("Failed to stop WebDriver:", err)
	}

}
