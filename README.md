# GoのEchoによるAPI開発のサンプルバージョン２
Go言語（Golang）のフレームワーク「Echo」によるバックエンドAPI開発用サンプルのバージョン２です。  
  
<br />
  
## 要件
・Goのバージョンは<span style="color:green">1.24.x</span>です。  
  
<br />
  
## ローカル開発環境構築
### 1. 環境変数ファイルをリネーム
```
cp ./src/.env.example ./src/.env
```  
  
### 2. コンテナのビルドと起動
```
docker compose build --no-cache
docker compose up -d
```  
  
### 3. コンテナの停止・削除
```
docker compose down
```  
  
<br />
  
## コード修正後に使うコマンド
ローカルサーバー起動中に以下のコマンドを実行可能です。  
  
### 1. go.modの修正
```
docker compose exec api go mod tidy
```  
  
### 2. フォーマット修正
```
docker compose exec api go fmt ./...
```  
  
### 3. コード解析チェック
```
docker compose exec api staticcheck ./...
```  
  
### 4. モック用ファイル作成（例）
```
docker compose exec api mockgen -source=./internal/repositories/XXX/XXX.go -destination=./internal/repositories/XXX/mock_XXX/mock_XXX.go
```  
  
### 5. テストコードの実行
```
docker compose exec api go test -v ./internal/handlers/...
```  
  
### 6. OpenAPIの仕様書修正
```
docker compose exec api swag i
```  
  
<br />
  
## マイグレーションに関する操作用コマンド
ローカルサーバー起動中に以下のコマンドを実行可能です。  
  
### 1. マイグレーションの状態確認
```
docker compose exec api go run ./database/cmd/migrate/status/main.go

docker compose exec -e ENV=testing api go run ./database/cmd/migrate/status/main.go
```  
> テスト用DBに対して実行したい場合はオプション「-e ENV=testing」を付ける  
  
### 2. マイグレーションの実行
```
docker compose exec api go run ./database/cmd/migrate/apply/main.go

docker compose exec -e ENV=testing api go run ./database/cmd/migrate/apply/main.go
```  
  
### 3. ロールバックの実行
```
docker compose exec api go run ./database/cmd/migrate/down/main.go

docker compose exec -e ENV=testing api go run ./database/cmd/migrate/down/main.go
```  
  
<br />
  
## API認証用のキーとパスワードを生成するためのコマンド
・APIキー認証を利用するには、事前に以下のコマンドでAPIキーとパスワードの生成が必要です。  
```
docker compose exec api go run ./cmd/create-apikey/main.go
```  
  
・APIキーは環境変数ファイル「.env」の「GO_ECHO_V2_API_KEY」に設定して下さい。  
```
GO_ECHO_V2_API_KEY=208c190373d51328cfda7b27993925bcc4c5edd0b50593f0a23cb730493f4711
```  
  