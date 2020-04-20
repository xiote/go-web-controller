package govkr

import (
	"github.com/tebeka/selenium"
)

func NewController(webDriver selenium.WebDriver) Controller {
	return Controller{webDriver}
}

type Controller struct {
	selenium.WebDriver
}

func (c *Controller) Navigate(url string) error {
	return c.WebDriver.Get(url)
}
