package govkr

import (
	"testing"

	"github.com/xiote/go-web-controller/govkr/mocks" // mockery --name InterfaceName
)

func TestGet(t *testing.T) {
	in := "http://www.daum.net"
	var want error = nil
	webDriverMock := &mocks.WebDriver{}
	webDriverMock.On("Get", in).Return(nil)

	c := NewController(webDriverMock)
	got := c.Get(in)
	if got != nil {
		t.Errorf("Get(%q) == %v, want %v", in, got, want)
	}
	webDriverMock.AssertExpectations(t)

}
