package videos

type VideosService interface {
}

type VideosServiceImpl struct {
}

func NewVideosServiceImpl() *VideosServiceImpl {
	return &VideosServiceImpl{}
}
