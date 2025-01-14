package ussd

import (
	"github.com/alainmucyo/ussd-go/sessionstores"
	"github.com/stretchr/testify/suite"
	"testing"
)

type DataBagSuite struct {
	suite.Suite
	store      sessionstores.Store
	databag    *DataBag
	request    *Request
	key, value string
}

type testStructValue struct {
	Name string
	Age  int
}

func (d *DataBagSuite) SetupSuite() {
	d.store = sessionstores.NewRedis("localhost:6379")
	err := d.store.Connect()
	d.Nil(err)
	d.request = &Request{}
	d.request.SessionId = "233246662003"
	d.request.PhoneNumber = "vodafone"
	d.request.Text = "*123#"
	d.key = "name"
	d.value = "Samora"
	d.databag = newDataBag(d.store, d.request)
}

func (d *DataBagSuite) TearDownSuite() {
	err := d.store.Close()
	d.Nil(err)
}

func (d *DataBagSuite) TestDataBag() {
	name := d.request.SessionId + "DataBag"

	err := d.databag.Set(d.key, d.value)
	d.Nil(err)
	val, err := d.store.HashGetValue(name, d.key)
	d.Nil(err)
	d.Equal(d.value, val)

	val, err = d.databag.Get(d.key)
	d.Nil(err)
	d.Equal(d.value, val)

	exists := d.databag.Exists(d.key)
	d.True(exists)

	err = d.databag.Delete(d.key)
	d.Nil(err)
	exists = d.databag.Exists(d.key)
	d.False(exists)

	err = d.databag.Clear()
	d.Nil(err)
	exists = d.store.HashExists(name)
	d.False(exists)

	v := &testStructValue{Name: "Samora", Age: 29}
	err = d.databag.SetMarshaled("user", v)
	d.Nil(err)

	v = new(testStructValue)
	err = d.databag.GetUnmarshaled("user", v)
	d.Nil(err)
	d.Equal("Samora", v.Name)
	d.Equal(29, v.Age)
}

func TestDataBagSuite(t *testing.T) {
	suite.Run(t, new(DataBagSuite))
}
