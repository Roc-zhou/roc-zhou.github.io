### 如何用一行代码实现网页变灰效果？
文章 摘抄于 菜鸟教程 [如何用一行代码实现网页变灰效果？]([https://mp.weixin.qq.com/s?__biz=MzA5NDIzNzY1OQ==&mid=2735615902&idx=1&sn=b72bebf2d87f63f6ab281bbde9478fa9&chksm=b6ab292a81dca03c9818df8fbf9a906dd7c1daccbbcc6fd25fee76eafcdea883b5a7eed58cc2&mpshare=1&scene=1&srcid=&sharer_sharetime=1586025842561&sharer_shareid=0a540c481054f8aab2b11efcd32ae53d&key=784d6f143f23d481fd3c2636c4d23767fb903c50a158850ee8a226b173e544274d472cb86e82dc9cd3e0194ed8468306c0fa88fc6f200e7a2f52eaf26ddb9b4af56d926dcb18adc82c3d97c3a3b6446c&ascene=1&uin=MjE4OTg4MjAxMA%3D%3D&devicetype=Windows+10&version=62080079&lang=zh_CN&exportkey=A0s2G7YzCLtJ1ZhxUXHdMqI%3D&pass_ticket=AgvJs22QcdCExtRNimNaruoRQTW2jf9PoBOLRXC1ggKVlbhKlGQ9a22KoP5vtncT](https://mp.weixin.qq.com/s?__biz=MzA5NDIzNzY1OQ==&mid=2735615902&idx=1&sn=b72bebf2d87f63f6ab281bbde9478fa9&chksm=b6ab292a81dca03c9818df8fbf9a906dd7c1daccbbcc6fd25fee76eafcdea883b5a7eed58cc2&mpshare=1&scene=1&srcid=&sharer_sharetime=1586025842561&sharer_shareid=0a540c481054f8aab2b11efcd32ae53d&key=784d6f143f23d481fd3c2636c4d23767fb903c50a158850ee8a226b173e544274d472cb86e82dc9cd3e0194ed8468306c0fa88fc6f200e7a2f52eaf26ddb9b4af56d926dcb18adc82c3d97c3a3b6446c&ascene=1&uin=MjE4OTg4MjAxMA%3D%3D&devicetype=Windows+10&version=62080079&lang=zh_CN&exportkey=A0s2G7YzCLtJ1ZhxUXHdMqI%3D&pass_ticket=AgvJs22QcdCExtRNimNaruoRQTW2jf9PoBOLRXC1ggKVlbhKlGQ9a22KoP5vtncT)
)

### 网站变灰
```css
html {
    -webkit-filter: grayscale(100%);
    -moz-filter: grayscale(100%);
    -ms-filter: grayscale(100%);
    -o-filter: grayscale(100%);
    filter: grayscale(100%);
    filter: progid:DXImageTransform.Microsoft.BasicImage(grayscale=1);
}

// 兼容写法
.gray {
    -webkit-filter: grayscale(100%);
    -moz-filter: grayscale(100%);
    -ms-filter: grayscale(100%);
    -o-filter: grayscale(100%);
    filter: grayscale(100%);
    filter: progid:DXImageTransform.Microsoft.BasicImage(grayscale=1);
}
```
这个样式名叫做 filter，搜下 MDN 的官方介绍，其链接为：https://developer.mozilla.org/zh-CN/docs/Web/CSS/filter。
官方介绍内容如下：
```
filter CSS 属性将模糊或颜色偏移等图形效果应用于元素。滤镜通常用于调整图像，背景和边框的渲染。
CSS 标准里包含了一些已实现预定义效果的函数。你也可以参考一个 SVG 滤镜，通过一个 URL 链接到 SVG 滤镜元素 (SVG filter element[1])。
```
其实就是一个滤镜的意思。
其所有用法示例如下：
```css
/* URL to SVG filter */
filter: url("filters.svg#filter-id");

/* <filter-function> values */
filter: blur(5px);
filter: brightness(0.4);
filter: contrast(200%);
filter: drop-shadow(16px 16px 20px blue);
filter: grayscale(50%);
filter: hue-rotate(90deg);
filter: invert(75%);
filter: opacity(25%);
filter: saturate(30%);
filter: sepia(60%);

/* Multiple filters */
filter: contrast(175%) brightness(3%);

/* Global values */
filter: inherit;
filter: initial;
filter: unset;
```
各个用法介绍大家可以参考官方的文档说明：https://developer.mozilla.org/zh-CN/docs/Web/CSS/filter