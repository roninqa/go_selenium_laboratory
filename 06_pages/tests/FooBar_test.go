package tests

import (
	"testing"
	"time"

	"github.com/tebeka/selenium"
)

func TestShouldBeOnTheCorrectPageII(t *testing.T) {
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

	// Text func() (string, error); returns a string and error
	// set 'there' to the string and throw away the error
	if there, _ := h1.Text(); there != "Welcome to the Internet" {
		t.Errorf("FAIL: Text is, %s\n", there)
	}

}
