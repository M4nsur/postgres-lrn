service-run:
	export conn="postgres://postgres:el0404@localhost:5432/postgres?sslmode=disable" && \
	go run main.go

some-target:
	export env="2213123" && \
	printenv env