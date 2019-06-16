package main

// 自定义异常
type InvalidRadiusError struct{
	Radius float64
	MinRadius float64
	MaxRadius float64
}

func (ire *InvalidRadiusError)Error() string {
	info := fmt.Sprintf("%.2f是非法半径，合法半径的范围是[%.2f,%.2f]\n",ire.Radius,ire.MinRadius, ire.MaxRadius)
	return info 
}
func NewInvalidRediusError(radius float64)*InvalidRadiusError {
		ire := new(InvalidRadiusError)
		ire.MinRadius = 5
		ire.MaxRadius = 50
		ire.Radius = radius
		return ire
}

func GetToBallVolumn3(radius float64)(volum float64,err error) {
	if radius < 0 {
		panic(NewInvalidRediusError(radius))
	}

	if radius < 5 || radius > 50 {
		
	}

	return (4/3.0) * math.Pi * math.Pow(radius, 3), nil
}
func main()  {
defer func () {
	err := recover()
	if err != nil {
		fmt.Println(err)
	}
}
	volum, err := GetToBallVolumn3(0.5)
	if err != nil {
		fmt.Println("获取体积失败，err=",err)
	}else{
		fmt.Println("小球的体积是",volum)
	}
}