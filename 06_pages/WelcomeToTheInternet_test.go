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
	h1, err := driver.FindElement(selenium.ByTagName, "h1")
	if err != nil {
		t.Errorf("ERROR: Cannot find h1: %v\n", err)
	} else if context, _ := h1.Text(); context == "Welcome to the Internet" {
		t.Logf("PASSED: The context of h1: %v\n", context)
	} else {
		t.Errorf("ERROR: The context does not match.")
	}
}
