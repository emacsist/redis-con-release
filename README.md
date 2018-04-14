# 说明

释放掉空闲时间 > idle 的 redis 连接.

# 启动

`app-macos` 这个是 Mac 的二进制格式

`app-linux-amd64` 这个是linux64 位的二进制格式.

`./应用名` 即可

`app.json` 这个要与二进制文件放在同一个目录.

# redis 的设置

默认情况下, redis 的 timeout 设置是0, 即永远不会主动释放的. 如果 redis server 没有设置的话, 就会导致 connection 一直会存在. 这时, 积累下来, 容易使用 client 端连接时会报 connection timeout 的问题.

当然, 也可以设置 timeout 的值, 但如果设置得不适当, 很容易出现以下异常


```bash
Unexpected error occurred in scheduled task.org.springframework.data.redis.RedisConnectionFailureException: Unexpected end of stream.; nested exception is redis.clients.jedis.exceptions.JedisConnectionException: Unexpected end of stream.
	at org.springframework.data.redis.connection.jedis.JedisExceptionConverter.convert(JedisExceptionConverter.java:67)
	at org.springframework.data.redis.connection.jedis.JedisExceptionConverter.convert(JedisExceptionConverter.java:41)
	at org.springframework.data.redis.PassThroughExceptionTranslationStrategy.translate(PassThroughExceptionTranslationStrategy.java:37)
	at org.springframework.data.redis.FallbackExceptionTranslationStrategy.translate(FallbackExceptionTranslationStrategy.java:37)
	at org.springframework.data.redis.connection.jedis.JedisConnection.convertJedisAccessException(JedisConnection.java:212)
	at org.springframework.data.redis.connection.jedis.JedisConnection.evalSha(JedisConnection.java:3173)
	at org.springframework.data.redis.connection.jedis.JedisConnection.evalSha(JedisConnection.java:3158)
	at sun.reflect.GeneratedMethodAccessor103.invoke(Unknown Source)
	at sun.reflect.DelegatingMethodAccessorImpl.invoke(DelegatingMethodAccessorImpl.java:43)
	at java.lang.reflect.Method.invoke(Method.java:601)
	at org.springframework.data.redis.core.CloseSuppressingInvocationHandler.invoke(CloseSuppressingInvocationHandler.java:57)
	at $Proxy95.evalSha(Unknown Source)
	at org.springframework.data.redis.core.script.DefaultScriptExecutor.eval(DefaultScriptExecutor.java:81)
	at org.springframework.data.redis.core.script.DefaultScriptExecutor$1.doInRedis(DefaultScriptExecutor.java:71)
	at org.springframework.data.redis.core.RedisTemplate.execute(RedisTemplate.java:202)
	at org.springframework.data.redis.core.RedisTemplate.execute(RedisTemplate.java:164)
	at org.springframework.data.redis.core.RedisTemplate.execute(RedisTemplate.java:152)
	at org.springframework.data.redis.core.script.DefaultScriptExecutor.execute(DefaultScriptExecutor.java:60)
	at org.springframework.data.redis.core.script.DefaultScriptExecutor.execute(DefaultScriptExecutor.java:54)
	at org.springframework.data.redis.core.RedisTemplate.execute(RedisTemplate.java:298)
```
