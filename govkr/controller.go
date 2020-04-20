package govkr

type WebDriver interface {
	Get(url string) error
}

type Controller struct {
	WebDriver
}

func NewController(webDriver WebDriver) Controller {
	return Controller{webDriver}
}

func (c *Controller) Get(url string) error {
	return c.WebDriver.Get(url)
}
