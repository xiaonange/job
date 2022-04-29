ES常用的数据类型可分为3大类
核⼼数据类型
复杂数据类型
专⽤数据类型
 

 

核心数据类型
（1）字符串
text  ⽤于全⽂索引，搜索时会自动使用分词器进⾏分词再匹配
keyword  不分词，搜索时需要匹配完整的值
（2）数值型
整型： byte，short，integer，long
浮点型： float, half_float, scaled_float，double
（3）日期类型
date
json没有date类型，插入|更新文档|字段时怎么表示date类型？
#mapping，将字段类型设置为date
"type" : "date" 
#插入|更新此字段的值时，有3种表示方式
#使用固定格式的字符串
"2020-04-18"、"2020/04/18 09:00:00"   
#值使用长整型的时间戳，1970-01-01 00:00:00，s
1610350870    
#值使用长整型的时间戳，ms
1641886870000
（4）范围型
integer_range， long_range， float_range，double_range，date_range
比如招聘要求年龄在[20, 40]上，mapping：
age_limit :{
　"type" : "integer_range"
}
插入|更新文档|字段时，值写成json对象的形式：
"age_limit" : {
　"gte" : 20,
　"lte" : 40
}
gt是大于，lt是小于，e是equals等于。
按此字段搜索时，值写常量：
"term" : {
　"age_limit" : 30
}
age_limit的区间包含了此值的文档都算是匹配。
（5）布尔
boolean     #true、false
（6）⼆进制
binary   会把值当做经过 base64 编码的字符串，默认不存储，且不可搜索
复杂数据类型
（1）对象
object
#定义mapping
"user" : {
    "type":"object"
}
#插入|更新字段的值，值写成json对象的形式
"user" : {
    "name":"chy",
    "age":12
}
#搜索时，字段名使用点号连接
"match":{
     "user.name":"chy"
 }
一个对象中可以嵌套对象。
（2）数组
#ES没有专门的数组类型，定义mapping，写成元素的类型
"arr" : {
    "type":"integer"
}
#插入|更新字段的值。元素可以是各种类型，但元素的类型要相同
"arr" : [1,3,4]
专用数据类型
ip
#定义mapping
"ip_address" : {
    "type":"ip"
}
#插入|更新字段的值，值写成字符串形式
"ip" : "192.168.1.1"
#搜索
"match":{
     "ip_address":"192.168.1.1"
 }
#ip在192.168.0.0 ~ 192.168.255.255上的文档都匹配
"match":{
     "ip_address":"192.168.0.0/16"
 }