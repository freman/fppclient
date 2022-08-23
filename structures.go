package fppclient

type Models []Model

type Model struct {
	ChannelCount        int    `json:"ChannelCount"`
	ChannelCountPerNode int    `json:"ChannelCountPerNode"`
	Name                string `json:"Name"`
	Orientation         string `json:"Orientation"`
	StartChannel        int    `json:"StartChannel"`
	StartCorner         string `json:"StartCorner"`
	StrandsPerString    int    `json:"StrandsPerString"`
	StringCount         int    `json:"StringCount"`
	Type                string `json:"Type"`
	AutoCreated         bool   `json:"autoCreated"`
	EffectRunning       bool   `json:"effectRunning"`
	Height              int    `json:"height"`
	IsActive            int    `json:"isActive"`
	Width               int    `json:"width"`
}

type ModelData struct {
	Data          []int `json:"data"`
	EffectRunning bool  `json:"effectRunning"`
	IsLocked      bool  `json:"isLocked"`
	RLE           bool  `json:"rle"`
}

type Fonts []string
type Plugins []string

func (p Plugins) Contains(plugin string) bool {
	for _, v := range p {
		if v == plugin {
			return true
		}
	}
	return false
}

type Status struct {
	Status  string `json:"Status"`
	Message string `json:"message"`
}

type fillPixelRequest struct {
	X   int   `json:"X,omitempty"`
	Y   int   `json:"Y,omitempty"`
	RGB []int `json:"RGB"`
}

type ChannelOutputsObj struct {
	ChannelOutputs `json:"channelOutputs"`
}

type ChannelOutputs []ChannelOutput

type ChannelOutput struct {
	Type                string              `json:"type"`
	SubType             string              `json:"subType"`
	Enabled             int                 `json:"enabled"`
	CfgVersion          int                 `json:"cfgVersion"`
	StartChannel        int                 `json:"startChannel"`
	ChannelCount        int                 `json:"channelCount"`
	ColorOrder          string              `json:"colorOrder"`
	Gamma               string              `json:"gamma"`
	WiringPinout        string              `json:"wiringPinout"`
	Brightness          int                 `json:"brightness"`
	PanelColorDepth     int                 `json:"panelColorDepth"`
	InvertedData        int                 `json:"invertedData"`
	PanelWidth          int                 `json:"panelWidth"`
	PanelHeight         int                 `json:"panelHeight"`
	PanelScan           int                 `json:"panelScan"`
	PanelOutputOrder    bool                `json:"panelOutputOrder"`
	PanelOutputBlankRow bool                `json:"panelOutputBlankRow"`
	Panels              ChannelOutputPanels `json:"panels"`
}

type ChannelOutputPanels []ChannelOutputPanel

type ChannelOutputPanel struct {
	OutputNumber int    `json:"outputNumber"`
	PanelNumber  int    `json:"panelNumber"`
	ColorOrder   string `json:"colorOrder"`
	XOffset      int    `json:"xOffset"`
	YOffset      int    `json:"yOffset"`
	Orientation  string `json:"orientation"`
	Row          int    `json:"row"`
	Col          int    `json:"col"`
}
