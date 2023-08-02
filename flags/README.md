<!--
 * @Author: facsert
 * @Date: 2023-08-01 22:33:00
 * @LastEditors: Please set LastEditors
 * @LastEditTime: 2023-08-02 22:40:43
 * @Description: 
-->

# Flag

```go
var paramList []map[string]string                // 自定义命令参数列表
func Main() (paramMap map[string]string)         // 解析列表, 获取命令参数, 返回参数与值的 map
```

自定义参数列表如下

```go
var paramList = []map[string]string{
    {
        "name": "host",
        "default": "127.0.0.1",
        "help": "server host",
    },
    {
        "name": "username",
        "default": "root",
        "help": "server username",
    },
    {
        "name": "password",
        "default": "admin",
        "help": "server password",
    },
}
                     
```

在`main.go` 添加打印 `fmt.Println(flags.Main())`  
命令行执行并设置参数, 命令行参数覆盖了默认值  

```bash
 $ go run main.go -username user -password root
 > map[host:127.0.0.1 password:root username:user]
```

