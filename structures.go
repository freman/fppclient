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

type ScheduleResponse struct {
	Status
	RespCode int      `json:"respCode"`
	Schedule Schedule `json:"schedule"`
}

type Schedule struct {
	Enabled int `json:"enabled"`
	Entries []struct {
		Args             []string `json:"args,omitempty"`
		Command          string   `json:"command,omitempty"`
		Day              int      `json:"day"`
		DayStr           string   `json:"dayStr"`
		Enabled          int      `json:"enabled"`
		EndDate          string   `json:"endDate"`
		EndTime          string   `json:"endTime"`
		ID               int      `json:"id"`
		MultisyncCommand bool     `json:"multisyncCommand,omitempty"`
		MultisyncHosts   string   `json:"multisyncHosts,omitempty"`
		Playlist         string   `json:"playlist"`
		Repeat           int      `json:"repeat"`
		RepeatInterval   int      `json:"repeatInterval"`
		StartDate        string   `json:"startDate"`
		StartTime        string   `json:"startTime"`
		StopType         int      `json:"stopType"`
		StopTypeStr      string   `json:"stopTypeStr"`
		Type             string   `json:"type"`
	} `json:"entries"`
	Items []struct {
		Args             []string `json:"args"`
		Command          string   `json:"command"`
		EndTime          int      `json:"endTime"`
		EndTimeStr       string   `json:"endTimeStr"`
		ID               int      `json:"id"`
		MultisyncCommand bool     `json:"multisyncCommand"`
		MultisyncHosts   string   `json:"multisyncHosts"`
		Priority         int      `json:"priority"`
		StartTime        int      `json:"startTime"`
		StartTimeStr     string   `json:"startTimeStr"`
	} `json:"items"`
}

type FPPDStatus struct {
	MQTT struct {
		Configured bool `json:"configured"`
		Connected  bool `json:"connected"`
	} `json:"MQTT"`
	Bridging        bool `json:"bridging"`
	CurrentPlaylist struct {
		Count       string `json:"count"`
		Description string `json:"description"`
		Index       string `json:"index"`
		Playlist    string `json:"playlist"`
		Type        string `json:"type"`
	} `json:"current_playlist"`
	CurrentSequence string `json:"current_sequence"`
	CurrentSong     string `json:"current_song"`
	DateStr         string `json:"dateStr"`
	Fppd            string `json:"fppd"`
	Mode            int    `json:"mode"`
	ModeName        string `json:"mode_name"`
	Multisync       bool   `json:"multisync"`
	NextPlaylist    struct {
		Playlist  string `json:"playlist"`
		StartTime string `json:"start_time"`
	} `json:"next_playlist"`
	RepeatMode string `json:"repeat_mode"`
	Scheduler  struct {
		Enabled      int `json:"enabled"`
		NextPlaylist struct {
			PlaylistName          string `json:"playlistName"`
			ScheduledStartTime    int    `json:"scheduledStartTime"`
			ScheduledStartTimeStr string `json:"scheduledStartTimeStr"`
		} `json:"nextPlaylist"`
		Status string `json:"status"`
	} `json:"scheduler"`
	SecondsPlayed    string `json:"seconds_played"`
	SecondsRemaining string `json:"seconds_remaining"`
	Sensors          []struct {
		Formatted string  `json:"formatted"`
		Label     string  `json:"label"`
		Postfix   string  `json:"postfix"`
		Prefix    string  `json:"prefix"`
		Value     float64 `json:"value"`
		ValueType string  `json:"valueType"`
	} `json:"sensors"`
	Status             int      `json:"status"`
	StatusName         string   `json:"status_name"`
	Time               string   `json:"time"`
	TimeStr            string   `json:"timeStr"`
	TimeStrFull        string   `json:"timeStrFull"`
	TimeElapsed        string   `json:"time_elapsed"`
	TimeRemaining      string   `json:"time_remaining"`
	Uptime             string   `json:"uptime"`
	UptimeDays         float64  `json:"uptimeDays"`
	UptimeHours        float64  `json:"uptimeHours"`
	UptimeMinutes      float64  `json:"uptimeMinutes"`
	UptimeSeconds      int      `json:"uptimeSeconds"`
	UptimeStr          string   `json:"uptimeStr"`
	UptimeTotalSeconds int      `json:"uptimeTotalSeconds"`
	UUID               string   `json:"uuid"`
	Volume             int      `json:"volume"`
	Warnings           []string `json:"warnings"`
}
