package diff_test // test文件package声明可以不同，需添加“_test”

import "go-learn/syntax/02_package/diff" // 可以理解为test文件不同包，调用包内方法时需import

func UseHello() {
	diff.Hello()
}
