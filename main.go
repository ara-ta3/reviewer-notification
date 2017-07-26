package main

import (
	"fmt"
	"net/http"
)

// func readRequestBody() {
//     body, _ := ioutil.ReadAll(resp.Body)
//     defer resp.Body.Close()
//     r := response{}
//     json.Unmarshal(body, &r)
//     c := 0
// }

func main() {
	h := GithubNotificationHandler{}
	http.Handle("/", h)
	http.ListenAndServe(":8080", nil)
}

type GithubNotificationHandler struct {
	n ReviewerNotification
}

func (h GithubNotificationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
	g := GithubEvent{}
	e := h.n.Notify(g)
	if e != nil {
		fmt.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

type ReviewerNotification struct {
	s SlackClient
}

func (n ReviewerNotification) Notify(g GithubEvent) error {
	r := "ara-ta3"
	message := "hogehoge"
	return n.s.Send(r, message)
}

type GithubEvent struct {
}

type SlackClient struct {
	token string
}

func (s SlackClient) Send(reviewer, message string) error {
	return nil
}
