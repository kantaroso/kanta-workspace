## 仕様

### トップページ
* アクセスカウンター
  * PV
* メニュー


## ディレクトリ構成

## goのモジュール関連の使い方

```sell
# モジュールについて
# https://qiita.com/tana6/items/df9a48eecb84576f618d

# gopath以下で作業していないので必ずpackage名を明記する
# https://castaneai.hatenablog.com/entry/2019/02/22/151213
cd app/gin
go mod init github.com/kantaroso/kanta-workspace

# go get
# 必要なライブラリは都度 go get
go get github.com/gin-gonic/gin

# ローカルモジュールの設定
# app/gin/go.mod に replaceを記載
# https://qiita.com/moto_pipedo/items/c5859e9c4ad81cdbe750
cd app/gin/controllers
go mod init github.com/kantaroso/kanta-workspace/controllers

# go 仮実行
cd app/gin
go run main.go

```
##


## ローカル構築手順

### php

* [docker-compose](infra/docker-compose/laravel/README.md)

```shell

# dokcerのイメージを作成する
make build-php-base

# 作成されたイメージをdocker-composeに設定する
# dockerを起動する
docker-compose -f infra/docker-compose/laravel/docker-compose.yml up -d

```

### go

* [docker-compose](infra/docker-compose/gin/README.md)

### vue

* [docker-compose](infra/docker-compose/vue_gin/README.md)


## データベース作成

* 準備
```sql
-- localhost:8080
create database kanta_workspace DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;
```

* laravelのファイル作成 [参考](https://qiita.com/shosho/items/a5a5839735dfef9214b1)、[参考2](https://readouble.com/laravel/5.7/ja/eloquent.html)

```shell
# ormは利用しない
# php artisan make:model Model/access_log -m

php artisan make:migration create_access_log_table --create=access_log
```

* マイグレーションファイル編集後にマイグレーションする
```shell
php artisan migrate
```

## k8s構築手順

* [参考](https://qiita.com/ocadaruma/items/efe720e46ae7ecb9ec25)

* [minikubeの使い方](https://github.com/kantaroso/kubernetes-training)

* [コマンドチートシート参考](https://qiita.com/suzukihi724/items/241f7241d297a2d4a55c)

```shell

eval $(minikube docker-env)

# php
make build-php-prod

# go
make build-go-prod

```

```shell

cd infra/k8s

# 全ての確認
kubectl get all
kubectl get pv
kubectl get pvc

# 各種登録
kubectl apply -f infra/k8s/namespace.yml

kubectl apply -f infra/k8s/go/pv.yml
kubectl apply -f infra/k8s/go/pvc.yml
kubectl apply -f infra/k8s/go/service.yml
kubectl apply -f infra/k8s/go/deployment.yml

kubectl apply -f infra/k8s/mysql/pv.yml
kubectl apply -f infra/k8s/mysql/pvc.yml
kubectl apply -f infra/k8s/mysql/service.yml
kubectl apply -f infra/k8s/mysql/deployment.yml

# アクセスURLの確認
minikube service mysql --url
minikube service go --url

# ssh で中身を確認
kubectl exec -it <pod-name> /bin/bash

```


## クエリ

```
create database kanta_workspace DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;

CREATE TABLE `access_log` (
  `id` bigint AUTO_INCREMENT NOT NULL,
  `method` varchar(255) NOT NULL,
  `endpoint` varchar(255) NOT NULL,
  `query_string` text NOT NULL,
  `user_agent` varchar(255) NOT NULL,
  `created_at` timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
  PRIMARY KEY (id)
)
```
