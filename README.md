## 仕様

### トップページ
* アクセスカウンター
  * PV
* メニュー


## ディレクトリ構成


## ローカル構築手順
```shell

# dokcerのイメージを作成する
infra/docker/make build-php-base

# 作成されたイメージをdocker-composeに設定する
# dockerを起動する
docker-compose -f infra/docker-compose/laravel/docker-compose.yml up -d

# イメージに問題なければイメージをdocker.hubにアップする
docker push kantaroso/php_base:xxxxxxxxx

```



## データベース作成

* 準備
```sql
-- localhost:8080
create database sample DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;
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

```shell

eval $(minikube docker-env)

cd docker/php

# docker imageを作成する
make build

```

```shell

cd k8s

# 全ての確認
kubectl get all

# 各種登録
kubectl apply -f namespace.yml
kubectl apply -f php/deployment.yml
kubectl apply -f php/service.yml

# アクセスURLの確認
minikube service php --url

# ssh で中身を確認
kubectl exec -it <pod-name> /bin/bash

```
