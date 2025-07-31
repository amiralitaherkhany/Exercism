package raindrops
import "fmt"
func Convert(number int) string {
	str := ""
    if number % 3 == 0{
        str += "Pling"
    }
    if number % 5 == 0{
        str += "Plang"
    }
    if number % 7 == 0{
        str += "Plong"
    }
    if str == ""{
        return fmt.Sprint(number)
    }
    return str
}
