## 生成go文件
```shell script
protoc --go_out=. hi.proto
```

## 生成js文件
```shell script
protoc --js_out=import_style=commonjs,binary:. hi.proto
```

## 打包为web可用的js文件
前置条件：需要安装npm。npm一般在安装nodejs的时候就会自动安装。

1.安装库文件的引用库
```shell script
npm install -g require
```
2.安装打包成前端使用的js文件
```shell script
npm install -g browserify
```
3.安装protobuf的库文件
```shell script
npm install google-protobuf
```
4.打包js文件export.js
```js
var address = require('./address_pb');
module.exports = {
    DataProto: address
}
```
5.编译生成可用js文件
```shell script
browserify exports.js -o  address_main.js
```