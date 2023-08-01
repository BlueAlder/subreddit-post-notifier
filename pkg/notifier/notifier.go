package notifier

import (
	"fmt"
	"os/exec"

	"github.com/gen2brain/beeep"
	"github.com/rs/zerolog/log"
)

func OpenURLInBrowser(url string) {
	cmd := exec.Command("open", url)
	_, err := cmd.Output()
	if err != nil {
		log.Warn().Err(err).Msg(fmt.Sprintf("Error after opening URL %s", url))
	}
}

func SendOSNotification(title string, desciption string) {
	if err := beeep.Alert(title, desciption, ""); err != nil {
		log.Warn().Err(err).Msg("unable to send os notification")
	}

}

func SendDiscordWebhook(webhook string, url string) {
	log.Fatal().Msg("Unimplemented")
}

func SendEmail(email string, url string) {
	log.Fatal().Msg("Unimplemented")
}

func SendLetterInPost(postAddress string, url string) {
	log.Fatal().Msg("Unimplemented")
}

func SendTextMessage(phoneNumber string, url string) {
	log.Fatal().Msg("Unimplemented")
}
