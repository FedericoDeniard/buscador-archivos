build:
	go build -o dist/buscador src/main.go
install:
	sudo cp dist/buscador /usr/local/bin
uninstall:
	sudo rm /usr/local/bin/buscador
	