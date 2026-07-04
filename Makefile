.PHONY: help test test-one fmt vet tidy build up down shell clean

# 例: make test-one P=3945
P ?=

# Docker コンテナ内で go を実行する (--rm で使い捨てコンテナを起動)
GO = docker compose run --rm app go

help:
	@echo "Usage:"
	@echo "  make test          全ての問題のテストを実行"
	@echo "  make test-one P=X  問題 X のテストのみ実行 (例: make test-one P=3945)"
	@echo "  make fmt           go fmt を全体に適用"
	@echo "  make vet           go vet を全体に実行"
	@echo "  make tidy          go mod tidy を実行"
	@echo "  make build         Docker イメージをビルド"
	@echo "  make up            Docker コンテナを起動 (バックグラウンド)"
	@echo "  make down          Docker コンテナを停止"
	@echo "  make shell         Docker コンテナ内でシェルを開く"
	@echo "  make clean         テストキャッシュを削除"

test:
	$(GO) test ./...

test-one:
	@if [ -z "$(P)" ]; then \
		echo "問題番号を指定してください。例: make test-one P=3945"; \
		exit 1; \
	fi
	$(GO) test -v ./problem-$(P)/...

fmt:
	$(GO) fmt ./...

vet:
	$(GO) vet ./...

tidy:
	$(GO) mod tidy

build:
	docker compose build

up:
	docker compose up -d

down:
	docker compose down

shell:
	docker compose run --rm app sh

clean:
	$(GO) clean -testcache
