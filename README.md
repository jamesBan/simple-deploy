# gogs简单代码发布系统

> 根据git钩子触发远程接口进行发布



![](https://raw.githubusercontent.com/o2team/misc/gh-pages/pfan123/githook/githook_3.png)




## 1.server端

### 配置server下的config.yaml 文件

```yaml
---
server: #hook对应的服务器配置
  host: 127.0.0.1 #服务器ip/域名
  port: 8080 # 端口
  secret: secret_token #密码

release_server: #项目列表
  project1: #项目1
  - 127.0.0.1:9530 #对应的agent地址
  project2:
  - 127.0.0.1:9530
```
### 运行build 将编译后的文件部署到服务器上





## 2.agent端

### 配置agent下的config.yaml文件
```yaml
project: #项目列表
- name: project1 #项目名
  worker_dir: D:\www\project1 #项目目录
  exec_command: release.bat #发布项目时部署代码
  timeout: 20 #超时时间
- name: project2
  worker_dir: D:\www\project2
  exec_command: release.bat
  timeout: 20

server: #对应server配置信息
  host: 127.0.0.1
  port: 9530
  allow_ips: #允许请求的ip
  - 127.0.0.1
```

### 运行build 将编译后的文件部署到相应的服务器上


## 3.配置gogsWeb 钩子
仓库管理/添加 Web 钩子--->推送地址(server配置地址)---->密钥文本(secret)--->选择指定事件（发布版本）