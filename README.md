# micro-falsework


创建Docker镜像 

```
docker build -t micro-falsework:latest .
```

再需要创建的目录生成项目

```
docker run --rm -v $(pwd):$(pwd) -w $(pwd) micro-falsework [项目名称]
```