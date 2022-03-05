## GitFork项目如何与原作者项目保持一致！
```shell
git branch #查看下当前分支
git remote -v
git remote add upstream git@github.com:xxxx/xxxxx.git
git remote -v
git fetch upstream
git merge upstream/master
```

## 获取第几次的 commit 的内容
n 从 0 开始
```shell
git log -n 1 --skip n --pretty=format:"%s"  
```
