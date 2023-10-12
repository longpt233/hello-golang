package usecase

//go:generate mockgen -destination=../mocks/mock_article_usecase.go -package=mocks github.com/longpt233/go-mock/usecase ArticleRepo
type ArticleRepo interface {
	Get(int) string
}

type ArticleUc struct {
	Repo ArticleRepo
}

func (u *ArticleUc) Do() string {
	out_repo := u.Repo.Get(123)
	out_repo = out_repo + "do" // thuc hien logic
	return out_repo
}
