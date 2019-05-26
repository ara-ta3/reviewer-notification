GO=go
HEROKU=$(shell which heroku)
goos_opt=GOOS=$(GOOS)
goarch_opt=GOARCH=$(GOARCH)
out=reviewer-notification
out_opt="-o $(out)"
slack_webhook_url=http://localhost
auth_token=token
labels=S-awaiting-review
map="ara-ta3:arata,hogehoge:fugafuga"
post_to_channel=
port=8080
host=localhost
path=
url=http://$(host):$(port)/$(path)

install:
	$(GO) mod vendor

run/binary: build_for_local
	./$(out)

run:
	env SLACK_WEBHOOK_URL=$(slack_webhook_url) \
		TOKEN=$(auth_token) \
		TARGET_LABELS=$(labels) \
		ACCOUNT_MAP=$(map) \
		SLACK_CHANNEL=$(post_to_channel) \
		PORT=$(port) \
		go run main.go

curl:
	curl -i $(url)

curl/post:
	curl -i $(url) -d '$(shell cat ./sample.json)' -X POST

build:
	$(goos_opt) $(goarch_opt) go build $(out_opt)

build_for_linux:
	$(MAKE) build GOOS=linux GOARCH=amd64 out_opt=""

build_for_local:
	$(MAKE) build goos_opt= goarch_opt= out_opt=

deploy:
	git push heroku master 

open: 
	$(HEROKU) open
