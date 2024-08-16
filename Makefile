run:
	clear
	templ generate
	ls -l
	go build -ldflags="-w -s" .
	./gosqlitetempl
