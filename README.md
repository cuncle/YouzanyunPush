# YouzanyunPush 

### 导入 

```
import (
	"github.com/cuncle/YouzanyunPush"
)
```

### 使用你的clientId和clientSecret 初始化

```
yclient:=YouzanyunPush.YouzanClient{"1","2"}
```

### 使用校验参考 

```
package main

import (
	"github.com/cuncle/YouzanyunPush"
	"net/http"
)

func main() {
	yclient:=YouzanyunPush.YouzanClient{"1","2"}
	http.HandleFunc("/", yclient.YouzanPush)
	http.ListenAndServe(":8888", nil)

}

```

### 检验测试curl 命令 

```
curl -X POST \
  http://127.0.0.1:8888/ \
  -H 'content-type: application/json' \
  -H 'event-sign: 4a53ae0b22d6ab013ec53ee291d688b' \
  -d '{"test":1}'

```
