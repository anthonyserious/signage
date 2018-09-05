signage: main.go
	go get -d -v && go build

run:
	./signage

clean:
	rm -rf ./signage
