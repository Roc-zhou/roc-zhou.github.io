### 在使用git 中过程中不免出现及其卡顿的情况，然而及其卡顿我们的处理一般的处理办法就是杀掉进程等等。。但是这样就影响了git程序。
问题：
Another git process seems to be running in this repository, e.g.an editor opened by 'git commit'. Please make sure all processes are terminated then try again. If it still fails, a git process may have crashed in this repository earlier:
remove the file manually to continue.
翻译：
另一个git进程似乎正在这个存储库中运行，例如。一个由“git提交”打开的编辑器。请确认所有流程终止，然后重试。如果仍然失败，则使用git进程可能在此存储库早前崩溃:
**手动删除文件以继续。**

处理办法：
**打开 .git 文件夹 删除 index.lock 这个文件即可**
```sh
rm -rf ./.git/index.lock
```