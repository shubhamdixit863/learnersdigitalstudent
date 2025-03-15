package services

import (
	"math/rand"
	"fmt"
	"time"
)

// "[2025-03-10 12:00:00] INFO: User logged in"

// "[2025-03-10 12:01:00] ERROR: Database connection failed"

// "[2025-03-10 12:02:00] WARN: Disk usage is high"


func RandomLog() string{
	Loglevel:=[]string{"Info","Error","Warn"}
	LogMeassage:=[]string{"User logged in","Database connection failed","Disk usage is high"}
	return fmt.Sprintf("[%s], %s ,%s",time.Now(),Loglevel[rand.Intn(len(Loglevel))],LogMeassage[rand.Intn(len(LogMeassage))])
}

func GenLogs(logchannel chan string,num int ){
	for i:=0;i<num;i++{
		logchannel<-RandomLog()
	}
	close(logchannel)
}

