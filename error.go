package misc
import (
	"fmt"
	"os"
)
func E(desc string, err interface{}) {
	switch fault := err.(type) {
	case error:
		{
			if fault != nil {
				panic(fmt.Errorf(desc+": %v", err))
			}
		}
	case bool:
		{
			if !fault {
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
