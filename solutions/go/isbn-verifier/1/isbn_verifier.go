package isbnverifier
import (
    "fmt"
    "strings")

func IsValidISBN(isbn string) bool {
	isbn=strings.ReplaceAll(isbn,"-","")
    fmt.Println(isbn)
    if len(isbn)!=10 {
        return false
    }
    var res int
    for i, s:= range isbn {
        if s=='X'&&i==9 {
            fmt.Println(s)
            res+=(10-i)*10
        } else {
            n := int(s-'0')
            fmt.Println(n)
            if n<0 || n>9 {
                return false
            }
            res+=(10-i)*n
        }
        fmt.Printf("res= %d\n", res)
    }
    return res%11==0
}
