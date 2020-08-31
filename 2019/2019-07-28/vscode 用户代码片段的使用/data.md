> 文件 > 首选项 > 用户代码片段  选择 新建代码片段
### 快速输出 console.log（）
![GIF.gif](https://b2.bmp.ovh/imgs/2019/07/67717d64c25be1e9.gif)

```sh
"Print to console": {
    "scope": "javascript,typescript",
    "prefix": "log",
    "body": [
        "console.log('$1');",
    ],
    "description": "console.log()"
}
```
### 新建 *.vue 文件 自动生成 模板
![GIF2.gif](https://b2.bmp.ovh/imgs/2019/07/2c44bf888deeffde.gif)
```sh
    "Print to console": {
		"scope": "vue",
		"prefix": "vue",
		"body": [
			"<template>",
			"  <div class='app'>$0</div>",
			"</template>",
			"",
			"<script>",
			"export default {",
      "  beforeRouteEnter(to, from, next) {",
      "    return next(vm => {});",
      "  },",
      "  name: 'app',",
      "  data() {",
      "    return {};",
      "  },",
			"}",
			"</script>",
      "<style scoped>",
      "",
			"</style>"
		],
		"description": "create vue"
	}
```