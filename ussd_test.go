package ussd

import (
	"github.com/alainmucyo/ussd-go/sessionstores"
	"github.com/stretchr/testify/suite"
	"testing"
)

const DummyServiceCode = "*123#"

type UssdSuite struct {
	suite.Suite
	ussd    *Ussd
	request *Request
	store   *sessionstores.Redis
}

func (u *UssdSuite) SetupSuite() {
	u.request = &Request{}
	u.request.SessionId = "233246662003"
	u.request.PhoneNumber = "vodafone"
	u.request.Text = DummyServiceCode

	u.store = sessionstores.NewRedis("localhost:6379")

	u.ussd = New("demo", "Menu")
	u.ussd.Middleware(addData("global", "i'm here"))
	u.ussd.Ctrl(new(demo))
}

/*func (u *UssdSuite) TearDownSuite() {
	u.ussd.end()
}
*/
func (u *UssdSuite) TestUssd() {

	u.Equal(1, len(u.ussd.middlewares))
	u.Equal(2, len(u.ussd.ctrls))

	data := Data{}

	response := u.ussd.process(u.store, data, u.request)
	u.False(response.Release)
	u.Contains(response.Message, "Welcome")

	u.request.Text = "1"
	response = u.ussd.process(u.store, data, u.request)
	u.Contains(response.Message, "Name")

	u.request.Text = "Samora"
	response = u.ussd.process(u.store, data, u.request)
	u.False(response.Release)
	u.Contains(response.Message, "Sex")

	u.request.Text = "1"
	response = u.ussd.process(u.store, data, u.request)
	u.False(response.Release)
	u.Contains(response.Message, "Age")

	u.request.Text = "twenty"
	response = u.ussd.process(u.store, data, u.request)
	u.False(response.Release)
	u.Contains(response.Message, "integer")

	u.request.Text = "29"
	response = u.ussd.process(u.store, data, u.request)
	u.True(response.Release)
	u.Contains(response.Message, "Master Samora")

	u.request.Text = "*123*"
	u.ussd.process(u.store, data, u.request)
	u.request.Text = "0"
	response = u.ussd.process(u.store, data, u.request)
	u.Equal("Bye bye.", response.Message)
}

func TestUssdSuite(t *testing.T) {
	suite.Run(t, new(UssdSuite))
}
