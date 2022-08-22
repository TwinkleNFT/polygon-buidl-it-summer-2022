
build: 
	env CGO_ENABLED=0 go build -o dist/app
	# env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o dist/app
	# env GOOS=linux CGO_ENABLED=0 go build -o dist/app
run: 
	go run main.go

docker-build: 
	docker build -t test .
docker-run:
	docker run -it test
docker-compose-up:
	docker-compose up
