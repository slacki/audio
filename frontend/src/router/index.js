import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'

Vue.use(VueRouter)

const routes = [
    {
        path: '/',
        name: 'Home',
        component: Home
    },
    {
        path: '/u/:id',
        name: 'UploadedFile',
        component: () => import('../views/Uploaded.vue')
    },
    {
        path: '/about',
        name: 'About',
        component: () => import('../views/About.vue')
    },
    {
        path: '/privacy',
        name: 'Privacy Policy',
        component: () => import('../views/Privacy.vue')
    },
    {
        path: '/terms',
        name: 'Terms',
        component: () => import('../views/Terms.vue')
    },
    {
        path: '/contact',
        name: 'Contact',
        component: () => import('../views/Contact.vue')
    },
]

const router = new VueRouter({
    mode: 'history',
    routes,
});

export default router;