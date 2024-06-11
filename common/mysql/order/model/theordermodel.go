package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TheorderModel = (*customTheorderModel)(nil)

type (
	// TheorderModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTheorderModel.
	TheorderModel interface {
		theorderModel
		withSession(session sqlx.Session) TheorderModel
	}

	customTheorderModel struct {
		*defaultTheorderModel
	}
)

// NewTheorderModel returns a model for the database table.
func NewTheorderModel(conn sqlx.SqlConn) TheorderModel {
	return &customTheorderModel{
		defaultTheorderModel: newTheorderModel(conn),
	}
}

func (m *customTheorderModel) withSession(session sqlx.Session) TheorderModel {
	return NewTheorderModel(sqlx.NewSqlConnFromSession(session))
}
