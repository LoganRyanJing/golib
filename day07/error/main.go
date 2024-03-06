package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Addree    string
}

func main() {
	employee, err := getInformation(1001)
	if err != nil {
		// Something is wrong. Do something
	} else {
		fmt.Println(employee)
	}
}

func getInformation(id int) (*Employee, error) {
	// employee, err := apiCallEmployee(1000)
	// if err != nil {
	// 	return nil, err // Simply return the error to the caller.
	// }

	// if err != nil {
	// 	return nil, fmt.Errorf("Got an error when getting the employee information: %v", err)
	// }
	// return employee, nil

	// 三次等待2秒钟
	for tries := 0; tries < 3; tries++ {
		employee, err := apiCallEmployee(1000)
		if err == nil {
			return employee, nil
		}

		fmt.Println("Server is not responding, retrying ...")
		time.Sleep(time.Second * 2)
	}

	return nil, fmt.Errorf("server has failed to respond to get the employee information")

}

func apiCallEmployee(id int) (*Employee, error) {
	employee := Employee{LastName: "Done", FirstName: "John"}
	return &employee, nil
}

// 在 Go 中处理错误时，请记住下面一些推荐做法：

// 始终检查是否存在错误，即使预期不存在。 然后正确处理它们，以免向最终用户公开不必要的信息。
// 在错误消息中包含一个前缀，以便了解错误的来源。 例如，可以包含包和函数的名称。
// 创建尽可能多的可重用错误变量。
// 了解使用返回错误和 panic 之间的差异。 不能执行其他操作时再使用 panic。 例如，如果某个依赖项未准备就绪，则程序运行无意义（除非你想要运行默认行为）。
// 在记录错误时记录尽可能多的详细信息（我们将在下一部分介绍记录方法），并打印出最终用户能够理解的错误。
