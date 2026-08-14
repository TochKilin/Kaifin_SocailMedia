package videos

type VidesoRepo interface {
}

type VidesoRepoImpl struct {
}

func NewVidesoRepoImpl() *VidesoRepoImpl {
	return &VidesoRepoImpl{}
}
