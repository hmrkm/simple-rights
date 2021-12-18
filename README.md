# simple-rights

[![codecov](https://codecov.io/gh/hmrkm/simple-rights/branch/main/graph/badge.svg?token=LE4923URW1)](https://codecov.io/gh/hmrkm/simple-rights)


シンプルな権限管理


## 必要なもの

- Docker Compose

## インストール

1. `.env.sample`をコピーして`.env`を作成
2. `.env`の内容を修正
3. `docker-compose up -d`
4. DBに`app/docs/migration.sql`の内容を反映

## 使い方

1. `/v1/rights`で権限チェック
