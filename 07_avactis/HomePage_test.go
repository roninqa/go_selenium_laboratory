package main

import (
	"testing"
	"time"

	"github.com/tebeka/selenium"
)

func TestHomePage(t *testing.T) {
	caps := selenium.Capabilities{"browserName": "chrome"}
	driver, err := selenium.NewRemote(caps, "")
	if err != nil {
		t.Fatalf("Oh, snap! Driver error: %s\n", err)
	}
	driver.SetImplicitWaitTimeout(time.Millisecond * 30000)
	driver.SetAsyncScriptTimeout(time.Millisecond * 30000)
	driver.SetPageLoadTimeout(time.Millisecond * 30000)

	defer driver.Quit()

	driver.Get("http://the-internet.herokuapp.com/")
	// now just navigate the pages...
}
