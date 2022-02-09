package structor

type Expr func() Constructor

var createExpr Expr

func SetCreateExpr(e Expr) {
	createExpr = e
}

func GetCreateExpr(table string, fields Fields) Constructor {
	c := createExpr()
	c.Set(table, "", fields...)
	return c
}

var addExpr Expr

func SetAddExpr(e Expr) {
	addExpr = e
}

func GetAddExpr(table string, fields Fields) Constructor {
	c := addExpr()
	c.Set(table, "", fields...)
	return c
}

var modifyExpr Expr

func SetModifyExpr(e Expr) {
	modifyExpr = e
}

func GetModifyExpr(table string, fields Fields) Constructor {
	c := modifyExpr()
	c.Set(table, "", fields...)
	return c
}

var primary Expr

func SetPrimaryExpr(e Expr) {
	primary = e
}

func GetPriMaryExpr(table string, fields Fields) Constructor {
	c := primary()
	c.Set(table, "", fields...)
	return c
}

var index Expr

func SetIndexExpr(e Expr) {
	index = e
}

func GetIndexExpr(table string, IndexName string, fields Fields) Constructor {
	c := index()
	c.Set(table, IndexName, fields...)
	return c
}

var unique Expr

func SetUniqueExpr(e Expr) {
	unique = e
}

func GetUniqueExpr(table string, IndexName string, fields Fields) Constructor {
	c := unique()
	c.Set(table, IndexName, fields...)
	return c
}

var dropIndex Expr

func SetDropIndexExpr(e Expr) {
	dropIndex = e
}

func GetDropIndexExpr(table string, IndexName string, fields Fields) Constructor {
	c := dropIndex()
	c.Set(table, IndexName, fields...)
	return c
}
