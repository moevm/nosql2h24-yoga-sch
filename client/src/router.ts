import { createRouter, createWebHistory } from 'vue-router';
import DataView from '@/views/DataView.vue';
import EntityList from "@/components/EntityList.vue";
import Auth from "@/components/Auth.vue";
import AdminPage from "@/components/AdminPage.vue";

const routes = [
    {
        path: '/',
        name: 'Auth',
        component: Auth,
    },
    {
        path: '/admin',
        name: 'AdminPage',
        component: AdminPage,
    },
    {
        path: '/admin/data',
        name: 'DataView',
        component: DataView,
    },
    {   path: '/admin/data/:entityType',
        name: 'EntityList',
        component: EntityList,
        props: true,
    },
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

export default router;