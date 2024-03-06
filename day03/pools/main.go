package main 

//定义一个任务Task
type Task struck {
	f func() error //一个Task里面应该有一个具体的业务，业务名称叫f
}

// 创建一个Task业务
func NewTask(arg_f func() error) *Task{
	t := Task{
		f : arg_f
	}
	return &t
}

func (t *Task) Execute(){
	t.f()
}