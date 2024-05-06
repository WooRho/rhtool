# rhtool

# 项目说明


# rhtool_connect_feie
 ## 飞鹅打印机对接
    提供增、删、打印等接口
    AddPrinter(snList string) (err error)
    Print(content string, sn string, times string) (orderId string, err error)
    Delete(sns string) (err error)
    PrinterStatus(sn string) (err error)
    IsPrintOk(strorderid string)
    PrinterLog(sn string, strdate string)

# rhtool_core
 ## -rcron 定时器
    参考test - TestCronByCronStr

 ## -remail email发送
    参考test - TestEmail

 ## -rexcel 导入导出 

 ## -rfile 文件操作

 ## -rgroup 开带context 协程并发执行

 ## -rmap 用来处理去重 value为空的struct
    NewStringSet： make一个key为string的map
    NewUint64Set： make一个key为uint64的map

 ## -rmath 数运算

 ## -rnumber 数字转换

 ## -rslice 切片转换

 ## -rstrings 字符串转换

 ## -rtime 时间转换

# 本地缓存

# 雪花算法
    可以配合本地緩存,实现单调递增
    dm := NewCustomNode()
	fmt.Println(dm.GenerateID().UInt64())
 
    