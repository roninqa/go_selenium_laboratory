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
	savebtn, _ := driver.FindElement(selenium.ByID, "savebtn")
	if true, err := savebtn.IsDisplayed(); err == nil {
		t.Log("PASSED", true)
	} else {
		t.Errorf("FAILED: Button is not displayed: %s\n", err)
	}

}
