package main

import (
	"encoding/json"
	"net/http"

	"github.com/sugyan/clova-extension-example/clova/cek"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func main() {
	http.HandleFunc("/callback", callbackHandler)
	appengine.Main()
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	// TODO check header
	ctx := appengine.NewContext(r)

	var message cek.RequestMessage
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		log.Errorf(ctx, "error: %s", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	log.Infof(ctx, "request: %#v", message)

	response := &cek.ResponseMessage{
		Response: &cek.Response{
			Card:       struct{}{},
			Directives: []*cek.Directive{},
			OutputSpeech: &cek.OutputSpeech{
				Type: cek.OutputSpeechTypeSpeechList,
				Values: []*cek.SpeechInfo{
					&cek.SpeechInfo{
						Lang:  cek.SpeechInfoLangJA,
						Type:  cek.SpeechInfoTypePlainText,
						Value: "サンプルのレスポンスです",
					},
				},
			},
			Reprompt: &cek.Reprompt{
				OutputSpeech: &cek.OutputSpeech{
					Type: cek.OutputSpeechTypeSpeechList,
					Values: []*cek.SpeechInfo{
						&cek.SpeechInfo{
							Lang:  cek.SpeechInfoLangJA,
							Type:  cek.SpeechInfoTypePlainText,
							Value: "これはリプロンプトです",
						},
					},
				},
			},
		},
		SessionAttributes: map[string]string{},
		Version:           "1.0",
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Errorf(ctx, "error: %s", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
