package dumptransport

import (
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/fatih/color"
)

type DumpTransport struct {
	Transport http.RoundTripper
}

func (d *DumpTransport) RoundTrip(h *http.Request) (*http.Response, error) {
	dump, _ := httputil.DumpRequestOut(h, true)
	color.HiRed("**** REQUEST ****")
	color.Red(string(dump))
	fmt.Println("")

	resp, err := d.Transport.RoundTrip(h)
	dump, _ = httputil.DumpResponse(resp, true)
	color.HiGreen("**** RESPONSE ****")
	color.Green(string(dump))
	color.HiBlue("\n********************")
	fmt.Println("")
	return resp, err
}
