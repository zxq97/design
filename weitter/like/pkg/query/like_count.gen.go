// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/zxq97/design/weitter/like/pkg/model"
)

func newLikeCount(db *gorm.DB, opts ...gen.DOOption) likeCount {
	_likeCount := likeCount{}

	_likeCount.likeCountDo.UseDB(db, opts...)
	_likeCount.likeCountDo.UseModel(&model.LikeCount{})

	tableName := _likeCount.likeCountDo.TableName()
	_likeCount.ALL = field.NewAsterisk(tableName)
	_likeCount.ID = field.NewInt64(tableName, "id")
	_likeCount.ObjID = field.NewInt64(tableName, "obj_id")
	_likeCount.ObjType = field.NewInt32(tableName, "obj_type")
	_likeCount.Count = field.NewInt32(tableName, "count")
	_likeCount.Ctime = field.NewTime(tableName, "ctime")
	_likeCount.Mtime = field.NewTime(tableName, "mtime")

	_likeCount.fillFieldMap()

	return _likeCount
}

type likeCount struct {
	likeCountDo likeCountDo

	ALL     field.Asterisk
	ID      field.Int64
	ObjID   field.Int64
	ObjType field.Int32
	Count   field.Int32
	Ctime   field.Time
	Mtime   field.Time

	fieldMap map[string]field.Expr
}

func (l likeCount) Table(newTableName string) *likeCount {
	l.likeCountDo.UseTable(newTableName)
	return l.updateTableName(newTableName)
}

func (l likeCount) As(alias string) *likeCount {
	l.likeCountDo.DO = *(l.likeCountDo.As(alias).(*gen.DO))
	return l.updateTableName(alias)
}

func (l *likeCount) updateTableName(table string) *likeCount {
	l.ALL = field.NewAsterisk(table)
	l.ID = field.NewInt64(table, "id")
	l.ObjID = field.NewInt64(table, "obj_id")
	l.ObjType = field.NewInt32(table, "obj_type")
	l.Count = field.NewInt32(table, "count")
	l.Ctime = field.NewTime(table, "ctime")
	l.Mtime = field.NewTime(table, "mtime")

	l.fillFieldMap()

	return l
}

func (l *likeCount) WithContext(ctx context.Context) *likeCountDo {
	return l.likeCountDo.WithContext(ctx)
}

func (l likeCount) TableName() string { return l.likeCountDo.TableName() }

func (l likeCount) Alias() string { return l.likeCountDo.Alias() }

func (l likeCount) Columns(cols ...field.Expr) gen.Columns { return l.likeCountDo.Columns(cols...) }

func (l *likeCount) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := l.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (l *likeCount) fillFieldMap() {
	l.fieldMap = make(map[string]field.Expr, 6)
	l.fieldMap["id"] = l.ID
	l.fieldMap["obj_id"] = l.ObjID
	l.fieldMap["obj_type"] = l.ObjType
	l.fieldMap["count"] = l.Count
	l.fieldMap["ctime"] = l.Ctime
	l.fieldMap["mtime"] = l.Mtime
}

func (l likeCount) clone(db *gorm.DB) likeCount {
	l.likeCountDo.ReplaceConnPool(db.Statement.ConnPool)
	return l
}

func (l likeCount) replaceDB(db *gorm.DB) likeCount {
	l.likeCountDo.ReplaceDB(db)
	return l
}

type likeCountDo struct{ gen.DO }

// sql(insert into @@table(obj_id, obj_type, count) values (@objID, @objType, @count) on duplicate key update count=count+@count)
func (l likeCountDo) IncrByLikeCount(objID int64, objType int32, count int32) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, objID)
	params = append(params, objType)
	params = append(params, count)
	params = append(params, count)
	generateSQL.WriteString("insert into like_count(obj_id, obj_type, count) values (?, ?, ?) on duplicate key update count=count+? ")

	var executeSQL *gorm.DB
	executeSQL = l.UnderlyingDB().Exec(generateSQL.String(), params...) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (l likeCountDo) Debug() *likeCountDo {
	return l.withDO(l.DO.Debug())
}

func (l likeCountDo) WithContext(ctx context.Context) *likeCountDo {
	return l.withDO(l.DO.WithContext(ctx))
}

func (l likeCountDo) ReadDB() *likeCountDo {
	return l.Clauses(dbresolver.Read)
}

func (l likeCountDo) WriteDB() *likeCountDo {
	return l.Clauses(dbresolver.Write)
}

func (l likeCountDo) Session(config *gorm.Session) *likeCountDo {
	return l.withDO(l.DO.Session(config))
}

func (l likeCountDo) Clauses(conds ...clause.Expression) *likeCountDo {
	return l.withDO(l.DO.Clauses(conds...))
}

func (l likeCountDo) Returning(value interface{}, columns ...string) *likeCountDo {
	return l.withDO(l.DO.Returning(value, columns...))
}

