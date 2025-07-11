# run templ generation in watch mode to detect all .templ files and
# re-create _templ.txt files on change, then send reload event to browser.
# Default url: http://localhost:7331
live/templ:
	templ generate --watch --proxy="http://localhost:3000" --open-browser=false -v

# run air to detect any go file changes to re-build and re-run the server.
live/server:
	air \
	--build.cmd "New-Item -Path Env:\BUILD_MODE -Value "develop"; go build -o tmp/bin/main.exe" --build.bin "tmp/bin/main.exe" --build.delay "100"; Remove-Item -Path Env:\BUILD_MODE -Verbose \
	--build.exclude_dir "node_modules" \
	--build.include_ext "go" \
	--build.stop_on_error "false" \
	--misc.clean_on_exit true

# run tailwindcss to generate the styles.css bundle in watch mode.
live/tailwind:
	bunx tailwindcss -i ./input.css -o ./static/css/styles.css --minify --watch

# watch for any js or css change in the assets/ folder, then reload the browser via templ proxy.
# --build.bin "true" 
live/sync_assets:
	air \
	--build.cmd "bunx tailwindcss -i ./input.css -o ./static/css/styles.css --minify; templ generate --notify-proxy" \
	--build.delay "100" \
	--build.exclude_dir "" \
	--build.include_dir "static" \
	--build.include_ext "js,css"

# start all 5 watch processes in parallel.
# live/tailwind
dev:
	make -j3 live/templ live/server live/sync_assets

run:
	bunx tailwindcss -i ./input.css -o ./static/css/style.css --minify && templ generate && go run .

build:
	go generate
	go build