
mod: 
	go build -buildmode=plugin -o user.so plugins/user/user.go
	go build -buildmode=plugin -o notification.so plugins/notification/notification.go

build: mod
	go build .

run: build
	./gomodss

clean:
	go clean
