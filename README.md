##music的Go版本
####文件目录及作用
> 按照垂直拆分的原则，将项目按照职责的方式进行目录拆分


- api:目录存放对外开放的api接口
- cmd:每个模块的main启动文件,每个子目录包含一个模块
- configs:yml文件相关的配置
- doc:说明相关的文档
- internal:每个子目录是一个独立的模块，并且该文件内是本项目中内部
使用的相关文件，不对外开放使用