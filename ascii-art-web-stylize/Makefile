build:
	docker build --rm -t ascii-image .
run:
	docker run -p 8080:8080 --rm --name ascii ascii-image
run-dev:
	docker run -d -p 8080:8080 --rm --name ascii ascii-image
	
stop:
	docker stop ascii