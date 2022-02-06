package handler

import (
    "inquiry/domain"
    "os"

    "github.com/gin-gonic/gin"

    sendgrid "github.com/sendgrid/sendgrid-go"
    "github.com/sendgrid/sendgrid-go/helpers/mail"
)

// SendMail - send mail include body data
func SendMail(c *gin.Context) {
    var m domain.Mail
    err := c.BindJSON(&m)
    if err != nil {
        c.JSON(http.InternalServerError, err.Error())
    }

    // set val to send mail
    from := mail.NewEmail(m.Name, m.Email)
    subject := m.Subject
    to := mail.NewEmail("admin", "haruki.lo.shelon@gmail.com") // メールを受けたい名前とアドレスを指定
    plainTextContent := m.Text
    htmlContent := m.Text

    // create message
    message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

    // send mail
    client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
    response, err := client.Send(message)
    if err != nil {
        c.JSON(response.StatusCode, err.Error())
    }
    c.JSON(response.StatusCode, "success to send mail !!")
}
