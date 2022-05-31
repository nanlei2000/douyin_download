echo "start test"
go test ./...
echo "start build"
go build -ldflags="-s -w" -o dydl.exe ./cmd/.