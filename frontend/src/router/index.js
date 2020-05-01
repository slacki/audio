import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'

Vue.use(VueRouter)

const routes = [
    {
        path: '/',
        name: 'home',
        component: Home
    },
    {
        path: '/about',
        name: 'about',
        component: () => import('../views/About.vue')
    },
    {
        path: '/privacy',
        name: 'privacyPolicy',
        component: () => import('../views/Privacy.vue')
    },
    {
        path: '/terms',
        name: 'terms',
        component: () => import('../views/Terms.vue')
    },
    {
        path: '/contact',
        name: 'contact',
        component: () => import('../views/Contact.vue')
    },
    {
        path: '/:id',
        name: 'uploadedFile',
        component: () => import('../views/Uploaded.vue')
    },
]

const router = new VueRouter({
    mode: 'history',
    routes,
});

export default router;