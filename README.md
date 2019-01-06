# Moly
这是一个proxy，支持 TCP 和 HTTP 协议。

## 转发规则
1. PathPrefix        匹配前缀
2. PathPrefixStrip   匹配前缀，并删除该前缀
3. Path              完全匹配 path
4. PathRegex         完全匹配 path，正则匹配
5. Method            方法匹配，GET/POST etc.
6. Host              host 字段匹配