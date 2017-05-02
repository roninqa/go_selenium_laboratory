// Exercises using http://the-internet.herokuapp.com/
package main

import (
	"testing"

	"github.com/tebeka/selenium"
)

func TestShouldBeOnTheCorrectPage(t *testing.T) {
	caps := selenium.Capabilities{"browserName": "chrome"}
	driver, err := selenium.NewRemote(caps, "")
	if err != nil {
		panic(err)
	}
	defer driver.Quit()

	driver.Get("http://the-internet.herokuapp.com/")
	h1, _ := driver.FindElement(selenium.ByTagName, "h1")
	if true, err := h1.Text(); err == nil {
		t.Log("PASSED", true)
	} else {
		t.Errorf("FAILED: H1 is not present: %s\n", err)
	}
}
