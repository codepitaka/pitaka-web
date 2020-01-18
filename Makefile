run: bin/pitaka-web
	@PATH="$(PWD)/bin:$(PATH)" heroku local

bin/pitaka-web: main.go
	go build -o bin/pitaka-web main.go

clean:
	rm -rf bin