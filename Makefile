build-php-base:
	$(eval TAGDATE := $(shell date "+%Y%m%d%H%M%S"))
	docker build -f ./infra/docker/php/base/Dockerfile -t kantaroso/kanta-workspace-php-base:latest .
	docker push kantaroso/kanta-workspace-php-base:latest
build-php-prod:
	$(eval TAGDATE := $(shell date "+%Y%m%d%H%M%S"))
	docker build -f ./infra/docker/php/laravel/Dockerfile -t kantaroso/kanta-workspace-php-prod:${TAGDATE} -t kantaroso/kanta-workspace-php-prod:latest .
	docker push kantaroso/kanta-workspace-php-prod:${TAGDATE}
	docker push kantaroso/kanta-workspace-php-prod:latest
build-go-pord:
	$(eval TAGDATE := $(shell date "+%Y%m%d%H%M%S"))
	docker run -it -v ${PWD}/app/gin:/go/src/kanta-workspace -w=/go/src/kanta-workspace golang:1.13.8-alpine3.11 go build main.go
	docker build -f ./infra/docker/go/gin/Dockerfile -t kantaroso/kanta-workspace-go-prod:${TAGDATE} -t kantaroso/kanta-workspace-go-prod:latest .
	docker push kantaroso/kanta-workspace-go-prod:${TAGDATE}
	docker push kantaroso/kanta-workspace-go-prod:latest
