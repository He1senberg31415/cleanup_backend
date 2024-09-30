package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
)

func main() {
	app := pocketbase.New()

	// app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
	// 	e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))
	// 	return nil
	// })

	app.OnMailerBeforeRecordVerificationSend().Add(func(e *core.MailerRecordEvent) error {
		// fmt.Println("UserID:", e.Record.Get("verificationCode"))
		// fmt.Println("Email html", e.Message.HTML)
		html := e.Message.HTML
		re := regexp.MustCompile(`%20%28.*?%29`)
		cleanText := re.ReplaceAllString(html, "")
		// fmt.Println(cleanText)
		replacedText := strings.Replace(cleanText, "CODE", fmt.Sprintf("%v", e.Record.Get("verificationCode")), 1)

		e.Message.HTML = replacedText

		return nil
	})

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/verify", func(c echo.Context) error {
			authRecord, _ := c.Get(apis.ContextAuthRecordKey).(*models.Record)
			verificationCode := c.QueryParam("verificationCode")
			fmt.Println("got request")
			fmt.Println(verificationCode)

			if authRecord != nil {
				fmt.Println("test1")
				username := authRecord.Username()
				fmt.Println(username)
				return c.String(http.StatusOK, "Authenticated user: ")
			}
			fmt.Println("test2")

			return c.String(http.StatusUnauthorized, "Unauthorized")
		})
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
