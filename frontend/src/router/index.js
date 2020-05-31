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
        path: '/faq',
        name: 'faq',
        component: () => import('../views/FAQ.vue')
    },
    {
        path: '/terms-and-privacy',
        name: 'terms',
        component: () => import('../views/TermsAndPrivacy.vue')
    },
    {
        path: '/abuse',
        name: 'abuse',
        component: () => import('../views/Abuse.vue')
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