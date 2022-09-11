# list
### 简介
使用`gin`+`gorm`编写的简单的待办表单
### 界面样式
![图片](https://user-images.githubusercontent.com/76742505/189521659-4d719f55-4d6d-4129-9d70-15df1e2df705.png)
### 食用指南
1. 首先运行`go mod tidy`来引用项目需要的依赖
2. 找到`Configs/config.ini`文件
填写一些必要的信息
```ini
port = 9090 ;default
release = false ;default

[mysql]
user = root  ;default
password = root  ;default
host = 127.0.0.1  ;default
port = 3306  ;default
database = test  ;default
```
3.运行`go mod`
运行之后会在项目里面生成一个和项目同名可执行文件
4. 运行`.\[name].exe Configs/config.ini`
然后到浏览器去访问
