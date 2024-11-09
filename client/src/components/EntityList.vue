<template>
  <div>
    <h2>{{ getEntityTitle(entityType) }}</h2>
    <table v-if="items.length" border="1" cellpadding="10">
      <thead>
      <tr>
        <th v-for="header in getTableHeaders(entityType)" :key="header">{{ header }}</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="item in items" :key="item._id">
        <td v-if="entityType === 'clients'">{{ item._id }}</td>
        <td v-if="entityType === 'clients'">{{ item.name }}</td>
        <td v-if="entityType === 'clients'">{{ item.phone }}</td>
        <td v-if="entityType === 'clients'">{{ item.gender }}</td>
        <td v-if="entityType === 'clients'">{{ item.birth_date }}</td>
        <td v-if="entityType === 'clients'">{{ item.created_at }}</td>
        <td v-if="entityType === 'clients'">{{ item.updated_at }}</td>
        <td v-if="entityType === 'clients'">{{ item.password }}</td>
        <td v-if="entityType === 'clients'">
          <a v-if="item.picture_uri" :href="item.picture_uri" target="_blank">{{ item.picture_uri }}</a>
        </td>
        <td v-if="entityType === 'clients'">
          <ul>
            <li v-for="classId in item.classes" :key="classId">{{ classId }}</li>
          </ul>
        </td>

        <td v-if="entityType === 'trainers'">{{ item._id }}</td>
        <td v-if="entityType === 'trainers'">{{ item.name }}</td>
        <td v-if="entityType === 'trainers'">{{ item.phone }}</td>
        <td v-if="entityType === 'trainers'">{{ item.gender }}</td>
        <td v-if="entityType === 'trainers'">{{ item.birth_date }}</td>
        <td v-if="entityType === 'trainers'">{{ item.created_at }}</td>
        <td v-if="entityType === 'trainers'">{{ item.updated_at }}</td>
        <td v-if="entityType === 'trainers'">{{ item.studio_id }}</td>
        <td v-if="entityType === 'trainers'">
          <a v-if="item.picture_uri" :href="item.picture_uri" target="_blank">Фото</a>
        </td>
        <td v-if="entityType === 'trainers'">
          <ul>
            <li v-for="classId in item.classes" :key="classId">{{ classId }}</li>
          </ul>
        </td>

        <td v-if="entityType === 'classes'">{{ item._id }}</td>
        <td v-if="entityType === 'classes'">{{ item.class_name }}</td>
        <td v-if="entityType === 'classes'">{{ item.time }}</td>
        <td v-if="entityType === 'classes'">{{ item.studio_id }}</td>
        <td v-if="entityType === 'classes'">{{ item.trainer_id }}</td>
        <td v-if="entityType === 'classes'">
          <ul>
            <li v-for="clientId in item.clients" :key="clientId">{{ clientId }}</li>
          </ul>
        </td>

        <td v-if="entityType === 'studios'">{{ item._id }}</td>
        <td v-if="entityType === 'studios'">{{ item.address }}</td>
        <td v-if="entityType === 'studios'">
          <ul>
            <li v-for="classId in item.classes" :key="classId">{{ classId }}</li>
          </ul>
        </td>
        <td v-if="entityType === 'studios'">
          <ul>
            <li v-for="trainerId in item.trainers" :key="trainerId">{{ trainerId }}</li>
          </ul>
        </td>
      </tr>
      </tbody>
    </table>
    <p v-else>Нет данных для отображения.</p>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, watch } from 'vue';
import { useRoute } from 'vue-router';

const route = useRoute();
const entityType = ref(route.params.entityType as 'clients' | 'trainers' | 'classes' | 'studios');
const items = ref([]);

async function loadData() {
  try {
    const response = await fetch(`address/${entityType.value}`); // заменить на правильный адрес
    const data = await response.json();
    items.value = data || [];
  } catch (error) {
    console.error("Ошибка при загрузке данных:", error);
  }
}

onMounted(loadData);

watch(() => route.params.entityType, (newType) => {
  entityType.value = newType as 'clients' | 'trainers' | 'classes' | 'studios';
  loadData();
});

function getEntityTitle(type: string) {
  switch (type) {
    case 'clients': return 'Клиенты';
    case 'trainers': return 'Тренеры';
    case 'classes': return 'Занятия';
    case 'studios': return 'Студии';
    default: return 'Список';
  }
}

// Функция для получения заголовков таблицы
function getTableHeaders(type: string) {
  switch (type) {
    case 'clients': return ['ID', 'Имя', 'Телефон', 'Гендер', 'Дата рождения', 'Создан', 'Обновлен', 'Пароль', 'Фото', 'Занятия'];
    case 'trainers': return ['ID', 'Имя', 'Телефон', 'Гендер', 'Дата рождения', 'Создан', 'Обновлен', 'Студия', 'Фото', 'Занятия'];
    case 'classes': return ['ID', 'Название занятия', 'Время', 'Студия', 'Тренер', 'Клиенты'];
    case 'studios': return ['ID', 'Адрес', 'Занятия', 'Тренеры'];
    default: return [];
  }
}
</script>

<style scoped>
h2 {
  margin-top: 1rem;
}

table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 1rem;
}

th, td {
  padding: 0.5rem;
  border: 1px solid #888;
  text-align: center;
}

a {
  color: #0066cc;
  text-decoration: none;
}

a:hover {
  text-decoration: underline;
}

ul {
  padding: 0;
  margin: 0;
}

li {
  list-style-type: none;
  margin: 0.5rem 0;
}
</style>
