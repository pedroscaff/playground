// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package playground

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
	"gorm.io/playground/model"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newExample(db *gorm.DB, opts ...gen.DOOption) example {
	_example := example{}

	_example.exampleDo.UseDB(db, opts...)
	_example.exampleDo.UseModel(&model.Example{})

	tableName := _example.exampleDo.TableName()
	_example.ALL = field.NewAsterisk(tableName)
	_example.Example = field.NewString(tableName, "example")

	_example.fillFieldMap()

	return _example
}

type example struct {
	exampleDo exampleDo

	ALL     field.Asterisk
	Example field.String

	fieldMap map[string]field.Expr
}

func (e example) Table(newTableName string) *example {
	e.exampleDo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e example) As(alias string) *example {
	e.exampleDo.DO = *(e.exampleDo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *example) updateTableName(table string) *example {
	e.ALL = field.NewAsterisk(table)
	e.Example = field.NewString(table, "example")

	e.fillFieldMap()

	return e
}

func (e *example) WithContext(ctx context.Context) *exampleDo { return e.exampleDo.WithContext(ctx) }

func (e example) TableName() string { return e.exampleDo.TableName() }

func (e example) Alias() string { return e.exampleDo.Alias() }

func (e example) Columns(cols ...field.Expr) gen.Columns { return e.exampleDo.Columns(cols...) }

func (e *example) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *example) fillFieldMap() {
	e.fieldMap = make(map[string]field.Expr, 1)
	e.fieldMap["example"] = e.Example
}

func (e example) clone(db *gorm.DB) example {
	e.exampleDo.ReplaceConnPool(db.Statement.ConnPool)
	return e
}

func (e example) replaceDB(db *gorm.DB) example {
	e.exampleDo.ReplaceDB(db)
	return e
}

type exampleDo struct{ gen.DO }

func (e exampleDo) Debug() *exampleDo {
	return e.withDO(e.DO.Debug())
}

func (e exampleDo) WithContext(ctx context.Context) *exampleDo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e exampleDo) ReadDB() *exampleDo {
	return e.Clauses(dbresolver.Read)
}

func (e exampleDo) WriteDB() *exampleDo {
	return e.Clauses(dbresolver.Write)
}

func (e exampleDo) Session(config *gorm.Session) *exampleDo {
	return e.withDO(e.DO.Session(config))
}

func (e exampleDo) Clauses(conds ...clause.Expression) *exampleDo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e exampleDo) Returning(value interface{}, columns ...string) *exampleDo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e exampleDo) Not(conds ...gen.Condition) *exampleDo {
	return e.withDO(e.DO.Not(conds...))
}

func (e exampleDo) Or(conds ...gen.Condition) *exampleDo {
	return e.withDO(e.DO.Or(conds...))
}

func (e exampleDo) Select(conds ...field.Expr) *exampleDo {
	return e.withDO(e.DO.Select(conds...))
}

func (e exampleDo) Where(conds ...gen.Condition) *exampleDo {
	return e.withDO(e.DO.Where(conds...))
}

func (e exampleDo) Order(conds ...field.Expr) *exampleDo {
	return e.withDO(e.DO.Order(conds...))
}

func (e exampleDo) Distinct(cols ...field.Expr) *exampleDo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e exampleDo) Omit(cols ...field.Expr) *exampleDo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e exampleDo) Join(table schema.Tabler, on ...field.Expr) *exampleDo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e exampleDo) LeftJoin(table schema.Tabler, on ...field.Expr) *exampleDo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e exampleDo) RightJoin(table schema.Tabler, on ...field.Expr) *exampleDo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e exampleDo) Group(cols ...field.Expr) *exampleDo {
	return e.withDO(e.DO.Group(cols...))
}

func (e exampleDo) Having(conds ...gen.Condition) *exampleDo {
	return e.withDO(e.DO.Having(conds...))
}

func (e exampleDo) Limit(limit int) *exampleDo {
	return e.withDO(e.DO.Limit(limit))
}

func (e exampleDo) Offset(offset int) *exampleDo {
	return e.withDO(e.DO.Offset(offset))
}

func (e exampleDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *exampleDo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e exampleDo) Unscoped() *exampleDo {
	return e.withDO(e.DO.Unscoped())
}

func (e exampleDo) Create(values ...*model.Example) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e exampleDo) CreateInBatches(values []*model.Example, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e exampleDo) Save(values ...*model.Example) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e exampleDo) First() (*model.Example, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Example), nil
	}
}

func (e exampleDo) Take() (*model.Example, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Example), nil
	}
}

func (e exampleDo) Last() (*model.Example, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Example), nil
	}
}

func (e exampleDo) Find() ([]*model.Example, error) {
	result, err := e.DO.Find()
	return result.([]*model.Example), err
}

func (e exampleDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Example, err error) {
	buf := make([]*model.Example, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e exampleDo) FindInBatches(result *[]*model.Example, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e exampleDo) Attrs(attrs ...field.AssignExpr) *exampleDo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e exampleDo) Assign(attrs ...field.AssignExpr) *exampleDo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e exampleDo) Joins(fields ...field.RelationField) *exampleDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Joins(_f))
	}
	return &e
}

func (e exampleDo) Preload(fields ...field.RelationField) *exampleDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Preload(_f))
	}
	return &e
}

func (e exampleDo) FirstOrInit() (*model.Example, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Example), nil
	}
}

func (e exampleDo) FirstOrCreate() (*model.Example, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Example), nil
	}
}

func (e exampleDo) FindByPage(offset int, limit int) (result []*model.Example, count int64, err error) {
	result, err = e.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = e.Offset(-1).Limit(-1).Count()
	return
}

func (e exampleDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e exampleDo) Scan(result interface{}) (err error) {
	return e.DO.Scan(result)
}

func (e exampleDo) Delete(models ...*model.Example) (result gen.ResultInfo, err error) {
	return e.DO.Delete(models)
}

func (e *exampleDo) withDO(do gen.Dao) *exampleDo {
	e.DO = *do.(*gen.DO)
	return e
}
