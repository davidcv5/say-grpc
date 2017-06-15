build:
	GOOS=linux go build -o app
	docker build -t gcr.io/demo24x7/say .
	rm -f app
push:
	gcloud docker -- push gcr.io/demo24x7/say