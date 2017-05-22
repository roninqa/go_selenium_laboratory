package pages

import (
	"time"

	"github.com/tebeka/selenium"
)

var driver selenium.WebDriver

// SetUp is exported
func SetUp() {
	caps := selenium.Capabilities{"browerName": "chrome"}
	driver, err := selenium.NewRemote(caps, "")
	if err != nil {
		panic(err)
	}
	driver.SetImplicitWaitTimeout(time.Millisecond * 30000)
	driver.SetAsyncScriptTimeout(time.Millisecond * 30000)
	driver.SetPageLoadTimeout(time.Millisecond * 30000)
}

// TearDown is exported
func TearDown() {
	defer driver.Quit()
}
