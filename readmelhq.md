
go mod init cuckoo

go mod tidy

go test -bench=. -v -cpu 1

go run ./cmd


# 可以修改使用

1. 处理key而放弃value
2. 编写 simple hashing 
3. 然后调用单个桶中的PSI就好了。
4. 能否将key的length调整为64位？