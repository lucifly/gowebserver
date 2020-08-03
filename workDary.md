# 2020-07-31
完成修改，从config.json中读取配置信息。  
Tips, 增加配置项需要在config.go中同步增加struct Config中的项。  

添加特性，添加命令行参数

# 2020-08-01
添加process处理的API接口   
    `/process/deploy` // 新增 process  
    `/process/delet` // 删除 process  
    `/process/start` // 开始 process  
    `/process/stop` // 结束 process  
    `/process/status` // 获取各个process的状态  

# 2020-08-03
设计了process执行的流程，以及信号的接入  
写了一个新的包 `processexe` 来负责process的五种接口的逻辑