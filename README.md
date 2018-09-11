# 说明
* golang：1.10.4
* gin：1.3  具体的文档，可以参考--https://github.com/gin-gonic/gin
* 依赖管理工具：govendor
* go get 命名的使用说明：http://wiki.jikexueyuan.com/project/go-command-tutorial/0.3.html


## conf：
* 该目录是保存应用的配置

## loggers：
* 该目录是保存应用的日志

## vendor：
* 该目录是下载依赖保存的地方


## 注意事项：
* 如果你发现下载第三方库的时候，被墙了，请参考此博文：https://www.jianshu.com/p/59b2b43850fd
* 你的项目必须放置在GOPATH的一级目录下

## todo：
1 日志记录
2 日志分析工具
3 性能监控--每个请求的耗时记录
4 Rating limit
5 Json Web Token
6 Etag
7 热替换
8 依赖管理工具使用