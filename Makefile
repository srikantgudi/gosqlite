run:
	clear
	rm views/*_templ.go
	templ generate
	ls -l
	go build -ldflags="-w -s" .
	./gosqlite
