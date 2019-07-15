# LunarBirthday

Google日历并不支持农历提醒，所以，只能靠自己转化为cvs文件，导入到Google日历中。

## 用法

```
Usage of LunarBirthday:
  -b string
    	输入日期，1990-01-01
  -h	显示用法
  -l	是否为阴历/农历
  -n int
    	次数，默认31 (default 31)
  -s string
    	标题
  -v	show version
```

## 例子

假设母亲的生日为，阴历一月初一，执行以下命令，重定向到一个cvs文件，使用Google日历网页版导入即可。

```
✗ LunarBirthday -s 老妈生日 -b 1950-01-01 -l > mon.cvs                                                                                        ✗
```

