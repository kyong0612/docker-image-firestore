build:
	docker build -t firestore-emulator .

run:
	docker run -it --rm -p 8080:8080 -v $(pwd):/app firestore-emulator


test:
	cd test && npm install && npm test