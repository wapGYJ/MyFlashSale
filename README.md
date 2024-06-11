# 使用go-zero框架的秒杀系统

## 项目大致框架
***
![Kiku](images/1.png)

## 客户端的完整请求路径
***
![Kiku](images/2.png)

### 客户端请求时请先将theorder,payment数据库里的数据清空。每次秒杀开启前，请先调用goods/rpc/internal/logic/goodscachelogic.go将所售商品缓存
