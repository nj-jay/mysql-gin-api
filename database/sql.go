package database

var (

	//查询数据表test中的所有数据,根据前段传入的page进行分页展示
	QueryAllData = `SELECT *FROM bookManage limit ?, 8;`

	//查询数据表login登录用户的信息
	QueryLogin = `SELECT * FROM login`

	//单例查询by name or types or author进行模糊匹配
	QueryDataByName = `SELECT * FROM bookManage WHERE concat(name, types, author) REGEXP ?;`

	//单例删除by id
	DeleteSingleDataById = `DELETE FROM bookManage WHERE id = ?;`

	//单例更新 price by id
	UpdateSingleDataById = `UPDATE bookManage SET name = ?, types = ?, author = ?, price = ?, addTime = ? WHERE id = ?;`

	//单例add
	PostSingleData = `insert into bookManage (name, types, author, price, addTime) values(?,?,?,?,?);`

	//添加用户
	PostUsername = `insert into login (username, password) values(?,?);`

)