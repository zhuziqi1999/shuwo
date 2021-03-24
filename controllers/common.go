package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/zhuziqi1999/shuwo/helpers"
	"github.com/zhuziqi1999/shuwo/system"
	"net/http"

	"strings"
)

const (
	SESSION_KEY          = "UserID"       // session key
	CONTEXT_USER_KEY     = "User"         // context user key
	SESSION_GITHUB_STATE = "GITHUB_STATE" // github state session key
	SESSION_CAPTCHA      = "GIN_CAPTCHA"  // captcha session key
)

func Handle404(c *gin.Context) {
	HandleMessage(c, "Sorry,I lost myself!")
}

func HandleMessage(c *gin.Context, message string) {
	c.HTML(http.StatusNotFound, "errors/error.html", gin.H{
		"message": message,
	})
}

func sendMail(to, subject, body string) error {
	c := system.GetConfiguration()
	return helpers.SendToMail(c.SmtpUsername, c.SmtpPassword, c.SmtpHost, to, subject, body, "html")
}

func NotifyEmail(subject, body string) error {
	notifyEmailsStr := system.GetConfiguration().NotifyEmails
	if notifyEmailsStr != "" {
		notifyEmails := strings.Split(notifyEmailsStr, ";")
		emails := make([]string, 0)
		for _, email := range notifyEmails {
			if email != "" {
				emails = append(emails, email)
			}
		}
		if len(emails) > 0 {
			return sendMail(strings.Join(emails, ";"), subject, body)
		}
	}
	return nil
}

func writeJSON(ctx *gin.Context, h gin.H) {
	if _, ok := h["succeed"]; !ok {
		h["succeed"] = false
	}
	ctx.JSON(http.StatusOK, h)
}
