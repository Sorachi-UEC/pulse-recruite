package mysqldb

import (
    "context"
    "database/sql"

    "github.com/pulse227/server-recruit-challenge-sample/model"
    "github.com/pulse227/server-recruit-challenge-sample/repository"
)

type albumRepository struct {
    db *sql.DB
}

func NewAlbumRepository(db *sql.DB) repository.AlbumRepository {
    return &albumRepository{db: db}
}

func (r *albumRepository) GetAll(ctx context.Context) ([]*model.Album, error) {
    rows, err := r.db.QueryContext(ctx, "SELECT id, title, singer_id FROM albums")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var albums []*model.Album
    for rows.Next() {
        album := &model.Album{}
        if err := rows.Scan(&album.ID, &album.Title, &album.SingerID); err != nil {
            return nil, err
        }
        albums = append(albums, album)
    }
    return albums, nil
}

func (r *albumRepository) Get(ctx context.Context, id model.AlbumID) (*model.Album, error) {
    album := &model.Album{}
    err := r.db.QueryRowContext(ctx, "SELECT id, title, singer_id FROM albums WHERE id = ?", id).
        Scan(&album.ID, &album.Title, &album.SingerID)
    if err != nil {
        return nil, err
    }
    return album, nil
}

func (r *albumRepository) Add(ctx context.Context, album *model.Album) error {
    _, err := r.db.ExecContext(ctx, "INSERT INTO albums (id, title, singer_id) VALUES (?, ?, ?)",
        album.ID, album.Title, album.SingerID)
    return err
}

func (r *albumRepository) Delete(ctx context.Context, id model.AlbumID) error {
    _, err := r.db.ExecContext(ctx, "DELETE FROM albums WHERE id = ?", id)
    return err
}