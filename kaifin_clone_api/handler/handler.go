package handler

import (
	//internal package
	//community package

	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	"kaifin_clone_api/internal/admin/auth"
	"kaifin_clone_api/internal/admin/user"
	"kaifin_clone_api/internal/admin/websocket"
	bookmark_mobile "kaifin_clone_api/internal/front/mobile/bookmark"
	comments_mobile "kaifin_clone_api/internal/front/mobile/comments"
	follower_mobile "kaifin_clone_api/internal/front/mobile/follower"
	likes_mobile "kaifin_clone_api/internal/front/mobile/likes"
	mobile_login "kaifin_clone_api/internal/front/mobile/mobile_login"
	post_mobile "kaifin_clone_api/internal/front/mobile/post"
	profile_mobile "kaifin_clone_api/internal/front/mobile/profile"
	story_mobile "kaifin_clone_api/internal/front/mobile/story"
	addcard "kaifin_clone_api/internal/front/web/add_card"
	"kaifin_clone_api/internal/front/web/article"
	articlecomment "kaifin_clone_api/internal/front/web/article_comment"
	"kaifin_clone_api/internal/front/web/bookmark"
	"kaifin_clone_api/internal/front/web/chat"
	"kaifin_clone_api/internal/front/web/comments"
	"kaifin_clone_api/internal/front/web/communities"
	"kaifin_clone_api/internal/front/web/course"
	courseinroll "kaifin_clone_api/internal/front/web/course_inroll"
	"kaifin_clone_api/internal/front/web/follower"
	"kaifin_clone_api/internal/front/web/likes"
	"kaifin_clone_api/internal/front/web/menu"
	"kaifin_clone_api/internal/front/web/music"
	"kaifin_clone_api/internal/front/web/mystickerset"
	"kaifin_clone_api/internal/front/web/notification"
	notificationbell "kaifin_clone_api/internal/front/web/notification_bell"
	"kaifin_clone_api/internal/front/web/playlist"
	"kaifin_clone_api/internal/front/web/post"
	"kaifin_clone_api/internal/front/web/profile"
	"kaifin_clone_api/internal/front/web/quote"
	quotereaction "kaifin_clone_api/internal/front/web/quote_reaction"
	quoteshare "kaifin_clone_api/internal/front/web/quote_share"
	quoteview "kaifin_clone_api/internal/front/web/quote_view"
	"kaifin_clone_api/internal/front/web/songs"
	"kaifin_clone_api/internal/front/web/sponsor"
	"kaifin_clone_api/internal/front/web/stickers"
	"kaifin_clone_api/internal/front/web/story"
	"kaifin_clone_api/internal/front/web/storyreaction"
	streaklevel "kaifin_clone_api/internal/front/web/streak_level"
	"kaifin_clone_api/internal/front/web/template"
	"kaifin_clone_api/internal/front/web/user_login"
	userregister "kaifin_clone_api/internal/front/web/user_register"
	"kaifin_clone_api/pkg/middleware"
)

type ServiceHandler struct {
	Admin  *AdminService
	Front  *FrontService
	Mobile *MobileService
}

func NewServiceHandler(app *fiber.App, db *sqlx.DB, rdb *redis.Client, ws *websocket.WebSocketManager) *ServiceHandler {
	middleware.NewJwtMinddleWare(app, db, rdb)
	return &ServiceHandler{
		Front:  NewFrontService(app, db, rdb, ws),
		Admin:  NewAdminService(app, db, rdb, ws),
		Mobile: NewMobileSerivice(app, db, rdb, ws),
	}
}

type AdminService struct {
	Auth      *auth.AuthRoute
	userRoute *user.UserRoute
}

func NewAdminService(app *fiber.App, dbpool *sqlx.DB, rdb *redis.Client, ws *websocket.WebSocketManager) *AdminService {
	r := auth.NewAuthRoute(app, dbpool, rdb, ws)
	us := user.NewUserRoute(app, dbpool, rdb, ws)
	return &AdminService{
		Auth:      r,
		userRoute: us,
	}
}

type FrontService struct {
	// Web front
	AuthUser       *user_login.AuthUserRouteImpl
	Register       *userregister.RegisterRouteImpl
	Post           *post.PostRouteImpl
	Menu           *menu.MenuRouteImpl
	Profile        *profile.ProfileRouteImpl
	Commu          *communities.CommunitiesRouteImpl
	Sto            *story.StoryRouteImgl
	li             *likes.LikesRouteImpl
	comm           *comments.CommentsRouteImpl
	Bookmark       *bookmark.BookMarkRouteImpl
	Follower       *follower.FollowersRouteImpl
	Sticker        *stickers.StickersRouteImpl
	StoryReaction  *storyreaction.StoryReactionRouteImpl
	MySticker      *mystickerset.MysetStickerRouteImpl
	Template       *template.TemplateRouteImpl
	Music          *music.PostMusicRouteImpl
	Song           *songs.SongRouteImpl
	Playlit        *playlist.PlaylistRouteImpl
	Quote          *quote.QuoteRoute
	QuoteReaction  *quotereaction.QuoteReactionRoute
	ViewsQuote     *quoteview.QuoteViewRoutel
	ShareQuote     *quoteshare.QuoteShareRoutel
	Article        *article.ArticlesRouteImpl
	Notification   *notification.NotificationsRouteImpl
	ArticleComment *articlecomment.CommentsRouteImpl

	Course       *course.CourseRouteImpl
	StreakLevel  *streaklevel.LevelRouteImpl
	Sponsor      *sponsor.SponsorsRouteImpl
	CourseInroll *courseinroll.EnrollmentRouteImpl
	AddCard      *addcard.CartRouteImpl
	Chat         *chat.ChatsRouteImpl
	Noti         *notificationbell.NotiRouteImpl
}

