# 2020-07-31
完成修改，从config.json中读取配置信息。  
Tips, 增加配置项需要在config.go中同步增加struct Config中的项。  

添加特性，添加命令行参数

# 2020-008-01
添加process处理的API接口   
    `/process/deploy` // 新增 process  
    `/process/delet` // 删除 process  
    `/process/start` // 开始 process  
    `/process/stop` // 结束 process  
    `/process/status` // 获取各个process的状态  