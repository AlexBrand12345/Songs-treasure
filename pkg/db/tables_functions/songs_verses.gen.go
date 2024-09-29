// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package tables_functions

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"songs-treasure/pkg/db/model"
)

func newSongsVerse(db *gorm.DB, opts ...gen.DOOption) songsVerse {
	_songsVerse := songsVerse{}

	_songsVerse.songsVerseDo.UseDB(db, opts...)
	_songsVerse.songsVerseDo.UseModel(&model.SongsVerse{})

	tableName := _songsVerse.songsVerseDo.TableName()
	_songsVerse.ALL = field.NewAsterisk(tableName)
	_songsVerse.SongID = field.NewInt32(tableName, "song_id")
	_songsVerse.Verses = field.NewString(tableName, "verses")
	_songsVerse.Tsv = field.NewString(tableName, "tsv")

	_songsVerse.fillFieldMap()

	return _songsVerse
}

type songsVerse struct {
	songsVerseDo

	ALL    field.Asterisk
	SongID field.Int32
	Verses field.String
	Tsv    field.String

	fieldMap map[string]field.Expr
}

func (s songsVerse) Table(newTableName string) *songsVerse {
	s.songsVerseDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s songsVerse) As(alias string) *songsVerse {
	s.songsVerseDo.DO = *(s.songsVerseDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *songsVerse) updateTableName(table string) *songsVerse {
	s.ALL = field.NewAsterisk(table)
	s.SongID = field.NewInt32(table, "song_id")
	s.Verses = field.NewString(table, "verses")
	s.Tsv = field.NewString(table, "tsv")

	s.fillFieldMap()

	return s
}

func (s *songsVerse) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *songsVerse) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 3)
	s.fieldMap["song_id"] = s.SongID
	s.fieldMap["verses"] = s.Verses
	s.fieldMap["tsv"] = s.Tsv
}

func (s songsVerse) clone(db *gorm.DB) songsVerse {
	s.songsVerseDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s songsVerse) replaceDB(db *gorm.DB) songsVerse {
	s.songsVerseDo.ReplaceDB(db)
	return s
}

type songsVerseDo struct{ gen.DO }

type ISongsVerseDo interface {
	gen.SubQuery
	Debug() ISongsVerseDo
	WithContext(ctx context.Context) ISongsVerseDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISongsVerseDo
	WriteDB() ISongsVerseDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISongsVerseDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISongsVerseDo
	Not(conds ...gen.Condition) ISongsVerseDo
	Or(conds ...gen.Condition) ISongsVerseDo
	Select(conds ...field.Expr) ISongsVerseDo
	Where(conds ...gen.Condition) ISongsVerseDo
	Order(conds ...field.Expr) ISongsVerseDo
	Distinct(cols ...field.Expr) ISongsVerseDo
	Omit(cols ...field.Expr) ISongsVerseDo
	Join(table schema.Tabler, on ...field.Expr) ISongsVerseDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISongsVerseDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISongsVerseDo
	Group(cols ...field.Expr) ISongsVerseDo
	Having(conds ...gen.Condition) ISongsVerseDo
	Limit(limit int) ISongsVerseDo
	Offset(offset int) ISongsVerseDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISongsVerseDo
	Unscoped() ISongsVerseDo
	Create(values ...*model.SongsVerse) error
	CreateInBatches(values []*model.SongsVerse, batchSize int) error
	Save(values ...*model.SongsVerse) error
	First() (*model.SongsVerse, error)
	Take() (*model.SongsVerse, error)
	Last() (*model.SongsVerse, error)
	Find() ([]*model.SongsVerse, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SongsVerse, err error)
	FindInBatches(result *[]*model.SongsVerse, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SongsVerse) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISongsVerseDo
	Assign(attrs ...field.AssignExpr) ISongsVerseDo
	Joins(fields ...field.RelationField) ISongsVerseDo
	Preload(fields ...field.RelationField) ISongsVerseDo
	FirstOrInit() (*model.SongsVerse, error)
	FirstOrCreate() (*model.SongsVerse, error)
	FindByPage(offset int, limit int) (result []*model.SongsVerse, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISongsVerseDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s songsVerseDo) Debug() ISongsVerseDo {
	return s.withDO(s.DO.Debug())
}

func (s songsVerseDo) WithContext(ctx context.Context) ISongsVerseDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s songsVerseDo) ReadDB() ISongsVerseDo {
	return s.Clauses(dbresolver.Read)
}

