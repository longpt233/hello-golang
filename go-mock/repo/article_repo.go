package repo

type ArticleRepo struct {
}

func (ArticleRepo) Get(a int) string {
	return "hello go mock"
}
