package videos_mobile

type VideosService interface {
}

type VideosServiceImpl struct {
}

func NewVideosServiceImpl() *VideosServiceImpl {
	return &VideosServiceImpl{}
}
