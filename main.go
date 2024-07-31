package main

import (
	"log"

	"gotexttospeech/htgotts"

	"github.com/hegedustibor/htgo-tts/handlers"
)

func main() {
	speech := htgotts.Speech{
		Folder:   "audio",
		Language: "en",
		Handler:  &handlers.MPlayer{},
	}

	err := speech.Speak("Hello there")
	if err != nil {
		log.Fatal(err)
	}
}
