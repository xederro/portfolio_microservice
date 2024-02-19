SHELL=cmd.exe
.PHONY: start stop watch build build-auth build-qrcode build-broker build-frontend

#starts containers
start:
	docker-compose up -d

#stops containers
stop:
	docker-compose down

#watches changes on codebase
watch:
	docker-compose watch

#builds all services
build:
	docker-compose build

#builds auth service
build-auth:
	docker-compose build auth

#builds qrcode service
build-qrcode:
	docker-compose build qrcode

#builds broker service
build-broker:
	docker-compose build broker

#builds frontend service
build-frontend:
	docker-compose build frontend
