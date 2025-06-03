dev:
	wails dev

build:
	wails build -clean

windows:
	wails build -clean -nsis -trimpath -upx