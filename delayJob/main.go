package delayJob

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis/v7"
	"gomailme/cache"
	"strconv"
	"time"
)

type DelayJob struct {
	Method map[string]JobMsg
	//JobChan chan string
}

// 定义job的方法接口
type JobMsg interface {
	JobRunFun(id string) error
}

// 定义方法函数
type JobRunFun func(id string) error

// JobRunFun 调用自己(JobRunFun)实现接口
func (f JobRunFun) JobRunFun(id string) error { return f(id) }

type Job struct {
	Id    string
	Delay int64
}

const (
	Key = "delayJob"
)

var DJ *DelayJob

// 初始化对象
func Init() {
	var dj = DelayJob{
		Method: make(map[string]JobMsg),
	}
	DJ = &dj
	DJ.SetMethods()
	// 创建主轮询携程
	go DJ.runConsumer()
}

// 轮询rides
func (dj *DelayJob) runConsumer() {
	for {
		jobs, err := dj.GetJobs()
		if err != nil {
			fmt.Println("rides查询错误，暂停轮询60s", err.Error())
			time.Sleep(time.Second * 60)
			continue
		}
		if len(jobs) == 0 {
			time.Sleep(time.Second * 1)
			continue
		}
		for _, id := range jobs {
			go dj.DoJob(SendMailType,id)
		}
		// 删除ids
		cache.RedisClient.ZRemRangeByRank(Key, 0, int64(len(jobs)-1))
	}
}

// 加入事件（唯一id,执行时间戳）
func (dj *DelayJob) AddJob(id string, delay int64) error {
	// 将key：id value:delay加入rides的zset
	err := cache.RedisClient.ZAdd(Key, &redis.Z{Score: float64(delay), Member: id,}).Err()
	return err
}

// 查询要执行的事件
func (dj *DelayJob) GetJobs() ([]string, error) {
	// 查询需执行的列表
	op := redis.ZRangeBy{
		Min:    "0",
		Max:    strconv.FormatInt(time.Now().Unix(), 10),
		Offset: 0,
		Count:  20,
	}
	jobs, err := cache.RedisClient.ZRangeByScore(Key, &op).Result()
	return jobs, err
}

func (dj *DelayJob) DelJob(id string) error {
	// 删除某个固定任务
	i,_:=cache.RedisClient.ZRem(Key, id).Result()
	if i==0{
		return errors.New("任务删除失败")
	}
	return nil
}

func (dj *DelayJob) DoJob(cmd string,id string) {
	// 做任务
	err:=dj.Method[cmd].JobRunFun(id)
	if err!=nil{
		fmt.Println(time.Now(),id,"号任务执行失败",err.Error())
	}
}

func (dj *DelayJob)AddJobFun(cmd string,jrf JobRunFun){
	dj.Method[cmd] = jrf
}
