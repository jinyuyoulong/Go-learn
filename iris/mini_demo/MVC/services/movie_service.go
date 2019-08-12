package services

import (
	datamodels2 "github.com/jinyuyoulong/Go-learn/iris/mini_demo/MVC/datamodels"
	repositories2 "github.com/jinyuyoulong/Go-learn/iris/mini_demo/MVC/repositories"
)

// MovieService处理电影数据模型的一些CRUID操作。
//这取决于影片库的动作。
//这是将数据源与更高级别的组件分离。
//因此，不同的存储库类型可以使用相同的逻辑，而无需任何更改。
//它是一个界面，它在任何地方都被用作界面
//因为我们可能需要在将来更改或尝试实验性的不同域逻辑。
type MovieService interface {
	GetAll() []datamodels2.Movie
	GetByID(id int64) (datamodels2.Movie, bool)
	DeleteByID(id int64) bool
	UpdatePosterAndGenreByID(id int64, poster string, genre string) (datamodels2.Movie, error)
}
type movieService struct {
	repo repositories2.MovieRepository
}

// NewMovieService返回默认 movie service.
func NewMovieService(repo repositories2.MovieRepository) MovieService {
	return &movieService{
		repo: repo,
	}
}

// GetAll 获取所有的movie.
func (s *movieService) GetAll() []datamodels2.Movie {
	return s.repo.SelectMany(func(_ datamodels2.Movie) bool {
		return true
	}, -1)
}

// GetByID 根据其ID返回一行。
func (s *movieService) GetByID(id int64) (datamodels2.Movie, bool) {
	return s.repo.Select(func(m datamodels2.Movie) bool {
		return m.ID == id
	})
}

// UpdatePosterAndGenreByID更新电影的海报和流派。
func (s *movieService) UpdatePosterAndGenreByID(id int64, poster string, genre string) (datamodels2.Movie, error) {
	// update the movie and return it.
	return s.repo.InsertOrUpdate(datamodels2.Movie{
		ID:     id,
		Poster: poster,
		Genre:  genre,
	})
}

// DeleteByID按ID删除电影。
//
//如果删除则返回true，否则返回false。
func (s *movieService) DeleteByID(id int64) bool {
	return s.repo.Delete(func(m datamodels2.Movie) bool {
		return m.ID == id
	}, 1)
}
