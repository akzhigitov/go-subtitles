import Vue from 'vue';
import VueRouter from 'vue-router';

Vue.use(VueRouter);

const routes = [
    {
        path: '/',
        name: 'Subtitle',
        component: Subtitle
    },
    {
        path: '/error',
        name: 'Error',
        component: Error,
        meta: {
            allowAnonymous: true
        }
    },
]

const router = new VueRouter({
    routes,
});


export default router;
