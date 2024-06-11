package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ GoodsModel = (*customGoodsModel)(nil)

type (
	// GoodsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGoodsModel.
	GoodsModel interface {
		goodsModel
		withSession(session sqlx.Session) GoodsModel
	}

	customGoodsModel struct {
		*defaultGoodsModel
	}
)

// NewGoodsModel returns a model for the database table.
func NewGoodsModel(conn sqlx.SqlConn) GoodsModel {
	return &customGoodsModel{
		defaultGoodsModel: newGoodsModel(conn),
	}
}

func (m *customGoodsModel) withSession(session sqlx.Session) GoodsModel {
	return NewGoodsModel(sqlx.NewSqlConnFromSession(session))
}
