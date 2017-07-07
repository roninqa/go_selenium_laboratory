package main

import (
	"testing"
	"time"

	"fmt"

	"github.com/tebeka/selenium"
)

func TestShouldBeOnTheCorrectPageI(t *testing.T) {
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

	// Get all of the 'li' and range to get the count
	li, _ := driver.FindElements(selenium.ByTagName, "li")
	for l := range li {
		fmt.Println("li are: ", l)
	}

}
