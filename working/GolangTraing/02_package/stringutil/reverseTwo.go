package stringutil

func reverseTwo(s string) string { //这里s="!oG ,olleH"
	r := []rune(s) //将s转换为rune类型的切片并赋值给r。rune 等同于int32,常用来处理unicode或utf-8字符
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i] //交换r[i], r[j]的值
	}
	return string(r)
	//第1次循环i=0,j=9,满足循环条件i < len(r)/2=5。r[0]=!,  r[9]=H，交换后r[0]=H，r[9]=！
	//第2次循环i=1,j=8,满足循环条件i < len(r)/2=5。r[1]=o,  r[8]=e，交换后r[1]=e，r[8]=o
	//第3次循环i=2,j=7,满足循环条件i < len(r)/2=5。r[1]=G,  r[7]=l，交换后r[2]=l，r[7]=G
	//第4次循环i=3,j=6,满足循环条件i < len(r)/2=5。r[1]=" ",r[6]=l，交换后r[3]=l，r[6]=" "
	//第5次循环i=4,j=5,满足循环条件i < len(r)/2=5。r[1]=o,  r[5]=o，交换后r[4]=o，r[5]=,
	//第6次循环i=5,j=6,不满足循环条件i < len(r)/2=5。退出循环。此时r:=[]rune(H,e,l,l,o,','," ",G,o,!)
	//最终将[]rune类型的r转换成string类型并返回

}
