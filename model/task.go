package model

import "github.com/jinzhu/gorm"

// Task 任务模型
type Task struct {
	gorm.Model
	Name  string
	Type  string
	User  string
	Pass  string
	Count int    //表示任务的执行次数，-1表示执行无限次，0表示不用执行了
	Cron  string //表示任务的执行规律，满足cron规则
	To    string
	Msg   string `gorm:"size:1000"`
	Desc  string
}

const (
	SendMail string ="SendMail"
)
// corn
//字段名    			是否必须		允许的值    		允许的特定字符
//秒(Seconds)    		是    	0-59    		* /, –
//分(Minutes)    		是    	0-59    		* /, –
//时(Hours)    			是    	0-23    		* /, –
//日(Day of month)    	是    	1-31    		* /, – ?
//月(Month)    			是    	1-12 or JAN-DEC	* /, –
//星期(Day of week)   	否   	0-6 or SUM-SAT  * /, – ?
