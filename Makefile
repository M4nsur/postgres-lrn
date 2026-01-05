service-run:
	export conn=postgres://postgres:el0404@localhost:5432/postgres &&
	go run main.go