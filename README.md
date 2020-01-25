## 仕様

### トップページ
* アクセスカウンター
  * アクセス数
  * 訪問人数
* メニュー


## ディレクトリ構成



## 構築手順

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


