templ generate
bunx tailwindcss -i ./input.css -o ./static/css/style.css --minify
go build -o ./tmp/main.exe .
