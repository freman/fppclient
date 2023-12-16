package fppclient

import (
	"context"
	"fmt"
)

func (c Client) GetFPPDStatus(ctx context.Context) (FPPDStatus, error) {
	var resp FPPDStatus
	if err := c.httpGet(ctx, "/api/fppd/status", &resp); err != nil {
		return FPPDStatus{}, fmt.Errorf("unable to retrieve fppd status: %w", err)
	}

	return resp, nil
}

func (c Client) GetFPPDSchedule(ctx context.Context) (Schedule, error) {
	var resp ScheduleResponse
	if err := c.httpGet(ctx, "/api/fppd/schedule", &resp); err != nil {
		return Schedule{}, fmt.Errorf("unable to retrieve schedule: %w", err)
	}

	return resp.Schedule, nil
}
