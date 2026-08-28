package clock
import "fmt"

// Define the Clock type here.
type Clock struct {
    h int
    m int
}

func New(h, m int) Clock {
    hres,mres,exh :=0,0,0
  
    if m<0 {
        if (-m)%60==0 {
        	mres=0
        	exh=-(-m)/60
        } else {
        	mres=60-(-m)%60
        	exh=-(1+(-m)/60)
        }
    } else {
          mres=m%60
         exh=m/60
    }
    h=(h+exh)
    if h<0 {
        hres=24-(-h)%24
    } else {
        hres=h%24
    }

	return Clock{hres, mres}
}

func (c Clock) Add(m int) Clock {

    res := New((c.h),(c.m+m))
    return res
}

func (c Clock) Subtract(m int) Clock {
    extH:=0
    minusM := m%60
    min := c.m - minusM
    if min<0 {
        extH=1
        min+=60
    }
    minusH := (m/60+extH)%24
    h:= c.h - minusH
    if h<0 {
        h+=24
    }
    res := New(h,min)
    return res	
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d",c.h,c.m)
}
