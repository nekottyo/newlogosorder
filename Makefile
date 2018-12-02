
build:
	docker build --tag newlogosorder_app .

up: build
	docker-compose up -d

rm:
	docker rm server
