test:
	@echo 'Makefile is work'

run:
	docker-compose up --build project
stop:
	docker-compose stop