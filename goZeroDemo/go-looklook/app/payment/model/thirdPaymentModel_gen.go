// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	thirdPaymentFieldNames          = builder.RawFieldNames(&ThirdPayment{})
	thirdPaymentRows                = strings.Join(thirdPaymentFieldNames, ",")
	thirdPaymentRowsExpectAutoSet   = strings.Join(stringx.Remove(thirdPaymentFieldNames, "`id`", "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`"), ",")
	thirdPaymentRowsWithPlaceHolder = strings.Join(stringx.Remove(thirdPaymentFieldNames, "`id`", "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`"), "=?,") + "=?"

	cacheLooklookPaymentThirdPaymentIdPrefix = "cache:looklookPayment:thirdPayment:id:"
	cacheLooklookPaymentThirdPaymentSnPrefix = "cache:looklookPayment:thirdPayment:sn:"
)

type (
	thirdPaymentModel interface {
		Insert(ctx context.Context, data *ThirdPayment) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*ThirdPayment, error)
		FindOneBySn(ctx context.Context, sn string) (*ThirdPayment, error)
		Update(ctx context.Context, data *ThirdPayment) error
		Delete(ctx context.Context, id int64) error
	}

	defaultThirdPaymentModel struct {
		sqlc.CachedConn
		table string
	}

	ThirdPayment struct {
		Id             int64     `db:"id"`
		Sn             string    `db:"sn"` // 流水单号
		CreateTime     time.Time `db:"create_time"`
		UpdateTime     time.Time `db:"update_time"`
		DeleteTime     time.Time `db:"delete_time"`
		DelState       int64     `db:"del_state"`
		Version        int64     `db:"version"`          // 乐观锁版本号
		UserId         int64     `db:"user_id"`          // 用户id
		PayMode        string    `db:"pay_mode"`         // 支付方式 1:微信支付
		TradeType      string    `db:"trade_type"`       // 第三方支付类型
		TradeState     string    `db:"trade_state"`      // 第三方交易状态
		PayTotal       int64     `db:"pay_total"`        // 支付总金额(分)
		TransactionId  string    `db:"transaction_id"`   // 第三方支付单号
		TradeStateDesc string    `db:"trade_state_desc"` // 支付状态描述
		OrderSn        string    `db:"order_sn"`         // 业务单号
		ServiceType    string    `db:"service_type"`     // 业务类型
		PayStatus      int64     `db:"pay_status"`       // 平台内交易状态   -1:支付失败 0:未支付 1:支付成功 2:已退款
		PayTime        time.Time `db:"pay_time"`         // 支付成功时间
	}
)

func newThirdPaymentModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultThirdPaymentModel {
	return &defaultThirdPaymentModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`third_payment`",
	}
}

func (m *defaultThirdPaymentModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	looklookPaymentThirdPaymentIdKey := fmt.Sprintf("%s%v", cacheLooklookPaymentThirdPaymentIdPrefix, id)
	looklookPaymentThirdPaymentSnKey := fmt.Sprintf("%s%v", cacheLooklookPaymentThirdPaymentSnPrefix, data.Sn)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, looklookPaymentThirdPaymentIdKey, looklookPaymentThirdPaymentSnKey)
	return err
}

func (m *defaultThirdPaymentModel) FindOne(ctx context.Context, id int64) (*ThirdPayment, error) {
	looklookPaymentThirdPaymentIdKey := fmt.Sprintf("%s%v", cacheLooklookPaymentThirdPaymentIdPrefix, id)
	var resp ThirdPayment
	err := m.QueryRowCtx(ctx, &resp, looklookPaymentThirdPaymentIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", thirdPaymentRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultThirdPaymentModel) FindOneBySn(ctx context.Context, sn string) (*ThirdPayment, error) {
	looklookPaymentThirdPaymentSnKey := fmt.Sprintf("%s%v", cacheLooklookPaymentThirdPaymentSnPrefix, sn)
	var resp ThirdPayment
	err := m.QueryRowIndexCtx(ctx, &resp, looklookPaymentThirdPaymentSnKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `sn` = ? limit 1", thirdPaymentRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, sn); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultThirdPaymentModel) Insert(ctx context.Context, data *ThirdPayment) (sql.Result, error) {
	looklookPaymentThirdPaymentIdKey := fmt.Sprintf("%s%v", cacheLooklookPaymentThirdPaymentIdPrefix, data.Id)
	looklookPaymentThirdPaymentSnKey := fmt.Sprintf("%s%v", cacheLooklookPaymentThirdPaymentSnPrefix, data.Sn)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, thirdPaymentRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Sn, data.DeleteTime, data.DelState, data.Version, data.UserId, data.PayMode, data.TradeType, data.TradeState, data.PayTotal, data.TransactionId, data.TradeStateDesc, data.OrderSn, data.ServiceType, data.PayStatus, data.PayTime)
	}, looklookPaymentThirdPaymentIdKey, looklookPaymentThirdPaymentSnKey)
	return ret, err
}

func (m *defaultThirdPaymentModel) Update(ctx context.Context, newData *ThirdPayment) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	looklookPaymentThirdPaymentIdKey := fmt.Sprintf("%s%v", cacheLooklookPaymentThirdPaymentIdPrefix, data.Id)
	looklookPaymentThirdPaymentSnKey := fmt.Sprintf("%s%v", cacheLooklookPaymentThirdPaymentSnPrefix, data.Sn)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, thirdPaymentRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.Sn, newData.DeleteTime, newData.DelState, newData.Version, newData.UserId, newData.PayMode, newData.TradeType, newData.TradeState, newData.PayTotal, newData.TransactionId, newData.TradeStateDesc, newData.OrderSn, newData.ServiceType, newData.PayStatus, newData.PayTime, newData.Id)
	}, looklookPaymentThirdPaymentIdKey, looklookPaymentThirdPaymentSnKey)
	return err
}

func (m *defaultThirdPaymentModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheLooklookPaymentThirdPaymentIdPrefix, primary)
}

func (m *defaultThirdPaymentModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", thirdPaymentRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultThirdPaymentModel) tableName() string {
	return m.table
}
