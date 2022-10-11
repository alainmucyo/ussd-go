package ussd

type NsanoRequest struct {
	MSISDN  string `json:"msisdn"`
	Network string `json:"network"`
	Message string `json:"msg"`
}

func (n *NsanoRequest) GetRequest() *Request {
	return &Request{
		SessionId:   n.MSISDN,
		Text:        n.Message,
		PhoneNumber: n.Network,
	}
}

type NsanoResponse struct {
	USSDResp ussdResp
}

type ussdResp struct {
	Action string `json:"action"`
	Menus  string `json:"menus"`
	Title  string `json:"title"`
}

func (n *NsanoResponse) SetResponse(response Response) {
	if response.Release {
		n.USSDResp.Action = "prompt"
		n.USSDResp.Menus = response.Message
	} else {
		n.USSDResp.Action = "input"
		n.USSDResp.Title = response.Message
	}
}
