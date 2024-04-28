package td

import "fmt"

func TdSample() {
	tConn := TestConn

	//插入数据
	insertStr := `INSERT INTO test.person (born, name, age) VALUES ('2024-04-18 12:00:02', 'Joh1n', 30)`
	tConn.Insert(insertStr)

	//获得所有数据
	res1, _ := tConn.GetAll("test", "person")
	fmt.Println(res1)

	//根据sql获得数据
	getStr := `select count(*) from test.person`
	res2, _ := tConn.GetBySql(getStr)
	fmt.Println(res2)
}
