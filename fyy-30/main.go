package main

// 写入文件 , os.OpenFile(name string, flag int, perm FileMode) (*File,error){...}

// flag 的选项如下
// os.O_WRONLY 只写
// os.O_CREATE 创建文件
// os.O_RDONLY 只读
// os.O_RDWR  读写
// os.O_TRUNC  清空
// os.O_APPEND  追加

// perm： 文件权限， 一个八进制数， r,w,x 分别对应为 读，写，执行
