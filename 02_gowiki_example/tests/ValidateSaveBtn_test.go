package main

import (
	"testing"

	"github.com/tebeka/selenium"
)

func TestIsSaveBtnThere(t *testing.T) {
	caps := selenium.Capabilities{"browserName": "chrome"}
	driver, err := selenium.NewRemote(caps, "")
	if err != nil {
		panic(err)
	}
	defer driver.Quit()

	driver.Get("localhost:8080/edit/fooBar")

	saveBtn, err := driver.FindElement(selenium.ByID, "savebtn")
	if err != nil {
		panic(err)
	}
	saveBtn.IsDisplayed()
}
