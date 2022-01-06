## 2.1 select

### 3.小结

- select 仅能操作管道；
- 每个case语句仅能处理一个管道，要么读要么写
- 多个case语句的执行顺序是随机的
- 存在default语句，select将不会阻塞