func (l likeCountDo) Not(conds ...gen.Condition) *likeCountDo {
	return l.withDO(l.DO.Not(conds...))
}

func (l likeCountDo) Or(conds ...gen.Condition) *likeCountDo {
	return l.withDO(l.DO.Or(conds...))
}

func (l likeCountDo) Select(conds ...field.Expr) *likeCountDo {
	return l.withDO(l.DO.Select(conds...))
}

func (l likeCountDo) Where(conds ...gen.Condition) *likeCountDo {
	return l.withDO(l.DO.Where(conds...))
}

func (l likeCountDo) Order(conds ...field.Expr) *likeCountDo {
	return l.withDO(l.DO.Order(conds...))
}

func (l likeCountDo) Distinct(cols ...field.Expr) *likeCountDo {
	return l.withDO(l.DO.Distinct(cols...))
}

func (l likeCountDo) Omit(cols ...field.Expr) *likeCountDo {
	return l.withDO(l.DO.Omit(cols...))
}

func (l likeCountDo) Join(table schema.Tabler, on ...field.Expr) *likeCountDo {
	return l.withDO(l.DO.Join(table, on...))
}

func (l likeCountDo) LeftJoin(table schema.Tabler, on ...field.Expr) *likeCountDo {
	return l.withDO(l.DO.LeftJoin(table, on...))
}

func (l likeCountDo) RightJoin(table schema.Tabler, on ...field.Expr) *likeCountDo {
	return l.withDO(l.DO.RightJoin(table, on...))
}

func (l likeCountDo) Group(cols ...field.Expr) *likeCountDo {
	return l.withDO(l.DO.Group(cols...))
}

func (l likeCountDo) Having(conds ...gen.Condition) *likeCountDo {
	return l.withDO(l.DO.Having(conds...))
}

func (l likeCountDo) Limit(limit int) *likeCountDo {
	return l.withDO(l.DO.Limit(limit))
}

func (l likeCountDo) Offset(offset int) *likeCountDo {
	return l.withDO(l.DO.Offset(offset))
}

func (l likeCountDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *likeCountDo {
	return l.withDO(l.DO.Scopes(funcs...))
}

func (l likeCountDo) Unscoped() *likeCountDo {
	return l.withDO(l.DO.Unscoped())
}

func (l likeCountDo) Create(values ...*model.LikeCount) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Create(values)
}

func (l likeCountDo) CreateInBatches(values []*model.LikeCount, batchSize int) error {
	return l.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (l likeCountDo) Save(values ...*model.LikeCount) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Save(values)
}

func (l likeCountDo) First() (*model.LikeCount, error) {
	if result, err := l.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.LikeCount), nil
	}
}

func (l likeCountDo) Take() (*model.LikeCount, error) {
	if result, err := l.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.LikeCount), nil
	}
}

func (l likeCountDo) Last() (*model.LikeCount, error) {
	if result, err := l.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.LikeCount), nil
	}
}

func (l likeCountDo) Find() ([]*model.LikeCount, error) {
	result, err := l.DO.Find()
	return result.([]*model.LikeCount), err
}

func (l likeCountDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.LikeCount, err error) {
	buf := make([]*model.LikeCount, 0, batchSize)
	err = l.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (l likeCountDo) FindInBatches(result *[]*model.LikeCount, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return l.DO.FindInBatches(result, batchSize, fc)
}

func (l likeCountDo) Attrs(attrs ...field.AssignExpr) *likeCountDo {
	return l.withDO(l.DO.Attrs(attrs...))
}

func (l likeCountDo) Assign(attrs ...field.AssignExpr) *likeCountDo {
	return l.withDO(l.DO.Assign(attrs...))
}

func (l likeCountDo) Joins(fields ...field.RelationField) *likeCountDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Joins(_f))
	}
	return &l
}

func (l likeCountDo) Preload(fields ...field.RelationField) *likeCountDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Preload(_f))
	}
	return &l
}

func (l likeCountDo) FirstOrInit() (*model.LikeCount, error) {
	if result, err := l.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.LikeCount), nil
	}
}

func (l likeCountDo) FirstOrCreate() (*model.LikeCount, error) {
	if result, err := l.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.LikeCount), nil
	}
}

func (l likeCountDo) FindByPage(offset int, limit int) (result []*model.LikeCount, count int64, err error) {
	result, err = l.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = l.Offset(-1).Limit(-1).Count()
	return
}

func (l likeCountDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = l.Count()
	if err != nil {
		return
	}

	err = l.Offset(offset).Limit(limit).Scan(result)
	return
}

func (l likeCountDo) Scan(result interface{}) (err error) {
	return l.DO.Scan(result)
}

func (l likeCountDo) Delete(models ...*model.LikeCount) (result gen.ResultInfo, err error) {
	return l.DO.Delete(models)
}

func (l *likeCountDo) withDO(do gen.Dao) *likeCountDo {
	l.DO = *do.(*gen.DO)
	return l
}
