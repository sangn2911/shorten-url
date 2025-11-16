package repository

import (
	"context"
	"fmt"
	"shorten-url/internal/model"
	"shorten-url/package/database/mysql"
	"shorten-url/package/logger"
	"slices"
	"strings"
)

const table_url_encode = "url_encode"

type repository struct {
	db *mysql.Client
}

func (r *repository) GetUrlEncodeByShortenId(ctx context.Context, shortenId string) (urlEncode model.UrlEncodeModel, err error) {
	columns := []string{
		"shorten_id",
		"shorten_number",
		"long_url",
	}
	sql := fmt.Sprintf(
		"SELECT %s FROM %s WHERE shorten_id = ? LIMIT 1",
		strings.Join(columns, ", "),
		table_url_encode,
	)
	rows, err := r.db.GetDBTX(ctx).QueryContext(ctx, sql, shortenId)
	if err != nil {
		return urlEncode, err
	}
	defer rows.Close()
	if rows.Next() {
		if err := rows.Scan(
			&urlEncode.ShortenId,
			&urlEncode.ShortenNumber,
			&urlEncode.LongUrl,
		); err != nil {
			return urlEncode, err
		}
		return urlEncode, nil
	}
	return urlEncode, err
}

func (r *repository) GetUrlEncodeByLongUrl(ctx context.Context, longUrl string) (urlEncode model.UrlEncodeModel, err error) {
	columns := []string{
		"shorten_id",
		"shorten_number",
		"long_url",
	}
	sql := fmt.Sprintf(
		"SELECT %s FROM %s WHERE long_url = ? LIMIT 1",
		strings.Join(columns, ", "),
		table_url_encode,
	)
	rows, err := r.db.GetDBTX(ctx).QueryContext(ctx, sql, longUrl)
	if err != nil {
		return urlEncode, err
	}
	defer rows.Close()
	if rows.Next() {
		if err := rows.Scan(
			&urlEncode.ShortenId,
			&urlEncode.ShortenNumber,
			&urlEncode.LongUrl,
		); err != nil {
			return urlEncode, err
		}
		return urlEncode, nil
	}
	return urlEncode, err
}

func (r *repository) InsertUrlEncode(ctx context.Context, urlEncode model.UrlEncodeModel) error {
	columns := []string{
		"shorten_id",
		"shorten_number",
		"long_url",
	}
	sql := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		table_url_encode,
		joinColumns(columns),
		getValueHolder(len(columns)),
	)
	if _, err := r.db.GetDBTX(ctx).ExecContext(
		ctx,
		sql,
		urlEncode.ShortenId,
		urlEncode.ShortenNumber,
		urlEncode.LongUrl,
	); err != nil {
		return err
	}
	return nil
}

func (r *repository) Begin(ctx context.Context) (context.Context, error) {
	return r.db.Begin(ctx)
}

func (r *repository) Commit(ctx context.Context) error {
	return r.db.Commit(ctx)
}

func (r *repository) RollBack(ctx context.Context) error {
	if err := r.db.RollBack(ctx); err != nil {
		logger.WithContext(ctx).Error(err, "fail to rollback transaction")
	}
	return nil
}

func (r *repository) GetLatestUrlEncode(ctx context.Context) (urlEncode model.UrlEncodeModel, err error) {
	columns := []string{
		"shorten_id",
		"shorten_number",
		"long_url",
	}
	sql := fmt.Sprintf(
		"SELECT %s FROM %s ORDER BY shorten_number DESC LIMIT 1",
		strings.Join(columns, ", "),
		table_url_encode,
	)
	rows, err := r.db.GetDBTX(ctx).QueryContext(ctx, sql)
	if err != nil {
		return urlEncode, err
	}
	defer rows.Close()
	if rows.Next() {
		if err := rows.Scan(
			&urlEncode.ShortenId,
			&urlEncode.ShortenNumber,
			&urlEncode.LongUrl,
		); err != nil {
			return urlEncode, err
		}
		return urlEncode, nil
	}
	return urlEncode, err
}

func NewRepository(db *mysql.Client) *repository {
	return &repository{
		db: db,
	}
}

func joinColumns(columns []string) string {
	return strings.Join(columns, ", ")
}

func getValueHolder(count int) string {
	return strings.Join(slices.Repeat([]string{"?"}, count), ", ")
}
