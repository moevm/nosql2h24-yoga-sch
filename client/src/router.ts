import { createRouter, createWebHistory } from 'vue-router';
import Home from '@/components/Home.vue';
import DataView from '@/views/DataView.vue';
import EntityList from "@/components/EntityList.vue";

const routes = [
    {
        path: '/',
        name: 'Home',
        component: Home,
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