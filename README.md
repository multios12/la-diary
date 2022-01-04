# la-diary

## initialize setting
> npm i -g yarn

## create new go project
> go mod init main
> go get github.com/gin-gonic/gin

## production build
> .\build.cmd

## production build for Linux
> .\build.cmd

## memo
* vsdcodeでサブディレクトリにgo.modを置く
  setting.jsonに下記を追加
  > "gopls": {
  >     "build.experimentalWorkspaceModule": true
  > }

## Project setup
```
yarn install
```

### Compiles and hot-reloads for development
```
yarn serve
```

### Compiles and minifies for production
```
yarn build
```

### Lints and fixes files
```
yarn lint
```

### Customize configuration
See [Configuration Reference](https://cli.vuejs.org/config/).
