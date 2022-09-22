.PHONY: build list github

VERSION := "0.0.1"

VARS    := -w -s -H=windowsgui

build:
	@cd internal/cmd/fyne-iconpark && go build -ldflags "${VARS}" -o "../../../bin/v${VERSION}/fyne-iconpark-${VERSION}.exe" && upx --best "../../../bin/v${VERSION}/fyne-iconpark-${VERSION}.exe"

github:
	@cd internal/cmd/fyne-iconpark && go build -ldflags "${VARS}" -o "../../../bin/fyne-iconpark-${VERSION}.exe"

list:
	@echo ${VERSION}
	@echo ${VARS}

vet:
	go vet