func NewFrontService(app *fiber.App, dbpool *sqlx.DB, rdb *redis.Client, ws *websocket.WebSocketManager) *FrontService {

	regRoute := userregister.NewRegisterRouteImpl(app, dbpool, rdb, ws)
	uRoute := user_login.NewAuthUserRouteImpl(app, dbpool, rdb, ws)
	profileRoute := profile.NewProfileRouteImpl(app, dbpool)
	menuRoute := menu.NewMenuRouteImpl(app, dbpool, rdb, ws)
	postRoute := post.NewPostRouteImpl(app, dbpool, rdb, ws)
	communities := communities.NewCommunitiesRoute(app, dbpool)
	story := story.NewStoryRouteImpl(app, dbpool, rdb, ws)
	likes := likes.NewLikesRouteImpl(app, dbpool, rdb)
	comments := comments.NewCommentsRouteImpl(app, dbpool)
	bookmark := bookmark.NewBookMarkRouteImpl(app, dbpool, rdb, ws)
	followers := follower.NewFollowersRouteImpl(app, dbpool, ws)
	sticker := stickers.NewStickersRouteImpl(app, dbpool)
	storyreaction := storyreaction.NewStoryReactionRouteImpl(app, dbpool)
	mysetsticker := mystickerset.NewMysetStickerRouteImpl(app, dbpool, rdb, ws)
	templ := template.NewTemplateRouteImpl(app, dbpool)
	music := music.NewPostMusicRouteImpl(app, dbpool, rdb, ws)
	songs := songs.NewSongRouteImpl(app, dbpool, rdb, ws)
	playlist := playlist.NewPlaylistRouteImpl(app, dbpool, rdb, ws)
	quote := quote.NewQuoteRoute(app, dbpool, rdb, ws)
	quoteReact := quotereaction.NewQuoteReactionRoute(app, dbpool, rdb, ws)
	quoteview := quoteview.NewQuoteViewRoute(app, dbpool)
	quoteshare := quoteshare.NewQuoteShareRoute(app, dbpool)
	article := article.NewArticlesRouteImpl(app, dbpool)
	notification := notification.NewNotificationsRouteImpl(app, dbpool, ws)
	articlecomment := articlecomment.NewCommentsRouteImpl(app, dbpool)
	course := course.NewCourseRouteImpl(app, dbpool, rdb, ws)
	streaklevel := streaklevel.NewLevelRouteImpl(app, dbpool)
	sponsor := sponsor.NewSponsorsRouteImpl(app, dbpool)
	courseenroll := courseinroll.NewCourseEnrollmentRoute(app, dbpool)
	addcard := addcard.NewCartRouteImpl(app, dbpool)
	chat := chat.NewChatsRouteImpl(app, dbpool, ws)
	notifi := notificationbell.NewNotiRoute(app, dbpool, ws)

	return &FrontService{
		AuthUser:       uRoute,
		Register:       regRoute,
		Post:           postRoute,
		Menu:           menuRoute,
		Profile:        profileRoute,
		Commu:          communities,
		Sto:            story,
		li:             likes,
		comm:           comments,
		Bookmark:       bookmark,
		Follower:       followers,
		Sticker:        sticker,
		StoryReaction:  storyreaction,
		MySticker:      mysetsticker,
		Template:       templ,
		Music:          music,
		Song:           songs,
		Playlit:        playlist,
		Quote:          quote,
		QuoteReaction:  quoteReact,
		ViewsQuote:     quoteview,
		ShareQuote:     quoteshare,
		Article:        article,
		Notification:   notification,
		ArticleComment: articlecomment,
		Course:         course,
		StreakLevel:    streaklevel,
		Sponsor:        sponsor,
		CourseInroll:   courseenroll,

		AddCard: addcard,
		Chat:    chat,
		Noti:    notifi,
	}
}

type MobileService struct {
	loginmobile    *mobile_login.AuthUserMobileRouteImpl
	postmobile     *post_mobile.PostMobileRouteImpl
	cmtmobile      *comments_mobile.CommentsMobileRouteImpl
	likesmobile    *likes_mobile.LikesMobileRouteImpl
	bookmarkmobile *bookmark_mobile.BookMarkMobileRouteImpl
	followmobile   *follower_mobile.FollowersMobileRouteImpl
	profilemobile  *profile_mobile.ProfileMobileRouteImpl
	storymobile    *story_mobile.StoryMobileRouteImgl
}

func NewMobileSerivice(app *fiber.App, db *sqlx.DB, rdb *redis.Client, ws *websocket.WebSocketManager) *MobileService {
	loginmobile := mobile_login.NewAuthUserMobileRouteImpl(app, db, rdb, ws)
	postMobile := post_mobile.NewPostMobilrRouteImpl(app, db, rdb, ws)
	cmtMobile := comments_mobile.NewCommentsMobileRouteImpl(app, db, rdb, ws)
	likesMobile := likes_mobile.NewLikesMobileRouteImpl(app, db, rdb)
	bookmarkMobile := bookmark_mobile.NewBookMarkMobileRouteImpl(app, db, rdb, ws)
	followMobile := follower_mobile.NewFollowersMobileRouteImpl(app, db)
	profileMobile := profile_mobile.NewProfileMobileRouteImpl(app, db)
	storyMobile := story_mobile.NewStoryMobileRouteImpl(app, db, rdb, ws)

	return &MobileService{
		loginmobile:    loginmobile,
		postmobile:     postMobile,
		cmtmobile:      cmtMobile,
		likesmobile:    likesMobile,
		bookmarkmobile: bookmarkMobile,
		followmobile:   followMobile,
		profilemobile:  profileMobile,
		storymobile:    storyMobile,
	}
}
