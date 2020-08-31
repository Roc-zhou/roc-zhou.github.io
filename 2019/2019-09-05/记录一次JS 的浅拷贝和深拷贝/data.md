简单理解浅拷贝和深拷贝
> 博客地址：[https://www.zhouhaipeng.com/](https://www.zhouhaipeng.com/)
> github [https://github.com/Roc-zhou](https://github.com/Roc-zhou)
### 什么是浅拷贝？
简单来说 有个 A 和 B，我们把A 复制给了B，当我们修改B 的时候 发现 A 的值也会发生改变，这就是浅拷贝
- 举个例子
```js
let a = [{ name: '张三' }],
b = a

b[0].name = '李四'
console.log(a);  // [{ name: '李四' }]
console.log(a); //  [{ name: '李四' }]
```
- 使用图解释一下
![9419407-1d09c5b4a0a06252.png](https://i.loli.net/2019/09/05/Rx53MBlsjQvczyp.png)
![9419407-ff2d0c88f14929a3.png](https://i.loli.net/2019/09/05/3fYTHq4dxDh9VaS.png)


从图中我们可以看到 a和b 全部 都指向了 同一个 堆内存的值 ，所以我们改变 b 的数据 a也会 改变 。

### 那什么是深拷贝呢？
理解了浅拷贝，我猜 深拷贝就是修改 b 的值 a 还是原来的值。
- 还是举个例子先
```js
// 深拷贝的方法有几种 这里只写一种 （JSON.stringify、jquery等）
function clone(obj) {
  let result = obj.constructor === Array ? [] : {};
  if (typeof obj !== 'object') {
    return obj;
  }
  for (const x in obj) {
    if (obj.hasOwnProperty(x)) {
      if (typeof obj[x] === 'object' && obj[x] !== null) {
        result[x] = clone(obj[x]);   //递归复制
      } else {
        result[x] = obj[x];
      }
    }
  }
  return result
}

let a = [{ name: '张三' }],
  b = clone(a)

console.log(a); // [ { name: '张三' } ]
console.log(b); // [ { name: '张三' } ]
// 成功复制
// 修改b的值
b[0].name = '李四'
console.log(a); // [ { name: '张三' } ]
console.log(b); // [ { name: '李四' } ]
```
- 还是上个图
![9419407-efe13a61709541e3.png](https://i.loli.net/2019/09/05/LJOBH9le5adCWQA.png)
![9419407-47e577f9b52cb1c7.png](https://i.loli.net/2019/09/05/UMPNlOYEkXtdF5A.png)

从图上我们可以看到 b 复制 a 是在 堆内存中复制了一个生成了一个新的，所以我们修改b 的值的时候 a 不会受影响，这样我们就得到了一个新的b 达到了深拷贝的效果！