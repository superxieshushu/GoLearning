
- Unicode vs UTF-8
Unicode 收集了世界上所有的字符、定长为 2 个字节。utf-8 编码是 Unicode 的一种实现。
UTF-8 是一种变长（1到4个字节）的编码方式。可以节省部分空间、兼容 ASCII 编码。

- rune 类型->int32的别名，几乎在所有方面等同于int32、它用来区分字符值和整数值