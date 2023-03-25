build:
	docker build -t firestore-emulator:latest .

run:
	docker run -it --rm -p 8080:8080 -v $(pwd):/app firestore-emulator

init: 
	docker exec -it firestore-emulator /bin/bash -c "firebase login && firebase projects:list"

test:
	cd test
	npm install
	npm test


