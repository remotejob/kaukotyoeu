all: push

# 0.0 shouldn't clobber any released builds
TAG =0.0
PREFIX = gcr.io/jntlserv0/godocker

binary: server.go
	CGO_ENABLED=0 GOOS=linux godep go build -a -installsuffix cgo -ldflags '-w' -o server

container: binary
	docker build -t $(PREFIX):$(TAG) .

push: container
	gcloud docker push $(PREFIX):$(TAG)

set: push
	 kubectl set image deployment/godocker godocker=$(PREFIX):$(TAG)

clean:
	docker rmi -f $(PREFIX):$(TAG) || true
