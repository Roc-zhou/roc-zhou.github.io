### 1、获取 shh 
- 在根目录 ~ 下执行 ssh-keygen（直接回车，是否覆盖之前的key yes）然后进入刚才所在目（windows 用户 一般 在 c盘用户目录下.ssh文件夹下，Mac 用户 在 ~/.ssh）
- 查看 id_rsa.pub 公钥 复制内容 cat id_rsa.pub
- 到GitHub Or GitLab 上 点击设置 然后 选择 SSH和GPG 选项 然后new SSH key
- 填写信息， 粘贴刚才复制的内容就ok了！；
>  不熟悉命令的童鞋 推荐一款软件  [*sourcetree*](https://www.google.com/search?newwindow=1&safe=strict&q=sourcetree&spell=1&sa=X&ved=0ahUKEwj0xIL9gZHjAhWCZs0KHcPXBHkQkeECCCwoAA) 



### 2、第一次使用git 需要告诉你是谁？

```sh
配置用户名：git config --global user.name '你的用户名(英文)'
配置邮箱：git config --global user.email'你的邮箱'
```
```sh
pwd  查看当前目录
git init  初始化一个仓库
git status  查看仓库的状态

git add   操作会把工作区的求该添加到暂存区
git add * 或者 git add ./   把当前目录所有的提交
git commit -m '修改描述'    把暂存区的所有修改内容一次性提交到mastor分支上

git log    查看修改日志
git log --pretty=oneline   查看日志 一行模式
git reflog   查看版本id前7位  并且可以查看只要commit的内容
```

**代码版本回退和切换**

```sh
git reset --hard '提交id(版本号)' 
​git reset --hard head^ 上一个本
git reset --hard head^^ 上两个版本
git reset --hard head-n
```

**撤销**:撤销分两种情况

```sh
git checkout 文件名   撤销文件  没有add提交到暂存区的时候
​已经add到暂存区的时候撤销 
git reset HEAD 文件名
```
**分支相关**
```sh
git branch  查看当前分支列表
git checkout -b 分支名字  创建并切换分支（把分支上的所有记录全部都拷贝过去）
git checkout 分支名字  切换分支
​git branch -d 分支名字  删除分支
git merge 分支名字  合并	某分支到当前分支	
```

**克隆远程仓库：**`git clone 仓库地址（HTTPS和SSH）协议都可以（ssh协议必须在github或者gitlab中 添加了你的公钥），克隆好之后，就会在本地生成一个该项目的仓库，不需要在手动去执行git init操作。 `
> **注意：克隆不需要克隆多次，一般只需要克隆一次（前提你的代码不丢失）
**更新远程仓库的代码**:`git pull origin 分支名字`。

**推送代码到远程仓库**：git push origin 分支名字, （在推送之前尽量要git pull 一下。要保证远端代码和本地代码保持同步。）

**本地文件和github仓库关联步骤**
---
```sh
git init  初始化仓库
git remote add origin ssh地址（仓库）
git pull origin master

// 新建文件等操作之后
git add *.文件
git commit -m '描述'
git push origin master
```
#### Git 自定义别名！！
```sh
git config --global alias.st status >> git st
git config --global alias.a add .  >> git a
```
有兴趣的童鞋可以自己去研究哈！在这里就不多写别的命令了，自己熟悉的命令才是敲的最爽的


