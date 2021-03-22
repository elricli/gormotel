package gormotel

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

var Plugin = &plugin{tracer: otel.Tracer("gorm")}

type plugin struct {
	tracer trace.Tracer
}

func (p plugin) Name() string {
	return "opentelemetry"
}

func (p plugin) Initialize(db *gorm.DB) error {
	// create callback
	if err := db.Callback().Create().Before("gorm:create").Register("opentelemetry:before_create", p.beforeCreate); err != nil {
		return err
	}
	if err := db.Callback().Create().After("gorm:create").Register("opentelemetry:after_create", p.afterCreate); err != nil {
		return err
	}

	// update callback
	if err := db.Callback().Update().Before("gorm:update").Register("opentelemetry:before_update", p.beforeUpdate); err != nil {
		return err
	}
	if err := db.Callback().Update().After("gorm:update").Register("opentelemetry:after_update", p.afterUpdate); err != nil {
		return err
	}

	// query callback
	if err := db.Callback().Query().Before("gorm:query").Register("opentelemetry:before_query", p.beforeQuery); err != nil {
		return err
	}
	if err := db.Callback().Query().After("gorm:query").Register("opentelemetry:after_query", p.afterQuery); err != nil {
		return err
	}

	// delete callback
	if err := db.Callback().Delete().Before("gorm:delete").Register("opentelemetry:before_delete", p.beforeDelete); err != nil {
		return err
	}
	if err := db.Callback().Delete().After("gorm:delete").Register("opentelemetry:after_delete", p.afterDelete); err != nil {
		return err
	}

	// row callback
	if err := db.Callback().Row().Before("gorm:row").Register("opentelemetry:before_row", p.beforeRow); err != nil {
		return err
	}
	if err := db.Callback().Row().After("gorm:row").Register("opentelemetry:after_row", p.afterRow); err != nil {
		return err
	}

	// raw callback
	if err := db.Callback().Raw().Before("gorm:raw").Register("opentelemetry:before_raw", p.beforeRaw); err != nil {
		return err
	}
	if err := db.Callback().Raw().After("gorm:raw").Register("opentelemetry:after_raw", p.afterRaw); err != nil {
		return err
	}

	return nil
}
