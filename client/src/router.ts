import {createRouter, createWebHistory} from 'vue-router';
import type { RouteRecordRaw } from 'vue-router';
import DataView from '@/views/DataView.vue';
import EntityList from "@/components/EntityList.vue";
import Auth from "@/components/Auth.vue";
import AdminPage from "@/components/AdminPage.vue";
import Home from "@/components/Home.vue";
import EntityCard from "@/components/EntityCard.vue";

const routes: Array<RouteRecordRaw> = [
    {
        path: '/',
        name: 'Auth',
        component: Auth,
        meta: { title: 'Login - Youga Places' },
    },
    {
        path: '/admin',
        name: 'AdminPage',
        component: AdminPage,
        meta: { title: 'Admin Dashboard - Youga Places' },
    },
    {
        path: '/admin/data',
        name: 'DataView',
        component: DataView,
        meta: { title: 'Data View - Admin - Youga Places' },
    },
    {
        path: '/:entityType/:id',
        name: 'EntityCard',
        component: EntityCard,
        props: true,
        meta: {
            title: (route: any) => `${route.params.entityType} - Admin - Youga Places`,
        },
    },
    {
        path: '/home',
        name: 'Home',
        component: Home,
        meta: { title: 'Home - Youga Places' },
    },
    {
        path: '/admin/data/:entityType',
        name: 'EntityList',
        component: EntityList,
        props: true,
        meta: {
            title: (route: any) => `${route.params.entityType} List - Admin - Youga Places`,
        },
    },
];


const router = createRouter({
    history: createWebHistory(),
    routes,
});

router.beforeEach((to, from, next) => {
    const defaultTitle = 'Youga Places';
    const title =
        typeof to.meta.title === 'function'
            ? to.meta.title(to)
            : to.meta.title;

    document.title = title || defaultTitle;
    next();
});

export default router;
