package helpers

import (
	"gopkg.in/gomail.v2"
)

func SendGoMail(email string, token string) error {
	m := gomail.NewMessage()

	m.SetHeader("From", "admin@example.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Hello!")

	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")

	// Set the HTML version of the email
	m.AddAlternative("text/html", `
        <html>
            <body>
                <h1 style="color: red;">This is a Test Email</h1>
                <p><b>Hello!</b> This is a test email with HTML formatting.</p>
                <p>Thanks,<br>Mailtrap</p>
            </body>
        </html>
    `)

	d := gomail.NewDialer("smtp.mailtrap.io", 587, "b9ae76c0da012f", "c380cda4abcab9")

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
