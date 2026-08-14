package videos_mobile

type VidesoRepo interface {
}

type VidesoRepoImpl struct {
}

func NewVidesoRepoImpl() *VidesoRepoImpl {
	return &VidesoRepoImpl{}
}
