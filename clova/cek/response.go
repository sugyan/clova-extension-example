package cek

// OutputSpeechType type
type OutputSpeechType string

// OutputSpeechType constants
const (
	OutputSpeechTypeSimpleSpeech OutputSpeechType = "SimpleSpeech"
	OutputSpeechTypeSpeechList   OutputSpeechType = "SpeechList"
	OutputSpeechTypeSpeechSet    OutputSpeechType = "SpeechSet"
)

// OutputSpeechVerboseType type
type OutputSpeechVerboseType string

// OutputSpeechVerboseType constants
const (
	OutputSpeechVerboseTypeSimpleSpeech OutputSpeechVerboseType = "SimpleSpeech"
	OutputSpeechVerboseTypeSpeechList   OutputSpeechVerboseType = "SpeechList"
)

// SpeechInfoLang type
type SpeechInfoLang string

// SpeechInfoLang constants
const (
	SpeechInfoLangEN    SpeechInfoLang = "en"
	SpeechInfoLangJA    SpeechInfoLang = "ja"
	SpeechInfoLangKO    SpeechInfoLang = "ko"
	SpeechInfoLangEmpty SpeechInfoLang = ""
)

// SpeechInfoType type
type SpeechInfoType string

// SpeechInfoType constants
const (
	SpeechInfoTypePlainText SpeechInfoType = "PlainText"
	SpeechInfoTypeURL       SpeechInfoType = "URL"
)

// ResponseMessage type
type ResponseMessage struct {
	Response          *Response         `json:"response"`
	SessionAttributes map[string]string `json:"sessionAttributes"`
	Version           string            `json:"version"`
}

// Response type
type Response struct {
	Card             interface{}   `json:"card"`
	Directives       []*Directive  `json:"directives"`
	OutputSpeech     *OutputSpeech `json:"outputSpeech"`
	Reprompt         *Reprompt     `json:"reprompt,omitempty"`
	ShouldEndSession bool          `json:"shouldEndSession"`
}

// Directive type
type Directive struct {
	Header *Header `json:"header"`
}

// Header type
type Header struct {
	MessageID string   `json:"messageId"`
	Name      string   `json:"name"`
	Namespace string   `json:"namespace"`
	Payload   *Payload `json:"payload"`
}

// Payload type
type Payload interface{}

// OutputSpeech type
type OutputSpeech struct {
	Brief   *SpeechInfo      `json:"brief,omitempty"`
	Type    OutputSpeechType `json:"type"`
	Values  []*SpeechInfo    `json:"values,omitempty"`
	Verbose *Verbose         `json:"verbose,omitempty"`
}

// SpeechInfo type
type SpeechInfo struct {
	Lang  SpeechInfoLang `json:"lang"`
	Type  SpeechInfoType `json:"type"`
	Value string         `json:"value"`
}

// Verbose type
type Verbose struct {
	Type   OutputSpeechVerboseType `json:"type"`
	Values []*SpeechInfo           `json:"values"`
}

// Reprompt type
type Reprompt struct {
	OutputSpeech *OutputSpeech `json:"outputSpeech"`
}
