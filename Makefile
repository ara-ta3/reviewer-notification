GODEP=$(shell go env GOPATH)/bin/godep
HEROKU=$(shell which heroku)
slack_webhook_url=http://localhost
auth_token=token
labels=S-awaiting-review
port=8080
host=localhost
path=
url=http://$(host):$(port)/$(path)

run:
	env SLACK_WEBHOOK_URL=$(slack_webhook_url) \
		TOKEN=$(auth_token) \
		TARGET_LABELS=$(labels) \
		PORT=$(port) \
		go run main.go

vendor/save: $(GODEP)
	$(GODEP) save ./...

curl:
	curl -i $(url)

deploy:
	git push heroku master 

open: 
	$(HEROKU) open

$(GODEP):
	go get -u github.com/tools/godep
