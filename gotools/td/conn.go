package td

import (
	"database/sql"
	"fmt"

	_ "github.com/taosdata/driver-go/v3/taosRestful"
)

var TdIp = "192.168.0.233"
var TdPort = "30247"
var TdUser = "jszhou"
var TdPassword = "123456"
var TdProtocol = "http"

var TestConn = TdEngine{
	TdIp, TdPort, TdUser, TdPassword, TdProtocol,
}

type TdEngine struct {
	IP       string
	Port     string
	User     string
	Password string
	Protocol string
}

func (t TdEngine) GetDSN() string {
	return fmt.Sprintf("%s:%s@%s(%s:%s)/", t.User, t.Password, t.Protocol, t.IP, t.Port)
}

func (t TdEngine) Insert(sqlStr string) error {
	DSN := t.GetDSN()
	taosConn, err := sql.Open("taosRestful", DSN)
	if err != nil {
		fmt.Println("failed to connect TDengine, err:", err)
		return err
	}
	defer taosConn.Close()
	_, err = taosConn.Exec(sqlStr)
	if err != nil {
		fmt.Println("failed to insert, err:", err)
		return err
	}
	return nil
}

func (t TdEngine) GetAll(database, table string) ([]map[string]interface{}, error) {
	var data []map[string]interface{}

	DSN := t.GetDSN()
	taosConn, err := sql.Open("taosRestful", DSN)
	if err != nil {
		fmt.Println("failed to connect TDengine, err:", err)
		return data, err
	}
	defer taosConn.Close()
	rows, err := taosConn.Query(fmt.Sprintf("select * from %s.%s", database, table))
	if err != nil {
		fmt.Println("failed to select from table, err:", err)
		return data, err
	}

	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return nil, fmt.Errorf("failed to get columns: %v", err)
	}

	var results []map[string]interface{}
	for rows.Next() {
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))
		for i := range columns {
			valuePtrs[i] = &values[i]
		}

		if err := rows.Scan(valuePtrs...); err != nil {
			return nil, fmt.Errorf("scan error: %v", err)
		}

		rowData := make(map[string]interface{})
		for i, col := range columns {
			val := values[i]
			if val == nil {
				rowData[col] = nil
			} else {
				rowData[col] = val
			}
		}
		results = append(results, rowData)
	}

	return results, nil
}

func (t TdEngine) GetBySql(sqlStr string) ([]map[string]interface{}, error) {
	var data []map[string]interface{}
	DSN := t.GetDSN()
	taosConn, err := sql.Open("taosRestful", DSN)
	if err != nil {
		fmt.Println("failed to connect TDengine, err:", err)
		return data, err
	}
	defer taosConn.Close()
	rows, err := taosConn.Query(sqlStr)
	if err != nil {
		fmt.Println("failed to select from table, err:", err)
		return data, err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return nil, fmt.Errorf("failed to get columns: %v", err)
	}
	var results []map[string]interface{}
	for rows.Next() {
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))
		for i := range columns {
			valuePtrs[i] = &values[i]
		}
		if err := rows.Scan(valuePtrs...); err != nil {
			return nil, fmt.Errorf("scan error: %v", err)
		}
		rowData := make(map[string]interface{})
		for i, col := range columns {
			val := values[i]
			if val == nil {
				rowData[col] = nil
			} else {
				rowData[col] = val
			}
		}
		results = append(results, rowData)
	}
	return results, nil
}

// TODO jszhou 后续可能修改成 sql+字段，通过sql获取单个字段的方法
// func TestTD() {
// 	var taosUri = "jszhou:123456@http(192.168.3.200:30247)/"
// 	fmt.Println(1111)
// 	taos, err := sql.Open("taosRestful", taosUri)
// 	if err != nil {
// 		fmt.Println("failed to connect TDengine, err:", err)
// 		return
// 	}
// 	defer taos.Close()
// 	fmt.Println(1112)
// 	// taos.Exec("create database if not exists test")
// 	// taos.Exec("use test")

// 	// taos.Exec("create table if not exists tb1 (ts timestamp, a int)")
// 	fmt.Println(1113)
// 	_, err = taos.Exec("INSERT INTO test.person (born, name, age) VALUES ('2024-04-18 12:00:02', 'Joh1n', 30)")
// 	if err != nil {
// 		fmt.Println("failed to insert, err:", err)
// 		return
// 	}
// 	rows, err := taos.Query("select * from test.person")
// 	if err != nil {
// 		fmt.Println("failed to select from table, err:", err)
// 		return
// 	}

// 	defer rows.Close()
// 	fmt.Println(1114)
// 	for rows.Next() {
// 		var r struct {
// 			born time.Time
// 			name string
// 			age  int
// 		}
// 		err := rows.Scan(&r.born, &r.name, &r.age)
// 		if err != nil {
// 			fmt.Println("scan error:\n", err)
// 			return
// 		}
// 		fmt.Println(r.born, r.name, r.age)
// 	}
// }
