.PHONY: build
build:
	docker build -t firestore-emulator:latest .

.PHONY: run
run:
	docker run -it --rm -p 8080:8080 -v $(pwd):/app firestore-emulator

.PHONY: test
test:
	cd test && \
	FIRESTORE_EMULATOR_HOST=localhost:8080 go run main.go

