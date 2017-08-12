GODEP=$(shell go env GOPATH)/bin/godep
slack_webhook_url=http://localhost
auth_token=token
labels=S-awaiting-review

run:
	env SLACK_WEBHOOK_URL=$(slack_webhook_url) \
		TOKEN=$(auth_token) \
		TARGET_LABELS=$(labels) \
		go run main.go

vendor/save: $(GODEP)
	$(GODEP) save ./...

curl:
	curl -i localhost:8080

deploy:
	git push heroku master 

$(GODEP):
	go get -u github.com/tools/godep
