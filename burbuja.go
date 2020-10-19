package main

func Burbuja(s []int64)  {
	for i:=0; i<len(s)-1; i++{
		for j:=0; j<len(s)-i-1; j++{
			if s[j] > s[j+1]{
				s[j], s[j+1] = s[j+1],s[j]
			}
		}
	}
}

func main()  {
	a:=[]int64{5,2,8,4,9,1}
	Burbuja(a)
}