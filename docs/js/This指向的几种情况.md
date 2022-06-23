

### 全局

无论是否在严格模式下，在全局执行环境中（在任何函数体外部）this 都指向全局对象 window

```js
// 在浏览器中, window 对象同时也是全局对象：
console.log(this === window); // true
a = 37;
console.log(window.a); // 37

this.b = "MDN";
console.log(window.b)  // "MDN"
console.log(b)         // "MDN"
```

### 函数

在函数内部，this 的取值取决于函数被调用的方式

- 严格模式

  ```js
  function f2(){
    "use strict"; // 这里是严格模式
    return this;
  }
  
  f2() === undefined; // true
  ```

  严格模式下，如果在执行的时候没有设置 this 的值，此时 this 会是 undefined

- 普通模式

  ```js
  function f1(){
    return this;
  }
  //在浏览器中：
  f1() === window;   //在浏览器中，全局对象是window
  
  //在Node中：
  f1() === globalThis;
  ```

  代码不在严格模式下，且 this 的值不是由该调用者设置的，所以 this 的默认指向是全局对象，也就是浏览器的 window

### 类中的 this

this 在类中的表现于函数中类似，因为类本质上也是函数，只是有些区别

在类的构造函数中，this 是一个常规对象。类中所有费静态的方法都会被添加到this 的原型中

```js
class Car {
  constructor() {
    // Bind sayBye but not sayHi to show the difference
    this.sayBye = this.sayBye.bind(this);
  }
  sayHi() {
    console.log(`Hello from ${this.name}`);
  }
  sayBye() {
    console.log(`Bye from ${this.name}`);
  }
  get name() {
    return 'Ferrari';
  }
}

class Bird {
  get name() {
    return 'Tweety';
  }
}

const car = new Car();
const bird = new Bird();

// The value of 'this' in methods depends on their caller
car.sayHi(); // Hello from Ferrari
bird.sayHi = car.sayHi;
bird.sayHi(); // Hello from Tweety

// For bound methods, 'this' doesn't depend on the caller
bird.sayBye = car.sayBye;
bird.sayBye();  // Bye from Ferrari
```

### call、apply、bind

- **`apply()`** 方法调用一个具有给定`this`值的函数，以及以一个数组（或[类数组对象](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Guide/Indexed_collections#working_with_array-like_objects)）的形式提供的参数。
- **`call()`** 方法使用一个指定的 `this` 值和单独给出的一个或多个参数来调用一个函数。
- **`bind()`** 方法创建一个新的函数，在 `bind()` 被调用时，这个新函数的 `this` 被指定为 `bind()` 的第一个参数，而其余参数将作为新函数的参数，供调用时使用。



在非严格模式下使用 call 和 apply 的时候回，如果用作 this 的值不是对象，则会被尝试转换为对象

null 和 undefined 被转换为全局对象。

```js
function add(c, d) {
  return this.a + this.b + c + d;
}

var o = {a: 1, b: 3};

// 第一个参数是用作“this”的对象
// 其余参数用作函数的参数
add.call(o, 5, 7); // 16

// 第一个参数是用作“this”的对象
// 第二个参数是一个数组，数组中的两个成员用作函数参数
add.apply(o, [10, 20]); // 34

function f(){
  return this.a;
}
```

es5引入了 Function.prototype.bind()。调用 f.bind(obj) 会创建一个与 f相同函数体和作用域的函数，但是在这个新函数中，this 永久的被绑定了 bind 的第一个参数，不需要关注函数是如何被调用的

```js
var g = f.bind({a:"azerty"});
console.log(g()); // azerty

var h = g.bind({a:'yoo'}); // bind只生效一次！
console.log(h()); // azerty

var o = {a:37, f:f, g:g, h:h};
console.log(o.a, o.f(), o.g(), o.h()); // 37, 37, azerty, azerty
```

### 箭头函数

在箭头函数中，this 与封闭词法环境的 this 保持一致（意思就是看在什么地方写的）。在全局代码中，它将被设置为全局对象：

```js
var globalObject = this;
var foo = (() => this);
console.log(foo() === globalObject); // true
```

### 构造函数

当一个函数用作构造函数的时候（使用 new 关键字），this 指向正在构造的新对象

> 虽然构造函数返回的默认值是 this 所指的那个对象，但是他仍然可以手动返回其他对象（如果返回值不是一个对象，则返回 this 对象）

```js
function C(){
  this.a = 37;
}

var o = new C();
console.log(o.a); // 37

function C2(){
  this.a = 37;
  return {a:38};
}

o = new C2();
console.log(o.a); // 3
```

### Dom 事件

当函数被用作事件处理函数时，它的 `this` 指向触发事件的元素（一些浏览器在使用非 `addEventListener` 的函数动态地添加监听函数时不遵守这个约定）。

```js
// 被调用时，将关联的元素变成蓝色
function bluify(e){
  console.log(this === e.currentTarget); // 总是 true

  // 当 currentTarget 和 target 是同一个对象时为 true
  console.log(this === e.target);
  this.style.backgroundColor = '#A5D9F3';
}

// 获取文档中的所有元素的列表
var elements = document.getElementsByTagName('*');

// 将bluify作为元素的点击监听函数，当元素被点击时，就会变成蓝色
for(var i=0 ; i<elements.length ; i++){
  elements[i].addEventListener('click', bluify, false);
}
```

未完...







