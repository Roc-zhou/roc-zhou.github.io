### 什么是MVVM？

MVVM是Model-View-ViewModel缩写，也就是把MVC中的Controller变成了ViewModel

Model 代表数据模型、View代表Ui组件、ViewModel是 他们两个的桥梁，数据会绑定到ViewModel层 然后将数据渲染到页面中，试图变化会通知ViewModel层更新数据

### vue2 的响应式数据原理？

简单来说就是 把一个普通的 js对象传入 vue实例 作为 data。这时Vue 会遍历这个对象中所有property，使用 `Object.defineProperty`把这些property 全部转为 getter/setter。然后对其监听，在属性被访问和修改时通知，更新视图。

每一个组件实例都对应一个 watcher 实例，当 setter 触发时会通知watcher，然后重新渲染

![image-20200923203044435.png](https://upload-images.jianshu.io/upload_images/9419407-c61428b8f029e9a9.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

### vue2的生命周期

创建前/后、挂载前/后、更新前/后、销毁前/后

beforeCreate/created、beforeMount/mounted、beforeUpdate/updated、beforeDestroy/destroyed

beforeCreate 阶段是不能访问 data、methods、computed 和 watch 上的 数据的 ！

![image-20200923214641670.png](https://upload-images.jianshu.io/upload_images/9419407-f9b0b9416a23ab80.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

### computed 和 watch 的区别

computed 是一个具有缓存的 watcher ，当依赖的属性发生变化就会更新视图，适用于比较消耗性能的场景（例如：购物车场景）

watch 没有缓存的特性。数据变，就会触发响应的操作，支持异步。如果需要深度监听对象的属性时，可以使用 `deep：true` 

### 组件间的几种常用传值 操作

```vue
props、$emit、$parent、$ref、eventbus、vuex
```

### v-if 和 v-show 的区别

v-if  当条件不成立时 时不会渲染 dom 元素的

v-show  操作的是 display ，是变化 dmo 的 显示 和 隐藏 （回流！！！）

display 和 visibility



















