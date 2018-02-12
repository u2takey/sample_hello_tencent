# sample_hello_tencent

## 1. 初始化

下载go语言, 设置好GOPATH
使用包管理工具[Dep](https://github.com/golang/dep)

```
dep init
```

创建dockerfile, makefile

```
touch dockerfile
touch Makefile
```

## 2. 编写代码
hello.go

## 3. 编写部署配置文件
deploy.yaml

## 4. 打包镜像，部署
```
➜  sample_hello_tencent git:(master) ✗ make push_docker
➜  sample_hello_tencent git:(master) ✗ make deploy_ccs
kubectl apply -f ./deploy.yaml
deployment "hello" created
service "demo" created
```

## 5. ccs设置暴露端口，绑定域名

访问 : http://118.25.33.62/hello
或者: http://hello.u2takey.tech/hello