# 基于kratos框架测试nacos服务发现

## 项目结构
* 项目总共2个服务
  * gateway: 提供http接口
  * user: 用户grpc服务

## 依赖关系
* user服务提供了2个service
  * User: 用户信息相关接口
  * Auth: 登录认证相关接口

* gateway服务内调用user提供的两个service
  * 查找用户接口
  
  ```shell
  curl -X POST "http://localhost:7001/user/find" -H "Content-Type: application/json" --data-raw "{
    \"id\":\"123456\"}"
  ```
  * 登录接口
  
  ```shell
  curl -X POST "http://localhost:7001/auth/login" -H "Content-Type: application/json" --data-raw "{
    \"id\":\"123456\"}"
  ```


## 遇到的问题
* 如果先启动gateway，然后再启动user，功能正常
* 如果交换启动顺序，会导致gateway服务中的 grpc ClientConn watcher卡住
  * 进而导致上面的两个接口只有其中一个能够成调用