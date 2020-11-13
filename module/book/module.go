package book

import "go.uber.org/fx"

func LoadBook() fx.Option {
	return fx.Options(
		fx.Provide(NewFindBook),
		fx.Provide(NewSQLBookRepository),
		fx.Provide(NewOtherServiceBookRepository),
		fx.Provide(NewRedisBookRepository),
		fx.Provide(NewRedundantBookRepository),
	)
}

//Example only to speed up wiring
func LoadRedundantBookParam() fx.Option {
	return fx.Options(

		fx.Provide(func(repo *RedundantBookRepository) IBookRepository {
			return repo
		}),

		fx.Provide(func(
			otherService *OtherServiceBookRepository,
			sqlRepo *SQLBookRepository,
			redisRepo *RedisBookRepository,
		) *RedundantBookRepositoryParam {
			return &RedundantBookRepositoryParam{
				Repos: []IBookRepository{
					redisRepo,
					sqlRepo,
					otherService,
				},
			}
		}),
	)
}