func (s songsVerseDo) WriteDB() ISongsVerseDo {
	return s.Clauses(dbresolver.Write)
}

func (s songsVerseDo) Session(config *gorm.Session) ISongsVerseDo {
	return s.withDO(s.DO.Session(config))
}

func (s songsVerseDo) Clauses(conds ...clause.Expression) ISongsVerseDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s songsVerseDo) Returning(value interface{}, columns ...string) ISongsVerseDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s songsVerseDo) Not(conds ...gen.Condition) ISongsVerseDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s songsVerseDo) Or(conds ...gen.Condition) ISongsVerseDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s songsVerseDo) Select(conds ...field.Expr) ISongsVerseDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s songsVerseDo) Where(conds ...gen.Condition) ISongsVerseDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s songsVerseDo) Order(conds ...field.Expr) ISongsVerseDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s songsVerseDo) Distinct(cols ...field.Expr) ISongsVerseDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s songsVerseDo) Omit(cols ...field.Expr) ISongsVerseDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s songsVerseDo) Join(table schema.Tabler, on ...field.Expr) ISongsVerseDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s songsVerseDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISongsVerseDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s songsVerseDo) RightJoin(table schema.Tabler, on ...field.Expr) ISongsVerseDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s songsVerseDo) Group(cols ...field.Expr) ISongsVerseDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s songsVerseDo) Having(conds ...gen.Condition) ISongsVerseDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s songsVerseDo) Limit(limit int) ISongsVerseDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s songsVerseDo) Offset(offset int) ISongsVerseDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s songsVerseDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISongsVerseDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s songsVerseDo) Unscoped() ISongsVerseDo {
	return s.withDO(s.DO.Unscoped())
}

func (s songsVerseDo) Create(values ...*model.SongsVerse) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s songsVerseDo) CreateInBatches(values []*model.SongsVerse, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s songsVerseDo) Save(values ...*model.SongsVerse) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s songsVerseDo) First() (*model.SongsVerse, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SongsVerse), nil
	}
}

func (s songsVerseDo) Take() (*model.SongsVerse, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SongsVerse), nil
	}
}

func (s songsVerseDo) Last() (*model.SongsVerse, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SongsVerse), nil
	}
}

func (s songsVerseDo) Find() ([]*model.SongsVerse, error) {
	result, err := s.DO.Find()
	return result.([]*model.SongsVerse), err
}

func (s songsVerseDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SongsVerse, err error) {
	buf := make([]*model.SongsVerse, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s songsVerseDo) FindInBatches(result *[]*model.SongsVerse, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s songsVerseDo) Attrs(attrs ...field.AssignExpr) ISongsVerseDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s songsVerseDo) Assign(attrs ...field.AssignExpr) ISongsVerseDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s songsVerseDo) Joins(fields ...field.RelationField) ISongsVerseDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s songsVerseDo) Preload(fields ...field.RelationField) ISongsVerseDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s songsVerseDo) FirstOrInit() (*model.SongsVerse, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SongsVerse), nil
	}
}

func (s songsVerseDo) FirstOrCreate() (*model.SongsVerse, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SongsVerse), nil
	}
}

func (s songsVerseDo) FindByPage(offset int, limit int) (result []*model.SongsVerse, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s songsVerseDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s songsVerseDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s songsVerseDo) Delete(models ...*model.SongsVerse) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *songsVerseDo) withDO(do gen.Dao) *songsVerseDo {
	s.DO = *do.(*gen.DO)
	return s
}
