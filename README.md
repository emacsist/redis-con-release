# 说明

释放掉空闲时间 > idle 的 redis 连接.

# 启动

`app-macos` 这个是 Mac 的二进制格式

`app-linux-amd64` 这个是linux64 位的二进制格式.

`./应用名` 即可

`app.json` 这个要与二进制文件放在同一个目录.

# redis 的设置

默认情况下, redis 的 timeout 设置是0, 即永远不会主动释放的. 如果 redis server 没有设置的话, 就会导致 connection 一直会存在. 这时, 积累下来, 容易使用 client 端连接时会报 connection timeout 的问题.

