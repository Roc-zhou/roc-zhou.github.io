# 命令格式
```sh
scp [参数] [原路径] [目标路径]
```
 
# 命令参数
```sh
-1  强制scp命令使用协议ssh1 

-2  强制scp命令使用协议ssh2 

-4  强制scp命令只使用IPv4寻址 

-6  强制scp命令只使用IPv6寻址 

-B  使用批处理模式（传输过程中不询问传输口令或短语） 

-C  允许压缩。（将-C标志传递给ssh，从而打开压缩功能） 

-p 保留原文件的修改时间，访问时间和访问权限。 

-q  不显示传输进度条。 

-r  递归复制整个目录。 

-v 详细方式显示输出。scp和ssh(1)会显示出整个过程的调试信息。这些信息用于调试连接，验证和配置问题。  

-c cipher  以cipher将数据传输进行加密，这个选项将直接传递给ssh。  

-F ssh_config  指定一个替代的ssh配置文件，此参数直接传递给ssh。 

-i identity_file  从指定文件中读取传输时使用的密钥文件，此参数直接传递给ssh。   

-l limit  限定用户所能使用的带宽，以Kbit/s为单位。    

-o ssh_option  如果习惯于使用ssh_config(5)中的参数传递方式，  

-P port  注意是大写的P, port是指定数据传输用到的端口号  

-S program  指定加密传输时所使用的程序。此程序必须能够理解ssh(1)的选项。
```
 
 
# 上传目录到服务器 
```sh
scp -r -P 22 local_dir username@servername:remote_dir

例如：scp -r -P 22 test root@192.168.0.101:/var/www/ 把当前目录下的test目录上传到服务器的/var/www/ 目录
```

### 上传本地文件到服务器
```sh
scp -P 22 /path/filename username@servername:/path/

例如 scp -P 22 /var/www/test.php root@192.168.0.101:/var/www/ 把本机/var/www/目录下的test.php文件上传到192.168.0.101这台服务器上的/var/www/目录中
```
### 从服务器上下载文件
下载文件我们经常使用wget，但是如果没有http服务，如何从服务器上下载文件呢？
```sh
scp username@servername:/path/filename /var/www/local_dir（本地目录）

例如scp root@192.168.0.101:/var/www/test.txt 把192.168.0.101上的/var/www/test.txt 的文件下载到/var/www/local_dir（本地目录）
```
### 从服务器下载整个目录
```sh
scp -r username@servername:/var/www/remote_dir/（远程目录）

/var/www/local_dir（本地目录）
例如:scp -r root@192.168.0.101:/var/www/test /var/www/
```


### SSH连接服务器
```sh
ssh root@192.168.0.11 -p12333(端口号) 默认端口号为 22

ssh root@192.168.1.173  -> 输入password
```