go build -o build/autoBrightness .

env GOOS=windows GOARCH=amd64 go build -o build/autoBrightness_windows_amd64.exe .

env GOOS=darwin GOARCH=arm64 go build -o build/autoBrightness_mac_arm64 .