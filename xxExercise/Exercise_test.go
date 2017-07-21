package main

import (
	"fmt"
	"testing"

	"github.com/roninqa/webdriver"
	"github.com/tebeka/selenium"
)

func TestShouldUseThePackage(t *testing.T) {
	driver := webdriver.SetUpChrome()
	defer driver.Quit()

	driver.Get("http://the-internet.herokuapp.com/")

	// Get all of the 'li' and range to get the count
	li, _ := driver.FindElements(selenium.ByTagName, "li")
	for l := range li {
		fmt.Println("li are: ", l)
	}

	t.Log("hello, world")
}
