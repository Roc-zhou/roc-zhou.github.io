问题：
同层目录下 一个 position 元素，一个 用了css 伪类，实现一个 三角 旋转的 效果 ，结果 position 层级 没有把底层的元素 彻底覆盖，原因是 写了一个 transform！！！！
![image.png](https://upload-images.jianshu.io/upload_images/9419407-406bf73911205bab.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)
一个很小的问题，浪费了时间，实属不该。所以写css 的时候能使用图片的还是 建议使用图片。
下面是 冲突部分代码

```css
<div>position 元素</div>
<p class="down">position 元素</p>

.down::after {
  content: "";
  width: 18px;
  height: 18px;
  margin-right: 15px;
  background-size: 100% 100%;
  background-image: url("../Images/menu_down.png"); // 一个向下的小箭头
  /* transform: rotate(90deg); */
  /* opacity: 0.6; */
}

.top::after {
  background-image: url("../Images/menu_top.png"); // 旋转变成 向上的小尖头
  /* transform: rotate(180deg); */
}
```