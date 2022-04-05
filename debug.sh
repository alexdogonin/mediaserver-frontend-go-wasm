cp wwwroot/* ./build && GOOS=js GOARCH=wasm go build -o ./build/app.wasm ./cmd/app && caddy run
