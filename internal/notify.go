package internal

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/smtp"

	"github.com/gin-gonic/gin"
)

func Notify(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	// Set up the authentication information
	auth := smtp.PlainAuth("", "avalama86@gmail.com", "mrjnhnokbhjtwksf", "smtp.gmail.com")

	// Set up the email message
	// to := []string{"romit.tajale04@gmail.com", "romittajale69@gmail.com", "<l.wong@computerfutures.com>", "<daniel.neil@turnberrysolutions.com>", "<t.kallen@computerfutures.com>", "<Deborah.Larrauri@theksquaregroup.com>", "<andre.lee@orangepeople.com>"}
	to := []string{"romit.tajale04@gmail.com", "romittajale69@gmail.com", "susanmaharjan03030@gmail.com"}
	msg := []byte(fmt.Sprintf("To: %s\r\nSubject: Notification\r\n\r\n%s", to[0], body))

	// Send the email message
	for i := 0; i < 10; i++ {
		err = smtp.SendMail("smtp.gmail.com:587", auth, "sender@example.com", to, msg)
		if err != nil {
			log.Println(err)
		}

	}
}
