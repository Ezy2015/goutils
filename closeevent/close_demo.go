package closeevent

import (
	"os"
	"fmt"
)

func demo(){
	c := make(chan os.Signal)
	signalNotify(c)
	for {
		fmt.Println("run...", err)
		
		s := <-c

		//�յ��źź�Ĵ�������ֻ������ź����ݣ�������һЩ������˼����
		fmt.Println("get signal:", s)
		fmt.Println("close...")
		
		//do your own close...
		//doCloseFunction()
			
		break
	}
}
