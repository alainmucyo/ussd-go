package ussd

type SmsghRequest struct {
	Mobile      string
	SessionId   string
	ServiceCode string
	Type        string
	Message     string
	Operator    string
}

func (s *SmsghRequest) GetRequest() *Request {
	return &Request{
		Mobile:      s.Mobile,
		ServiceCode: s.ServiceCode,
		Type:        s.Type,
		Message:     s.Message,
		Operator:    s.Operator,
	}
}

type SmsghResponse struct {
	*Response
}

func (s *SmsghResponse) SetResponse(response *Response) {
	s.Response = response
}