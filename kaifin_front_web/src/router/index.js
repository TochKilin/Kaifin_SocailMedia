import { createRouter, createWebHistory } from "vue-router"
import Register from "../components/frontend/auth/Register.vue"
import Login from "@/components/frontend/auth/Login.vue"
import Feed from "@/components/frontend/feed/feed.vue"
import Home from "@/components/frontend/home/home.vue"
import MainProfileContainer from "@/components/frontend/social_profile/MainProfileContainer.vue"
import HomeCourse from "@/components/frontend/course/home/HomeCourse.vue"
import EducationalGames from "@/components/frontend/educational_game/EducationalGames.vue"
// import Home from "@/components/frontend/aticle/home/Home.vue"
import Article from "@/components/frontend/article/Article.vue"
import MyLearning from "@/components/frontend/my_learning/MyLearning.vue"
import Music from "@/components/frontend/music/Music.vue"
// import Quotes from "@/components/frontend/quotes/quotes.vue"
import AboutUs from "@/components/frontend/about_us/AboutUs.vue"
import Contact from "@/components/frontend/contact/Contact.vue"
import ProfileInstructor from "@/components/frontend/home/ProfileInstructor/ProfileInstructor.vue"
import CreateCourse from "@/components/frontend/home/ProfileInstructor/CreateCourse.vue"
import HoleUP from "@/components/frontend/educational_game/HoleUP.vue"
import AddCardCourse from "@/components/frontend/course/AddCardCourse.vue"
import ShoppingCartCourse from "@/components/frontend/course/home/ShoppingCartCourse.vue"
import QrCodeCourse from "@/components/frontend/course/home/QrCodeCourse.vue"
import CourseDetail from "@/components/frontend/course/home/CourseDetail.vue"
import MoreCourseInstructor from "@/components/frontend/course/home/MoreCourseInstructor.vue"
import PaymentMethod from "@/components/frontend/course/home/PaymentMethod.vue"
import About from "@/components/frontend/about_us/About.vue"
import Quotes from "@/components/frontend/quotes/Quotes.vue"
import GroupDetail from "@/components/frontend/sidebar_left_feed/GroupDetail.vue"
// import Quotes from "@/components/frontend/quotes/quotes.vue"



const router = createRouter({

    history:createWebHistory(),

    routes:[
        {
            path:"/",
            redirect:"/login"
        },

        {
            path:"/login",
            component:Login
        },

        {
            path:"/register",
            component:Register
        },

        {
         path:"/feed",
         component:Feed
        },

        {
            path:"/home",
            component:Home
        },
         // ===== PROFILE ROUTES =====
        {
            path: "/profile/:id",
            component: MainProfileContainer,
        },
        {
            path: "/course",
            component: HomeCourse,
        },
        {
            path: "/game",
            component: EducationalGames,
        },
        {
            path: "/article",
            component: Article,
        },
        {
            path: "/my-learning",
            component: MyLearning,
        },
        // {
        //     path: "/music",
        //     component: Music,
        // },
        {
            path: "/quote",
            component: Quotes,
        },
        {
            path: "/about-us",
            component: About,
        },
        {
            path: "/contact",
            component: Contact,
        },
        {
            path: "/profile-instructor/:id",
            name: "ProfileInstructor",
            component: ProfileInstructor,
        },
        {
            path: "/create-course",
            component: CreateCourse,
        },
        {
            path: "/game/:id",
            name: "HoleUP",
            component: HoleUP,
        },
        {
        path: '/course/:id',
        name: 'AddCardCourse',
        component: AddCardCourse
        },
        {
        path: '/shopping-cart', 
        name: 'ShoppingCartCourse',
        component: ShoppingCartCourse
        },
        {
        path: '/qrcode-course',
        name: 'QrCodeCourse', // <-- This name must match exactly what you pass to router.push()
        component: QrCodeCourse
        },
        {
        path: '/course-detail/:id',
        name: 'CourseDetail',
        component: CourseDetail
        },
        {
            path: '/more-course-instructor',
            name: 'MoreCourseInstructor',
            component: MoreCourseInstructor
        },
        {
            path: '/payment-method',
            name: 'PaymentMethod',
            component: PaymentMethod
        },
        {
            path: '/group/:id', 
            name: 'GroupDetail', 
            component: GroupDetail
        }




    ]

})


export default router