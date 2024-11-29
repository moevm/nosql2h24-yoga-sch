import { createRouter, createWebHistory } from 'vue-router';
import DataView from '@/views/DataView.vue';
import EntityList from "@/components/EntityList.vue";
import Auth from "@/components/Auth.vue";

const routes = [
    {
        path: '/',
        name: 'Auth',
        component: Auth,
    },
    {
        path: '/data',
        name: 'DataView',
        component: DataView,
    },
    {   path: '/data/:entityType',
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