package sumofmultiples
import "fmt"
func SumMultiples(limit int, divisors ...int) int {
	m:=make(map[int]bool)
    res:=0
    for _,i := range divisors {
        val := i
        fmt.Println(val)
        for val<limit&&val>0 {
            if m[val]==false {
          		m[val]=true
                res+=val
            }
            val+=i
        }
    }
    return res
}
