# run templ generation in watch mode to detect all .templ files and
# re-create _templ.txt files on change, then send reload event to browser.
# Default url: http://localhost:7331
live/templ:
	templ generate --watch --proxy="http://localhost:3000" --open-browser=false -v

# run air to detect any go file changes to re-build and re-run the server.
live/server:
	air
	--build.full_bin "BUILD_MODE=develop go build -o tmp/bin/main.exe" --build.bin "tmp/bin/main.exe" --build.delay "100" \
	--build.exclude_dir "node_modules" \
	--build.include_ext "go" \
	--build.stop_on_error "false" \
	--misc.clean_on_exit true

# run tailwindcss to generate the styles.css bundle in watch mode.
live/tailwind:
	bunx tailwindcss -i ./input.css -o ./static/css/style.css --watch=forever

# watch for any js or css change in the assets/ folder, then reload the browser via templ proxy.
live/sync_assets:
	air
	--build.cmd "templ generate --notify-proxy" \
	--build.bin "true" \
	--build.delay "100" \
	--build.exclude_dir "" \
	--build.exclude_dir "node_modules" \
	--build.include_dir "static" \
	--build.include_ext "js,css"

# start all 5 watch processes in parallel.
dev:
	make -j4 live/tailwind live/templ live/server live/sync_assets

build:
	go generate
	go build