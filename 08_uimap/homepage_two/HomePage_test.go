package homepage_two

import (
	"testing"

	"time"

	"log"

	"github.com/roninqa/uimap"
	"github.com/roninqa/webdriver"
	"github.com/tebeka/selenium"
)

func TestShouldBeAbleToClickOnABTesting(t *testing.T) {

	// Start time to run test
	start := time.Now()

	driver := webdriver.SetUpChrome()
	defer driver.Quit()

	// Navigate to the site
	driver.Get("http://the-internet.herokuapp.com/")

	// Click on the AB Testing link
	abtesting, err := driver.FindElement(selenium.ByLinkText, uimap.LnkABTesting)
	if err != nil {
		t.Log("***ERROR with abtesting: ", err)
	}
	abtesting.Click()

	// Validate you're in the AB Testing page
	h3, err := driver.FindElement(selenium.ByTagName, uimap.PgABTest)
	if err != nil {
		t.Log("***ERROR with h3: ", err)
	}

	if there, _ := h3.Text(); there != "A/B Test Variation 1" {
		t.Logf("***FAIL: 'There' is: %s\n", there)
	}

	// End time to run test
	elapsed := time.Since(start)
	log.Printf("***Elapsed time: %v\n", elapsed)
}

func TestShouldBeAbleToClickOnCheckBoxes(t *testing.T) {

	// Start time to run test
	start := time.Now()

	driver := webdriver.SetUpChrome()
	defer driver.Quit()

	// Navigate to the site
	driver.Get("http://the-internet.herokuapp.com/")

	// Click on the Checkboxes link
	checkBoxes, err := driver.FindElement(selenium.ByXPATH, uimap.LnkCheckBoxes)
	if err != nil {
		t.Log("***ERROR with checkBoxes: ", err)
	}
	checkBoxes.Click()

	// Validate you're in the Checkboxes page
	h3, err := driver.FindElement(selenium.ByTagName, uimap.PgCheckBoxes)
	if err != nil {
		t.Log("***ERROR with h3: ", err)
	}

	if there, _ := h3.Text(); there != "Checkboxes" {
		t.Logf("***FAIL: 'There' is: %s\n", there)
	}

	// End time to run test
	elapsed := time.Since(start)
	log.Printf("***Elapsed time: %v\n", elapsed)
}

func TestShouldBeAbleToClickOnABTestingOne(t *testing.T) {

	// Start time to run test
	start := time.Now()

	driver := webdriver.SetUpChrome()
	defer driver.Quit()

	// Navigate to the site
	driver.Get("http://the-internet.herokuapp.com/")

	// Click on the AB Testing link
	abtesting, err := driver.FindElement(selenium.ByLinkText, uimap.LnkABTesting)
	if err != nil {
		t.Log("***ERROR with abtesting: ", err)
	}
	abtesting.Click()

	// Validate you're in the AB Testing page
	h3, err := driver.FindElement(selenium.ByTagName, uimap.PgABTest)
	if err != nil {
		t.Log("***ERROR with h3: ", err)
	}

	if there, _ := h3.Text(); there != "A/B Test Variation 1" {
		t.Logf("***FAIL: 'There' is: %s\n", there)
	}

	// End time to run test
	elapsed := time.Since(start)
	log.Printf("***Elapsed time: %v\n", elapsed)
}

func TestShouldBeAbleToClickOnCheckBoxesTwo(t *testing.T) {

	// Start time to run test
	start := time.Now()

	driver := webdriver.SetUpChrome()
	defer driver.Quit()

	// Navigate to the site
	driver.Get("http://the-internet.herokuapp.com/")

	// Click on the Checkboxes link
	checkBoxes, err := driver.FindElement(selenium.ByXPATH, uimap.LnkCheckBoxes)
	if err != nil {
		t.Log("***ERROR with checkBoxes: ", err)
	}
	checkBoxes.Click()

	// Validate you're in the Checkboxes page
	h3, err := driver.FindElement(selenium.ByTagName, uimap.PgCheckBoxes)
	if err != nil {
		t.Log("***ERROR with h3: ", err)
	}

	if there, _ := h3.Text(); there != "Checkboxes" {
		t.Logf("***FAIL: 'There' is: %s\n", there)
	}

	// End time to run test
	elapsed := time.Since(start)
	log.Printf("***Elapsed time: %v\n", elapsed)
}
