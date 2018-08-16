src: 放源代码

如果有多个文件或多个包

1) 配置GOPATH环境变量， 配置src同级目录的绝对路径
	如：C:\Users\superman\Desktop\code\go\src\25_工程管理：不同目录
	
2) 自动生成bin或pkg目录，需要使用go install命令(了解)
	除了要配置GOPATH环境变量，还要配置GOBIN环境变量
	
src: 放源代码
bin: 放可执行程序
pkg: 放平台相关的库
