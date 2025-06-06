package lpa

import (
	"net/url"

	sgp22 "github.com/damonto/euicc-go/v2"
)

// Discovery discovers the downloadable profiles from SM-DS.
//
// See https://aka.pw/sgp22/v2.5#page=212 (Section 5.8.2, ES11.AuthenticateClient)
func (c *Client) Discovery(address *url.URL, IMEI []byte) ([]*sgp22.EventEntry, error) {
	response, err := c.InitiateAuthentication(address)
	if err != nil {
		return nil, err
	}
	cardRequest := response.CardRequest()
	cardRequest.IMEI = IMEI
	request, err := sgp22.InvokeAPDU(c.APDU, cardRequest)
	if err != nil {
		return nil, err
	}
	clientResponse, err := sgp22.InvokeHTTP(c.HTTP, address, &sgp22.ES11AuthenticateClientRequest{
		ES9AuthenticateClientRequest: request,
	})
	if err != nil {
		return nil, err
	}
	return clientResponse.EventEntries, nil
}
