# grpc_demo

```
                             |                    
   _` |   _| _ \   _|     _` |   -_)   ` \    _ \ 
 \__, | _|  .__/ \__|   \__,_| \___| _|_|_| \___/ 
 ____/     _|                                     
```

### 证书生成
``` shell script
$ cd certs/
$ ./gen_certs.sh
```

### ProtoBuf 文件清理与生成
```shell script
$ make
Choose a command run in grpc_demo:

Usage: make [target]

Valid target values are:

proto	generate proto file
clean	clean proto file
help	print this help message and exit.
```
