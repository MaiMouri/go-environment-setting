# 環境ごとに設定ファイルを読み込む  

godotenvとGetenvで環境ごとに設定ファイルを読み込むデモプロジェクトです。  
初期状態ではローカル環境とDocker環境で確認可能。

## 概要  

動作環境上の環境変数`GO_ENV`に環境名を入れることで実行時に環境ごとの設定ファイルを読み込むようになる。  
envfiles/`環境名`.env

## 使用技術、ライブラリ  

|  項目  |  用途  |
| ---- | ---- |
|  Go  |  言語  |
|  MySQL  |  DB  |
|  docker  |  コンテナ  |
|  docker-compose  | マルチコンテナ |  
|  gin  |  フレームワーク  |
|  gorm  |  ORM  |
|  godotenv  | 環境変数 |
|  gomod  |  モジュール管理  |

## 環境構築(Dockerの場合)  

[docker](https://www.docker.com/get-started)  
[docker-compose](https://docs.docker.com/compose/install/)  
がインストールされていること。

1.クローン  
> git clone https://github.com/s-moteki/go-playground.git

2.移動  
> cd go-playground/environment-setting-file

3.コンテナ起動  
> docker-compose up --build

## 環境構築(ローカル環境の場合)  

任意のDBがインストールされていること  

1.クローン  
> git clone https://github.com/s-moteki/go-playground.git

2.DB準備  
`go-playground/environment-setting-file/envfiles/local.env`をローカル環境のDBに合わせて修正する。  
その後DBに`go-playground/environment-setting-file/docker/db/initdb.d/create_tables.sql`を流す

3.移動  
> cd go-playground/environment-setting-file/app

4.起動
> go run server.go

## 動作確認  

一覧取得
> curl localhost:8080/users  
