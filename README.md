#This is where the example is:
https://github.com/tebeka/selenium

#This is where the code is:
https://github.com/tebeka/selenium/blob/master/doc.go

#This is how you execute the tests
https://godoc.org/github.com/tebeka/selenium
1. Run your version of the .jar. e.g., java -jar selenium-server-standalone-2.24.1.jar
2. Now, use the command "go test". This will run all tests in the folder. Test in Go are specific. 
You must write tests in a certain way. Notice Example_test.go. This test file is denoted with "_test". 
Test methods starts with "func TestSelenium(t *testing.T)". More reading here, https://golang.org/pkg/testing/. 