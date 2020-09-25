package service

type HelloService struct {

}
//只有满足如下标准的方法才能用于远程访问，其余方法会被忽略：
//- 方法是导出的
//- 方法有两个参数，都是导出类型或内建类型
//- 方法的第二个参数是指针
//- 方法只有一个error接口类型的返回值
//func (t *T) MethodName(argType T1, replyType *T2) error

func (p *HelloService) Hello(req string, reply *string) error {
	*reply = "hello:"+req
	return nil
}
