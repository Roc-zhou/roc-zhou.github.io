# 总结

### css盒模型

css盒模型可以分为 标准模型 和 IE模型，由 margin、border、padding、content 构成

![b9dc3737a4e74d15bcb62888f54c1f43_tplv-k3u1fbpfcp-zoom-1.png](https://i.loli.net/2020/09/19/ESGA7DcyjBCUJQf.png)

标准盒模型 和 IE 盒模型的区别：

主要区别就是 width 和 height 不同

标准模型的宽高是 content的宽高，IE盒模型的宽高 是 border + padding + content 的总尺寸

### JS部分

#### Event 对象常见应用有哪些

取消时间的默认动作： `event.preventDefault()`

阻止事件冒泡：`event.stopProagation()`

#### 什么是构造函数

构造函数：用于批量创建同类对象的函数

构造函数和普通函数的差异

- 构造函数首字母大写
- 构造函数this代表新对象用于初始化新对象的私有属性
- 并不是用return 返回函数的执行结果
- 需要用new来执行函数创建新对象

#### 原型（**prototype**）总结

首先 js分为 **函数对象** 和 **普通对象**，每个对象都有`__proto__`属性，但是只有函数对象才会有`prototype`属性

Object、Function 都是js内置的函数，类似的还有 Array、RegExp、Date、Boolean、Number、String

每一个实例都有一个 `__proto__` 属性，指向其构造函数的原型 （prototype）

```js
var s1 = new Bird
	s2 = new Bird
s1.__proto__ === Bird.prototype
s1.__proto__ === s2.__proto__
```

**`__proto__`和`prototype`的区别是什么？**

属性`__proto__`是一个对象，它有两个属性 `constructor` 和 `__proto__`

原型对象 prototype 有一个默认的 constructor 属性，用来记录实例是哪一个构造函数创建的（**constructor** 的作用！！！）

**继承**

继承分为 原型继承 和 非原型继承

原型继承：用来实现下级函数实例使用上级函数原型中的属性，也就是说 将上级函数的实例赋值给下级函数的原型

```js
function Person(name, age){
    this.name = name;
    this.age = age;
}
Person.prototype.say = function(){
    console.log("我叫"+ this.name);
};
function Student(name, age){
    this.name = name;
    this.age = age;
}
//原型继承
Student.prototype = new Person();
Student.prototype.study = function(){
    console.log(this.name + "爱学习");
}
var stu = new Student("tom","18");
console.log(stu.say()) // 我叫tom
console.log(stu.study()) // tom爱学习
```

非原型继承：用 call 和 apply 实现 非原型方式继承

call 和 apply 用于修改函数执行时内部 this 的指向。功能一致

call 和 apply 的区别 就是 参数的区别，第一个参数相同，而call 的第二个参数是已列表的形式展现，apply 是已数组的形式

```js
//非原型继承
function Person(name, age, sex){
    this.name = name;
    this.age = age;
    this.sex = sex;
}
function Student(name, age, sex, stuNum){
    // call 和 apply 的区别
    // Person.call(this,name, age, sex);
    Person.apply(this,arguments);
    // this.name = name;
	// this.age = age;
	// this.sex = sex;
    this.stuNum=stuNum;
}
```

对象底层的查找规则：当访问对象属性时，首先查找自身的私有属性，没有就沿着 `__proto__` 属性逐级向上查找原型中的属性。有则返回属性值，没有就返回 undefined，优先级逐渐下降

**原型链**

![image.png](https://i.loli.net/2020/09/19/KVeUOxtnXNFboay.png)

什么是原型链？

原型通过继承关系逐级连接组成的链系结构叫做原型链

看个例子

```js
function A(){}
function B(){}
function C(){}

A.prototype.fa = function(){}
B.prototype.fb = function(){}
C.prototype.fc = function(){}

// 继承
B.prototype = new A();
B.prototype.constructor = B;
C.prototype = new B();
C.prototype.constructor = C;


//创建c的实例
var c = new C();
console.log(c);

console.log(c.__proto__ === C.prototype);//true
console.log(c.__proto__.__proto__ === B.prototype);//true
console.log(c.__proto__.__proto__.__proto__ === A.prototype); //true
console.log(c.__proto__.__proto__.__proto__.__proto__ === Object.prototype); //true
console.log(c.__proto__.__proto__.__proto__.__proto__.__proto__); //Object.prototype.__proto__ === null;

```

#### this的典型场景

this 代表什么？

this 是函数执行时，函数内不创建的临时对象，代表当前函数对象

1、对象内部属性值是函数，通过对象调用函数时，this代表的当前对象

```js
var person = {
    name: "tom",
    say(){ 
        console.log(this.name)
    }
}
person.say(); // tom
```

2、Dom 对象事件函数执行时，this 代表当前Dom 对象

```js
document.body.onkydown = function(){
    console.log(this)
}
```

3、构造函数中，this代表新创建的对象

```js
function Person(name, age){
    this.name = name;
    this.age = age;
    a = this;
}
var a, b;
b = new Person("cat", 28);
console.log(b.name) // cat
console.log(b.age) // 28
console.log(b.age) // Person
```

4、call 和 apply 通过函数执行的主函数，内部this 是第一个实惨

```js
var person = {
    name: "tom",
    say(){ 
        console.log(this.name)
    }
}
function Person(name, age){
    this.name = name;
    this.age = age;
}
var b = new Person("cat", 28);
person.say.call(b); // cat
```

5、其他 this 代表 window

```js
function fn(){
    console.log(this);
}
fn();
```

#### 闭包

什么是闭包？简单来说就是函数返回内层函数，内层函数持有外层函数变量 就是闭包 (就是把变量封装在外层函数内部，保证长期有效)

```js
function fn(){
    var name = "xiaoming"
    return function fn2() {
        return name
    }
}
```

#### Proime 和 async、await



### Http 部分

#### 1、http 常见状态码

- 1xx：指示信息--表示请求已经接收，继续处理
- 2xx：成功--表示请求已经被成功接受、理解、接收
- 3xx：重定向--要完成请求必须进行更进一步的操作
- 4xx：客户端错误—请求有语法错误或者请求无法实现
- 5xx：服务端错误—服务器未能实现合法的请求

常见状态码：

- 200 ok：成功
- 400 Bad request：请求错误
- 401 Unauthorized：请求未经授权，这个状态码必须和WWW-Authenticate报头一起使用
- 403 Forbidden：服务器收到请求，拒绝服务
- 404 Not Found： 请求资源不存在，通常是api不存在 或者未部署
- 500 server error： 服务器错误

#### 2、浏览器输入了Url 之后发生了哪些事情

首先浏览器中输入了url 会首先进行 解析，判断你是一个合法的url，然后会去进行 dns 解析，建立 tcp 连接，然后处理请求、响应请求，最后渲染到页面

![image.png](https://i.loli.net/2020/09/19/nsutFYoGQVTZHmi.png)

### 安全部分

- CSRF 和 xss攻击

**CSRF（Cross-site request forgery）跨站请求伪造**

**简单原理：**就是浏览器打开网站 A 进行 操作，然后又打开了网站B（恶意网站），B网站收到用户请求，返回攻击代码得到信息，然后访问网站A。然后去发送请求

**解决方法：**

- Token验证
- Refer验证
- 自定义令牌

**XSS（cross-site scripting）跨域脚本攻击**

**简单原理：**在网站页面输入框内插入 恶意的代码

**解决方法：**

输入框对特殊字符做校验，或者禁止输入特殊字符 或者 进行字符转译等等