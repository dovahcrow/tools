package misc

import (
	"fmt"
	"os"
	"runtime"
)

func E(desc string, callframe bool, err interface{}) {
	switch fault := err.(type) {
	case error:
		{
			if fault != nil {
				if callframe {
					for i := 1; ; i++ {
						_, b, c, d := runtime.Caller(i)
						if !d {
							break
						}
						fmt.Printf("Call From %+v At %+v\n", b, c)
					}
				}
				panic(fmt.Errorf(desc+": %v", err))
			}
		}
	case bool:
		{
			if !fault {
				if callframe {

					for i := 1; ; i++ {
						_, b, c, d := runtime.Caller(i)
						if !d {
							break
						}
						fmt.Printf("Call From %+v At %+v\n", b, c)
					}
				}
				panic(fmt.Errorf(desc))
			}
		}
	}

}

func ErrHandle(err error, c string, des ...string) {
	if err != nil {
		switch c {
		case `p`:

			{
				for _, v := range des {
					fmt.Println(v)
				}
				panic(err)
			}
		case `x`:
			{
				for _, v := range des {
					fmt.Println(v)
				}
				fmt.Println(err)
				os.Exit(0)
			}
		case `n`:
			{
				for _, v := range des {
					fmt.Println(v)
				}
				fmt.Println(err)
			}
		}
	}
}
