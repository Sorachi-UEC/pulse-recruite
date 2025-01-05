package service

import (
    "context"
	"fmt"

    "github.com/pulse227/server-recruit-challenge-sample/model"
    "github.com/pulse227/server-recruit-challenge-sample/repository"
)

type AlbumService interface {
    GetAlbumListService(ctx context.Context) ([]*model.Album, error)
    GetAlbumService(ctx context.Context, albumID model.AlbumID) (*model.Album, error)
    PostAlbumService(ctx context.Context, album *model.Album) error
    DeleteAlbumService(ctx context.Context, albumID model.AlbumID) error
}

type albumService struct {
    albumRepository repository.AlbumRepository
	singerRepository repository.SingerRepository //課題4
}

var _ AlbumService = (*albumService)(nil)

//SingerRepositoryの追加
func NewAlbumService(albumRepository repository.AlbumRepository, singerRepository repository.SingerRepository) AlbumService {
    return &albumService{
		albumRepository: albumRepository,
		singerRepository: singerRepository,
	}
}

func (s *albumService) GetAlbumListService(ctx context.Context) ([]*model.Album, error) {
    //return s.albumRepository.GetAll(ctx) //課題3はこれ

    albums, err := s.albumRepository.GetAll(ctx)
    if err != nil {
        return nil, err
    }

    // 各アルバムに対応する歌手情報を埋め込む
    for _, album := range albums {
        singer, err := s.singerRepository.Get(ctx, model.SingerID(album.SingerID))
        if err != nil {
            return nil, fmt.Errorf("failed to get singer: %w", err)
        }
        album.Singer = singer
    }

    return albums, nil
}

func (s *albumService) GetAlbumService(ctx context.Context, albumID model.AlbumID) (*model.Album, error) {
    //return s.albumRepository.Get(ctx, albumID)

    album, err := s.albumRepository.Get(ctx, albumID)
    if err != nil {
        return nil, err
    }

    // アルバムに対応する歌手情報を埋め込む
    singer, err := s.singerRepository.Get(ctx, model.SingerID(album.SingerID))
    if err != nil {
        return nil, fmt.Errorf("failed to get singer: %w", err)
    }
    album.Singer = singer

    return album, nil
}

func (s *albumService) PostAlbumService(ctx context.Context, album *model.Album) error {
    return s.albumRepository.Add(ctx, album)


}

func (s *albumService) DeleteAlbumService(ctx context.Context, albumID model.AlbumID) error {
    return s.albumRepository.Delete(ctx, albumID)
}