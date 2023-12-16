package fppclient

import (
	"context"
	"errors"
	"fmt"
)

func (c Client) GetSchedule(ctx context.Context) (schedule []ScheduleEntries, err error) {

	if err = c.httpGet(ctx, "/api/schedule", &schedule); err != nil {
		return nil, fmt.Errorf("unable to retrieve schedule: %w", err)
	}

	return schedule, err
}

func (c Client) PostSchedule(ctx context.Context, scheduleIn []ScheduleEntries) (schedule []ScheduleEntries, err error) {
	if err = c.httpPost(ctx, "/api/schedule", &scheduleIn, &schedule); err != nil {
		return nil, fmt.Errorf("unable to update schedule: %w", err)
	}

	return schedule, err
}

func (c Client) PostScheduleReload(ctx context.Context) (err error) {
	var s Status
	if err = c.httpPost(ctx, "/api/schedule/reload", nil, &s); err != nil {
		return err
	}

	if s.Status != "OK" {
		return errors.New(s.Message)
	}

	return nil
}

type ScheduleEntries []struct {
	Enabled          int      `json:"enabled"`
	Sequence         int      `json:"sequence"`
	Day              int      `json:"day"`
	StartTime        string   `json:"startTime"`
	StartTimeOffset  int      `json:"startTimeOffset"`
	EndTime          string   `json:"endTime"`
	EndTimeOffset    int      `json:"endTimeOffset"`
	Repeat           int      `json:"repeat"`
	StartDate        string   `json:"startDate"`
	EndDate          string   `json:"endDate"`
	StopType         int      `json:"stopType"`
	Playlist         string   `json:"playlist"`
	Command          string   `json:"command,omitempty"`
	Args             []string `json:"args,omitempty"`
	MultisyncCommand bool     `json:"multisyncCommand,omitempty"`
}
