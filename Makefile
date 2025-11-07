
mod:
	go build -buildmode=plugin -o plugins/user.so plugins/user/user.go
	go build -buildmode=plugin -o plugins/notification.so plugins/notification/notification.go

build: mod
	go build .

run: build
	./gomodss

clean:
	go clean
	rm plugins/*.so
