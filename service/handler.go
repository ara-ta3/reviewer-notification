package service

import (
	"log"
	"net/http"
)

type GithubNotificationHandler struct {
	NotificationService ReviewerNotification
	Logger              log.Logger
}

func (h GithubNotificationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h := r.Header.Get("X-Hub-Signature")
	e := h.NotificationService.NotifyWithRequestBody(r.Body, h)
	if e == nil {
		w.WriteHeader(http.StatusNoContent)
	} else {
		h.Logger.Printf("%+v\n", e)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
