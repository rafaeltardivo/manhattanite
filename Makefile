build:
	docker build . --no-cache -t manhattanite:1.0
up:
	docker run -p 8080:8080  manhattanite:1.0