![cover](https://user-images.githubusercontent.com/33691840/122493790-dd921e80-d01a-11eb-994a-446274d8464a.jpg)

### 安装node
[https://nodejs.org/en/](https://nodejs.org/en/)
### 安装 pm2-logrotate-ext，为什么不是安装 pm2-logrotate ？
>这里为什么不是安装 pm2-logrotate ,前人经验=>  只有当日志文件大小超过上限时都会把日志追加到最新的文件中，比如设置了日志文件名为app-out.log，日志文件大小为1M，那么在app-out.log文件没有超过1M时，即使第二天新建了日志文件app-out_yyyy-mm-dd_hh-mm-ss.log，但是依然会把日志追加到app-out.log中而非新建的日志文件中，更坑的是，在超过设定的大小后，它是把超出部分写到新的日志文件中，然后继续在app-out.log中追加- -，在官方文档中也没有发现对这个问题的说明。[文章](https://www.cnblogs.com/daner1257/p/10763888.html)

> 具体 是不是这个原因，也没有时间去验证，交给广大同学了去验证了。
### 什么是 pm2-logrotate-ext 
pm2-logrotate-ext 就是 在 pm2-logrotate的基础上 修改了pm2-logrotate 的不足。其他功能还是和pm2-logrotate 是一样的
[pm2-logrotate-ext](https://github.com/Lujo5/pm2-logrotate-ext)

### 安装 pm2-logrotate-ext 首先你要先安装 pm2 
```sh
pm2 install pm2-logrotate-ext
```
### 配制
```sh
max_size（默认为10M）：当文件大小大于此值时，它将旋转它（工作人员可能在实际超过限制后检查文件）。您可以在随后结束指定单位：10G，10M，10K
retain（默认为30文件日志）：此数字是一次保留的循环日志的数量，这意味着如果保留= 7，则最多将有7个循环日志和当前的日志。
compress（默认为false）：通过gzip为所有旋转的日志启用压缩
dateFormat（默认为YYYY-MM-DD_HH-mm-ss）：使用的数据格式名称为日志文件
rotateModule（默认为true）：像其他应用一样旋转pm2模块的日志
workerInterval（默认为30秒）：您可以控制工作人员检查日志大小的间隔（最小为1）
rotateInterval（默认为0 0 * * *每天午夜）：该Cron用于在执行时强制旋转。我们正在使用节点计划来计划cron，因此节点计划的所有有效cron 对该选项来说都是有效的cron。Cron风格：
TZ（默认为系统时间）：这是用于偏移已保存日志文件的标准tz数据库时区。例如Etc/GMT-1，带有小时日志的值，将在文件名称为GMT-1的14GMT时保存文件13。
forced（默认为true）：在上启用或禁用强制轮播rotateInterval，如果设置为false，则仅在max_size达到限制时才会发生日志文件轮换。
```
### 如何设置这些值
```sh
pm2 set pm2-logrotate-ext:<param> <value>
```

### 个人设置
```sh
pm2 set pm2-logrotate-ext:max_size 100M
pm2 set pm2-logrotate-ext:retain 7
pm2 set pm2-logrotate-ext:compress false
pm2 set pm2-logrotate-ext:dateFormat YYYY-MM-DD_HH-mm-ss
pm2 set pm2-logrotate-ext:workerInterval 30
pm2 set pm2-logrotate-ext:rotateInterval 0 0 * * *
pm2 set pm2-logrotate-ext:rotateModule true
pm2 set pm2-logrotate-ext:forced false

设置完成重启下 pm2 reloadLogs
```