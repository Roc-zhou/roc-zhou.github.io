## 1、下载需要合并的仓库A

```sh
git clone xxx.git
```

## 2、添加需要被合并的仓库 B

```sh
git remote add base xxx.git
```
> base 只是名字可以随意取

使用 `git remote` 查看所有远程仓库 将看到一个本地默认仓库 origin 和 另外一个我们新增的 base

## 3、把 base 远程仓库 B 中的数据加到当前 A 仓库中

```sh
git fetch base  // base 就是上面 add remote 的名字
```

## 4、基于 base 仓库的 master 分支（也可以是其他分支），新建一个分支并切换到该分支

```sh
git checkout -b masterB base/master  // base/xxx
```
`git branch` 查看所有分支

## 5、 切换需要合并的分支 master 或其他分支 然后进行 merge 操作

```sh
git checkout master

git merge masterB --allow-unrelated-histories
```
如果不加 --allow-unrelated-histories 就会出现 `fatal: refusing to merge unrelated histories` 这个错误。


## 6、解决冲突

。。。。

## 7、提交

push

