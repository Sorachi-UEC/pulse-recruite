package repository

import (
    "context"
    "github.com/pulse227/server-recruit-challenge-sample/model"
    "database/sql"
    "fmt"
)

// AlbumRepository はアルバムのデータを操作するためのインターフェースです
type AlbumRepository interface {
    GetAll(ctx context.Context) ([]*model.Album, error)
    Get(ctx context.Context, id model.AlbumID) (*model.Album, error)
    Add(ctx context.Context, album *model.Album) error
    Delete(ctx context.Context, id model.AlbumID) error
}

type albumRepository struct {
    db *sql.DB // 仮に *sql.DB を使用したデータベース接続と仮定
}

// NewAlbumRepository は albumRepository のコンストラクタです
func NewAlbumRepository(db *sql.DB) AlbumRepository {
    return &albumRepository{
        db: db,
    }
}

// GetAll はすべてのアルバムを取得するメソッドです
func (r *albumRepository) GetAll(ctx context.Context) ([]*model.Album, error) {
    rows, err := r.db.QueryContext(ctx, "SELECT id, title, singer_id FROM albums")
    if err != nil {
        return nil, fmt.Errorf("failed to query albums: %w", err)
    }
    defer rows.Close()

    var albums []*model.Album
    for rows.Next() {
        var album model.Album
        if err := rows.Scan(&album.ID, &album.Title, &album.SingerID); err != nil {
            return nil, fmt.Errorf("failed to scan album: %w", err)
        }
        albums = append(albums, &album)
    }

    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("rows iteration error: %w", err)
    }

    return albums, nil
}

// Get は指定したIDのアルバムを取得するメソッドです
func (r *albumRepository) Get(ctx context.Context, id model.AlbumID) (*model.Album, error) {
    row := r.db.QueryRowContext(ctx, "SELECT id, title, singer_id FROM albums WHERE id = ?", id)
    var album model.Album
    if err := row.Scan(&album.ID, &album.Title, &album.SingerID); err != nil {
        if err == sql.ErrNoRows {
            return nil, fmt.Errorf("album with id %d not found", id)
        }
        return nil, fmt.Errorf("failed to scan album: %w", err)
    }

    return &album, nil
}

// Add は新しいアルバムを追加するメソッドです
func (r *albumRepository) Add(ctx context.Context, album *model.Album) error {
    _, err := r.db.ExecContext(ctx, "INSERT INTO albums (title, singer_id) VALUES (?, ?)", album.Title, album.SingerID)
    if err != nil {
        return fmt.Errorf("failed to insert album: %w", err)
    }
    return nil
}

// Delete は指定したIDのアルバムを削除するメソッドです
func (r *albumRepository) Delete(ctx context.Context, id model.AlbumID) error {
    _, err := r.db.ExecContext(ctx, "DELETE FROM albums WHERE id = ?", id)
    if err != nil {
        return fmt.Errorf("failed to delete album: %w", err)
    }
    return nil
}
