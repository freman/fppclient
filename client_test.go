package fppclient_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"

	"github.com/freman/fppclient"
	"github.com/freman/fppclient/debug/dumptransport"
)

func TestGetOverlaysModels(t *testing.T) {
	c, err := fppclient.New("http://10.0.0.249")
	require.NoError(t, err)

	models, err := c.GetOverlaysModels(context.TODO())
	require.NoError(t, err)

	spew.Dump(models)
}

func TestGetOverlaysModel(t *testing.T) {
	c, err := fppclient.New("http://10.0.0.249")
	require.NoError(t, err)

	model, err := c.GetOverlaysModel(context.TODO(), "LED Panels")
	require.NoError(t, err)

	spew.Dump(model)
}

func TestGetOverlaysModelData(t *testing.T) {
	c, err := fppclient.New("http://10.0.0.249")
	require.NoError(t, err)

	modelData, err := c.GetOverlaysModelData(context.TODO(), "LED Panels", false)
	require.NoError(t, err)

	modelDataRLE, err := c.GetOverlaysModelData(context.TODO(), "LED Panels", true)
	require.NoError(t, err)

	_, _ = modelData, modelDataRLE
}

func TestGetFonts(t *testing.T) {
	c, err := fppclient.New("http://10.0.0.249")
	require.NoError(t, err)

	fonts, err := c.GetOverlaysFonts(context.TODO())
	require.NoError(t, err)

	spew.Dump(fonts)
}

func TestFillModel(t *testing.T) {
	c, err := fppclient.New("http://10.0.0.249", fppclient.WithHTTPClient(&http.Client{Timeout: 10 * time.Second, Transport: &dumptransport.DumpTransport{http.DefaultTransport}}))
	require.NoError(t, err)

	require.NoError(t, c.FillOverlaysModel(context.TODO(), "LED Panels", 0, 0, 0))
}

func TestSetModelPixell(t *testing.T) {
	c, err := fppclient.New("http://10.0.0.249", fppclient.WithHTTPClient(&http.Client{Timeout: 10 * time.Second, Transport: &dumptransport.DumpTransport{http.DefaultTransport}}))
	require.NoError(t, err)

	require.NoError(t, c.SetOverlaysModelPixel(context.TODO(), "LED Panels", 0, 0, 90, 0, 0))
}
