### 1、利用ES6方法 set

```js
var arr = [1,2,3,3,3,3,4,4,5,6,7]
function unique(arr) {
	return Array.from(new Set(arr))
}
console.log(unique(arr)); // => [ 1, 2, 3, 4, 5, 6, 7 ]
```

### 2、for循环 splice去重 去重

```js
var arr = [1,2,3,3,3,3,4,4,5,6,7]
function unique(arr) {
	for (var i = 0; i < arr.length; i++) {
		for (var j = i+1; j < arr.length; j++) {
			if (arr[i] === arr[j]) {
				arr.splice(j,1);
				j--
			}
		}
	}
	return arr
}
console.log(unique(arr)); // => [ 1, 2, 3, 4, 5, 6, 7 ]
```

### 3、利用indexOf去重

```js
var arr = [1,2,3,3,3,3,4,4,5,6,7]
function unique(arr) {
	const array = []
	for(const x of arr) {
		if (array.indexOf(x) === -1) {
			array.push(x)
		}
	}
	return array
}
console.log(unique(arr)); // => [ 1, 2, 3, 4, 5, 6, 7 ]
```

### 4、...new Set(arr)

```js
var arr = [1,2,3,3,3,3,4,4,5,6,7]
console.log([...new Set(arr)]);// => [ 1, 2, 3, 4, 5, 6, 7 ]
```

### 5、利用sort()方法

```js
var arr = [1, 2, 3, 3, 3, 3, 4, 4, 5, 6, 7]

function unique(arr) {
	arr = arr.sort()
	var array = [arr[0]];
	for (var i = 1; i < arr.length; i++) {
		if (arr[i] !== arr[i-1]) {
			array.push(arr[i])
		}
	}
	return array
}
console.log(unique(arr)); // => [ 1, 2, 3, 4, 5, 6, 7 ]
```

