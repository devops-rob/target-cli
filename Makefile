version = 1.0.0

build:
	go build -o /usr/local/bin/target .
multi_build:
	@echo ""
	@echo "Compile cli"

	# Clear the output
	rm -rf ./bin

	GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o ./bin/linux_arm64/target_v$(version) main.go
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./bin/linux_amd64/target_v$(version) main.go
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -o ./bin/darwin_arm64/target_v$(version) main.go
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o ./bin/darwin_amd64/target_v$(version) main.go
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o ./bin/windows_amd64/target_v$(version).exe main.go
	GOOS=windows GOARCH=386 CGO_ENABLED=0 go build -o ./bin/windows_386/target_v$(version).exe main.go
zip:
	pwd
	zip -j ./bin/target_v$(version)_linux_arm64.zip ./bin/linux_arm64/target_v$(version)
	zip -j ./bin/target_v$(version)_linux_amd64.zip ./bin/linux_amd64/target_v$(version)
	zip -j ./bin/target_v$(version)_darwin_arm64.zip ./bin/linux_arm64/target_v$(version)
	zip -j ./bin/target_v$(version)_darwin_amd64.zip ./bin/linux_arm64/target_v$(version)
	zip -j ./bin/target_windows_amd64.zip ./bin/windows_amd64/target_v$(version).exe
	zip -j ./bin/target_v$(version)_windows_386.zip ./bin/windows_386/target_v$(version).exe
	ls -lha ./bin
shasum:
	cd bin/; shasum -a 256 *.zip > target_v$(version)_SHA256SUMS
gpg:
	gpg --detach-sign ./bin/target_v$(version)_SHA256SUMS