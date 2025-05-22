build:
	docker build -t brandscout_test_task .
run:
	docker run -p 8080:8080 brandscout_test_task
test:	
	go test ./...