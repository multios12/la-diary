# la-diary
  this app is line diary.

## デバッグ実行
vscode上での実行を前提。chromeを利用
フロントデバッグサーバ：ポート3001
バックエンドサーバ：ポート3000

### 実行手順
1. 実行とデバッグで「go API Server」を選択、実行
2. 実行とデバッグで「debug react」を選択、実行

-------------------------------------------------------------

## initialize setting
> npm i -g yarn

## create new react project
> create-react-app
> yarn add react-router-dom

## create new go project
> mkdir srv
> cd srv
> go mod init main
> go get github.com/gin-gonic/gin
