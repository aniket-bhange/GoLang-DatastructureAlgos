package utils

//Swap method to sawp address of two veriables
func Swap(xp *int, yp *int) {
	temp := *xp
	*xp = *yp
	*yp = temp
}
