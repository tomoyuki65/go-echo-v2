-- テストユーザの作成
CREATE USER testuser;
ALTER USER testuser WITH PASSWORD 'test-password';

-- テストDBの作成
CREATE DATABASE testdb;

-- Atlas用のDB作成
CREATE DATABASE tmpdb;

-- テストユーザーにDBの接続権限付与
GRANT CONNECT ON DATABASE testdb TO testuser;
GRANT CONNECT ON DATABASE tmpdb TO testuser;

-- -- テストDBの権限付与
\c testdb
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO testuser;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO testuser;
GRANT USAGE ON SCHEMA public TO testuser;
GRANT CREATE ON SCHEMA public TO testuser;

-- -- Atlas用のDBの権限付与
\c tmpdb
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO testuser;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO testuser;
GRANT USAGE ON SCHEMA public TO testuser;
GRANT CREATE ON SCHEMA public TO testuser;
