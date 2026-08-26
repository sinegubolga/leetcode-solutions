package phonenumber
import (
    "strings"
    "errors"
    "fmt"
)
func Number(phoneNumber string) (string, error) {
	phoneNumber = strings.TrimSpace(phoneNumber)
    if len(phoneNumber)<10 {
        return phoneNumber, errors.New("too short number")
    }
    var ph []int
    for _,val:=range phoneNumber {        
        if val>='0' && val<='9' {
            digit:=int(val-'0')
            ph=append(ph,digit)
        }
    }
    fmt.Println(ph)
    if len(ph)>11 || len(ph)<10 {
        return phoneNumber, errors.New("wrong number og digits")
    } else if len(ph)==11 {
        if ph[0]!=1 {
            return phoneNumber, errors.New("first digit must be 1")
        } else {
            ph=ph[1:]
        }
    }
    if ph[0]<2 || ph[3]<2 {
      return phoneNumber, errors.New("can't be 1")  
    }    
    return fmt.Sprintf("%d%d%d%d%d%d%d%d%d%d",ph[0],ph[1],ph[2],ph[3],ph[4],ph[5],ph[6],ph[7],ph[8],ph[9]),nil
    
    
}

func AreaCode(phoneNumber string) (string, error) {
	cleanPh,ok:=Number(phoneNumber)
    if ok!=nil {
        return phoneNumber,errors.New("Wrong number")
    }
    return cleanPh[:3],nil
}

func Format(phoneNumber string) (string, error) {
	clear,ok:=Number(phoneNumber)
    if ok!=nil {
        return phoneNumber,errors.New("Wrong number")
    }
    s := []rune(clear)
    return fmt.Sprintf("(%c%c%c) %c%c%c-%c%c%c%c",s[0],s[1],s[2],s[3],s[4],s[5],s[6],s[7],s[8],s[9]),nil

}
