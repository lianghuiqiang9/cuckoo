
go mod init cuckoo

go mod tidy

go test -bench=. -v -cpu 1

go run ./cmd