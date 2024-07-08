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

	err := speech.Speak("Awesome! Is there anything specific you'd like to talk about or need help with?")
	if err != nil {
		log.Fatal(err)
	}
}
