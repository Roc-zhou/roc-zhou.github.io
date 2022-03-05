# fmt包的各种使用

## Print() 函数
使用其操作数的默认格式打印格式，并写入标准输出。
当两个操作数都不是字符串时，在操作数之间添加空格。
返回写入的字节数和遇到的任何写入错误。
```go
func Print(a ...interface{}) (n int, err error)
```
## Println() 函数
作用和Print函数一样，只是输出的时候会换行
```go
func Println(a ...interface{}) (n int, err error)
```
## Printf() 函数
Printf根据格式说明符格式化并写入标准输出。
返回写入的字节数和遇到的任何写入错误。
```go
func Printf(format string, a ...interface{}) (n int, err error)
```
## Fprint、Fprintln、Fprintf
作用和上面的三个函数一样，只是把转换的结果写入到w中
```go
func Fprint(w io.Writer, a ...interface{}) (n int, err error)
func Fprintln(w io.Writer, a ...interface{}) (n int, err error)
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
```
## Sprint、Sprintln、Sprintf
作用和上面的三个函数一样，不同的是把结果以字符串形式返回
```go
func Sprint(a ...interface{}) string
func Sprintln(a ...interface{}) string
func Sprintf(format string, a ...interface{}) string
```
## Errorf() 函数
同Sprintf 函数，结果返回error类型的字符串
```go
func Errorf(format string, a ...interface{}) error
```
