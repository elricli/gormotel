package gormotel

import (
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

func (p plugin) beforeCreate(db *gorm.DB) {
	p.before("gorm.create", db)
}

func (p plugin) afterCreate(db *gorm.DB) {
	p.after(db)
}

func (p plugin) beforeUpdate(db *gorm.DB) {
	p.before("gorm.update", db)
}

func (p plugin) afterUpdate(db *gorm.DB) {
	p.after(db)
}

func (p plugin) beforeDelete(db *gorm.DB) {
	p.before("gorm.delete", db)
}

func (p plugin) afterDelete(db *gorm.DB) {
	p.after(db)
}

func (p plugin) beforeQuery(db *gorm.DB) {
	p.before("gorm.query", db)
}

func (p plugin) afterQuery(db *gorm.DB) {
	p.after(db)
}

func (p plugin) beforeRow(db *gorm.DB) {
	p.before("gorm.row", db)
}

func (p plugin) afterRow(db *gorm.DB) {
	p.after(db)
}

func (p plugin) beforeRaw(db *gorm.DB) {
	p.before("gorm.raw", db)
}

func (p plugin) afterRaw(db *gorm.DB) {
	p.after(db)
}

func (p plugin) before(name string, db *gorm.DB) {
	if !trace.SpanFromContext(db.Statement.Context).IsRecording() {
		return
	}
	ctx, span := p.tracer.Start(db.Statement.Context, name)
	span.SetAttributes(
		attribute.String("gorm.table", db.Statement.Table),
	)
	db.Statement.Context = ctx
}

func (p plugin) after(db *gorm.DB) {
	span := trace.SpanFromContext(db.Statement.Context)
	if db.Error != nil && db.Error != gorm.ErrRecordNotFound {
		span.RecordError(db.Error)
		span.SetStatus(codes.Error, db.Error.Error())
	} else {
		span.SetStatus(codes.Ok, "OK")
	}
	span.SetAttributes(
		attribute.String("gorm.statement", db.Statement.SQL.String()),
		attribute.Int64("gorm.rowsAffected", db.RowsAffected),
	)
	span.End()
}
