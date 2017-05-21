// Exercises using http://the-internet.herokuapp.com/
package main

import (
	"testing"
	"time"

	"github.com/tebeka/selenium"
)

func TestShouldBeOnTheCorrectPage(t *testing.T) {
	caps := selenium.Capabilities{"browserName": "chrome"}
	driver, err := selenium.NewRemote(caps, "")
	if err != nil {
		panic(err)
	}
	driver.SetImplicitWaitTimeout(time.Millisecond * 30000)
	driver.SetAsyncScriptTimeout(time.Millisecond * 30000)
	driver.SetPageLoadTimeout(time.Millisecond * 30000)
	defer driver.Quit()

	driver.Get("http://the-internet.herokuapp.com/")
	h1, err := driver.FindElement(selenium.ByTagName, "h1")
	if err != nil {
		panic(err)
	}

	t.Log(h1.Text())

}
