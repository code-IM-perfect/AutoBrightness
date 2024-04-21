go build -o build/autoBrightness .

env GOOS=windows GOARCH=amd64 go build -o build/autoBrightness-windows-x86_64.exe .

env GOOS=darwin GOARCH=arm64 go build -o build/autoBrightness-mac-arm64 .