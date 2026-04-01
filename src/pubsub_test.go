package gopattern

import (
	"log"
	"os"
	"testing"
	"time"
)

type CfgPublisher struct {
	cfgevent  *EventChannel
	filesName []string
}

func (cp CfgPublisher) Publish() {

}

func (cp CfgPublisher) Subscribe() *EventChannel {
	return cp.cfgevent
}

func (cp CfgPublisher) CheckOnUpdate() {
	for _, filePath := range cp.filesName {
		fileInfo, err := os.Stat(filePath)
		if err != nil {
			log.Println(err)
		} else {
			cp.cfgevent.SetEvent(EventMessage{Topic: filePath, Value: fileInfo.ModTime().Format("2006-01-02 15:04:05")})
		}
	}
}
func NewCfgPublisher(list []string) (*CfgPublisher, error) {
	var st CfgPublisher
	var se EventChannel
	st.cfgevent = &se
	st.filesName = list
	return &st, nil
}

func UpdateFile(fileName string) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Println("Ошибка при открытии/создании файла:", err)
		return
	}
	defer file.Close()
	update_str := time.Now().Format("2006-01-02 15:04:05")
	file.WriteString(update_str)
}

func Test_Publisher(t *testing.T) {
	tArr := []string{"test1.txt", "test2.txt"}
	tS, err := NewCfgPublisher(tArr)
	if err != nil {
		log.Println(err)
	}
	log.Println(*tS)
	timer0 := time.NewTimer(35 * time.Second)
	timer1 := time.NewTimer(55 * time.Second)
	timerChkUpd := time.NewTimer(10 * time.Second)
	for {
		select {
		case <-timer0.C:
			UpdateFile(tArr[0])
			timer0.Reset(15 * time.Second)
		case <-timer1.C:
			UpdateFile(tArr[1])
			timer1.Reset(25 * time.Second)		
		case <-timerChkUpd.C:
			tS.CheckOnUpdate()
			for _,tstr := range tArr {
				log.Println(tstr+"="+tS.cfgevent.GetValue(tstr))
			}
			timerChkUpd.Reset(10 * time.Second)
		}
	}
}
