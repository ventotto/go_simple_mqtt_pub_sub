.EXPORT_ALL_VARIABLES:
MQTT_URL=mqtt://test:test2@127.0.0.1:9999

all: deps run

deps:
	@dep ensure

build:
	@go build -o bin  main.go
	@go build -o bin  main_mqtt.go

run-main:
	@go run main.go

run-mqtt:
	@echo $$MQTT_URL
	@go run main_mqtt.go





