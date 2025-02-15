main_package_path = ./
binary_name = clientele
tmp_dir = ./tmp

.PHONY: wasm/build
wasm/build:
	GOOS=js GOARCH=wasm go build -o ${tmp_dir}/clientele.wasm ./cmd/client

.PHONY: wasm/deploy
wasm/deploy: wasm/build
	cp ${tmp_dir}/clientele.wasm web/static/

.PHONY: audit
audit: test
	go mod tidy -diff
	go mod verify
	test -z "$(shell gofmt -l .)"
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...

.PHONY: test
test:
	go test -v -race -buildvcs ./...

.PHONY: test/cover
test/cover:
	go test -v -race -buildvcs -coverprofile=${tmp_dir}/coverage.out ./...
	go tool cover -html=${tmp_dir}/coverage.out

.PHONY: tidy
tidy:
	go mod tidy -v
	go fmt ./...

.PHONY: build
build:
	# Include additional build steps, like TypeScript, SCSS or Tailwind compilation here...
	go build -o=${tmp_dir}/bin/${binary_name} ${main_package_path}

.PHONY: run
run: build
	/tmp/bin/${binary_name}

.PHONY: run/live
run/live:
	go run github.com/cosmtrek/air@v1.43.0 \
		--build.cmd "make build" --build.bin "${tmp_dir}/bin/${binary_name}" --build.delay "100" \
		--build.exclude_dir "" \
		--build.include_ext "go, tpl, tmpl, html, css, scss, js, ts, sql, jpeg, jpg, gif, png, bmp, svg, webp, ico" \
		--misc.clean_on_exit "true"

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

.PHONY: no-dirty
no-dirty:
	@test -z "$(shell git status --porcelain)"
