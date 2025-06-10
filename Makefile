# Checks and installs tools

check-air: # Install air if not already installed.
	@if ! [ -e "$(GOBIN)/air" ]; then \
		echo "go install github.com/air-verse/air@latest"; \
		go install github.com/air-verse/air@latest; \
	fi

check-templ:
	@if ! [ -e "$(GOBIN)/templ" ]; then \
		echo "go install github.com/a-h/templ/cmd/templ@latest"; \
		go install github.com/a-h/templ/cmd/templ@latest; \
	fi

check-delve: # Install dlv if not already installed.
	@if ! [ -e "$(GOBIN)/dlv" ]; then \
		echo "go install github.com/go-delve/delve/cmd/dlv@latest"; \
		go install github.com/go-delve/delve/cmd/dlv@latest; \
	fi

check-pkgsite: # Install pkgsite if not already installed.
	@if ! [ -e "$(GOBIN)/pkgsite" ]; then \
		echo "go install golang.org/x/pkgsite/cmd/pkgsite@latest"; \
		go install golang.org/x/pkgsite/cmd/pkgsite@latest; \
	fi

check-gosec: # Install gosec if not already installed.
	@if ! [ -e "$(GOBIN)/gosec" ]; then \
		echo "go install github.com/securego/gosec/v2/cmd/gosec@latest"; \
		go install github.com/securego/gosec/v2/cmd/gosec@latest; \
	fi

check-govuln: # Install govulncheck if not already installed.
	@if ! [ -e "$(GOBIN)/govulncheck" ]; then \
		echo "go install golang.org/x/vuln/cmd/govulncheck@latest"; \
		go install golang.org/x/vuln/cmd/govulncheck@latest; \
	fi


# Common commands

.PHONY: test
test:
	@go test ./...

.PHONY: build
build:
	@go build -o ./bin/shortlink ./main.go

.PHONY: docs
docs: check-pkgsite
	@pkgsite -http=localhost:6060 -open .

.PHONY: debug
debug: check-delve
	@dlv debug ./main.go


# Check vulnerabilities

.PHONY: gosec
gosec: check-gosec
	@gosec -tests -exclude-dir=.ignore ./...

.PHONY: govuln
govuln: check-govuln
	@govulncheck ./...

.PHONY: vuln
vuln:
	make -j1 gosec govuln


# Live reload

.PHONY: tailwind
tailwind:
	@npx @tailwindcss/cli -i ./ui/assets/css/input.css -o ./ui/assets/css/app.css --minify --watch

.PHONY: templ
templ: check-templ
	@templ generate --watch --proxy="http://localhost:8080" --open-browser=false

.PHONY: server
server: check-air
	@air \
	--build.cmd "go build -o tmp/bin/shortlink" \
	--build.full_bin "GO_ENV=dev tmp/bin/shortlink" \
	--build.args_bin "serve" \
	--build.delay "100" \
	--build.exclude_dir "node_modules" \
	--build.include_ext "go" \
	--build.stop_on_error "false" \
	--misc.clean_on_exit false

.PHONY: dev
dev:
	@make -j3 tailwind templ server
