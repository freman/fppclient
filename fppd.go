package fppclient

import "fmt"

func (c Client) GetFPPDStatus() (FPPDStatus, error) {
	var resp FPPDStatus
	if err := c.httpGet("/api/fppd/schedule", &resp); err != nil {
		return FPPDStatus{}, fmt.Errorf("unable to retrieve fppd status: %w", err)
	}

	return resp, nil
}
