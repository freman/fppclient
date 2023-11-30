package fppclient_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/freman/fppclient"
	"github.com/stretchr/testify/require"
)

func TestFPPTime(t *testing.T) {
	checks := []struct {
		In  string
		Out fppclient.FPPTime
	}{{
		`"Sun Dec  1 @ 07:00 PM"`,
		fppclient.FPPTime{time.Date(2023, time.December, 1, 19, 0, 0, 0, time.UTC)},
	}, {
		`"Sun Dec 10 @ 07:00 PM"`,
		fppclient.FPPTime{time.Date(2023, time.December, 10, 19, 0, 0, 0, time.UTC)},
	}}

	for _, check := range checks {
		var out fppclient.FPPTime

		err := json.NewDecoder(strings.NewReader(check.In)).Decode(&out)
		require.NoError(t, err)
		require.Equal(t, check.Out, out)
	}
}
