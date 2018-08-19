package cek

// PlayerActivity type
type PlayerActivity string

// PlayerActivity constants
const (
	PlayerActivityIDLE    PlayerActivity = "IDLE"
	PlayerActivityPLAYING PlayerActivity = "PLAYING"
	PlayerActivityPAUSED  PlayerActivity = "PAUSED"
	PlayerActivitySTOPPED PlayerActivity = "STOPPED"
)

// Orientation type
type Orientation string

// Orientation constants
const (
	OrientationLandscape Orientation = "landscape"
	OrientationPortrait  Orientation = "portrait"
)

// Size type
type Size string

// Size constants
const (
	SizeNone   Size = "none"
	SizeS100   Size = "s100"
	SizeM100   Size = "m100"
	SizeL100   Size = "l100"
	SizeXL100  Size = "xl100"
	SizeCustom Size = "custom"
)

// RequestType type
type RequestType string

// RequestType constants
const (
	RequestTypeEvent        RequestType = "EventRequest"
	RequestTypeIntent       RequestType = "IntentRequest"
	RequestTypeLaunch       RequestType = "LaunchRequest"
	RequestTypeSessionEnded RequestType = "SessionEndedRequest"
)

// RequestMessage type
type RequestMessage struct {
	Context *Context `json:"context"`
	Request *Request `json:"request"`
	Session *Session `json:"session"`
	Version string   `json:"version"`
}

// Context type
type Context struct {
	AudioPlayer *AudioPlayer `json:"AudioPlayer,omitempty"`
	System      *System      `json:"System"`
}

// AudioPlayer type
type AudioPlayer struct {
	OffsetInMilliseconds int            `json:"offsetInMilliseconds,omitempty"`
	PlayerActivity       PlayerActivity `json:"playerActivity"`
	Stream               interface{}    `json:"stream,omitempty"`
	TotalInMilliseconds  int            `json:"totalInMilliseconds,omitempty"`
}

// System type
type System struct {
	Application *Application `json:"application"`
	Device      *Device      `json:"device"`
	User        *User        `json:"user"`
}

// Application type
type Application struct {
	ApplicationID string `json:"applicationId"`
}

// Device type
type Device struct {
	DeviceID string   `json:"deviceId"`
	Display  *Display `json:"display"`
}

// Display type
type Display struct {
	ContentLayer *ContentLayer `json:"contentLayer,omitempty"`
	DPI          *int          `json:"dpi,omitempty"`
	Orientation  *Orientation  `json:"orientation,omitempty"`
	Size         Size          `json:"size"`
}

// ContentLayer type
type ContentLayer struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

// User type
type User struct {
	UserID      string  `json:"userId"`
	AccessToken *string `json:"accessToken,omitempty"`
}

// Request type
type Request struct {
	Type   RequestType `json:"type"`
	Intent interface{} `json:"intent"`
}

// Session type
type Session struct {
	New               bool              `json:"new"`
	SessionAttributes map[string]string `json:"sessionAttributes"`
	SessionID         string            `json:"sessionId"`
	User              *User             `json:"user"`
}
