package database

var (

	//查询数据表test中的所有数据
	QueryAllData = `SELECT *FROM test;`

	//查询数据表login登录用户的信息
	QueryLogin = `SELECT *FROM login`

	//单例查询by name前缀
	QueryDataByName = `SELECT *FROM test WHERE name like ?;`

	//单例删除by id
	DeleteSingleDataById = `DELETE FROM test WHERE id = ?;`

	//单例更新 price by id
	UpdateSingleDataById = `UPDATE test SET name = ?, price = ? WHERE id = ?;`

	//单例add
	PostSingleData = `insert into test (name, price) values(?,?);`

	//添加用户
	PostUsername = `insert into login (username, password) values(?,?);`

)