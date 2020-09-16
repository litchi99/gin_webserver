# gin_webserver
Golang  RESTful  webserver

## 1. viper 
- 使用viper反序列化,将config.yaml配置读取到对应的构造体当中
- 将yaml配置文件反序列化到`v.Unmarshal(&global.GVA_CONFIG)`构造体当中`type Server struct`

## 2. gorm对象关系映射框架
- 数据库表名是结构体struct名字的蛇形小写、复数形式
- 数据库列名是struct内的字段名的蛇形小写