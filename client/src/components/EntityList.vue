<template xmlns="http://www.w3.org/1999/html">
  <div class="table-container">
    <h2>{{ (items?.length || 0) + " " + entityType }}</h2>

    <div class="filters">
      <h3>Фильтры поиска</h3>
      <div
          v-for="header in getColumnConfig(entityType)"
          :key="header.key"
          class="filter-group">
      <label :for="'filter-' + header.key">
          {{ header.label }}
        </label>

        <select v-if="header.label === 'Gender'" :id="'filter-' + header.key" v-model="filters[header.key]">
          <option selected disabled hidden value="">All</option>
          <option value="MALE">Male</option>
          <option value="FEMALE">Female</option>
        </select>

        <input v-else-if="header.isDate && header.isTime" type="datetime-local"
               :id="'filter-' + header.key" v-model="filters[header.key]" />
        <input v-else-if="header.isDate && !header.isTime" type="date"
               :id="'filter-' + header.key" v-model="filters[header.key]" />

        <input v-else type="text"
               :id="'filter-' + header.key"
               :placeholder="'Введите ' + header.label"
               v-model="filters[header.key]" />
      </div>
      <button class="apply-filters" @click="applyFilters">Apply</button>
    </div>

    <table v-if="items.length" border="1" cellpadding="10">
      <thead>
      <tr>
        <template v-for="header in getColumnConfig(entityType)" :key="header.label">
          <th v-if="header.label !== 'Password'">{{ header.label }}</th>
        </template>
      </tr>
      </thead>
      <tbody>
      <tr v-for="item in items" :key="item['id']">
        <template v-for="column in getColumnConfig(entityType)" :key="column.key">
          <td v-if="!column.isLink && !column.isList && !column.isDate && !column.isPassword">{{ item[column.key as keyof Entity] }}</td>
          <td v-else-if="column.isLink">
            <a v-if="item[column.key as keyof Entity]" :href="item[column.key as keyof Entity]" target="_blank">
              {{ item[column.key as keyof Entity] }}
            </a>
          </td>
          <td v-else-if="column.isDate && column.isTime">{{ getDateTime(item[column.key as keyof Entity]) }}</td>
          <td v-else-if="column.isDate && !column.isTime">{{ getDate(item[column.key as keyof Entity]) }}</td>
          <td v-else-if="column.isList">
            <ul>
              <li v-for="elem in item[column.key as keyof Entity]" :key="elem">{{ elem }}</li>
            </ul>
          </td>
        </template>
      </tr>
      </tbody>
    </table>
    <p v-else>Нет данных для отображения.</p>

    <button @click="showModal = true" class="add-button">Добавить новый элемент</button>

    <div v-if="showModal" class="modal">
      <div class="modal-content">
        <h3>Добавить новый элемент</h3>
        <form>
          <div v-for="header in getColumnConfig(entityType)" :key="header.label">
            <div v-if="header.isNeeded" class="form-group">
              <label :for="header.key">{{ header.label }}</label>
              <input v-model="formData.phone"
                     v-if="header.label === 'Phone'"
                     type="tel"
                     :id="header.key"
                     placeholder="+7(999)999-9999"
                     required
              />
              <select v-model="formData.gender"
                      v-else-if="header.label === 'Gender'"
                      :id="header.key"
                      required
              >
                <option value="MALE">Male</option>
                <option value="FEMALE">Female</option>
              </select>
              <input v-model="formData.password"
                     v-else-if="header.label === 'Password'"
                     type="password"
                     :id="header.key"
                     required
              />
              <input v-model="formData[header.key]"
                     v-else-if="header.label === 'Time'"
                     type="datetime-local"
                     :id="header.key"
                     required
              />
              <input v-model="formData[header.key]"
                     v-else-if="header.label === 'Birth Date'"
                     type="date"
                     :max="today"
                     :id="header.key"
                     required
              />
              <input v-model="formData[header.key]"
                     v-else-if="!header.isLink && !header.isList && !header.isDate && !header.isTime"
                     :id="header.key"
                     :placeholder="'Enter ' + header.label"
                     required
              />
            </div>
          </div>
          <div class="modal-actions">
            <button type="submit" class="modal-btn-create" @click="addNewItem()">Создать</button>
            <button type="button" class="modal-btn-close" @click="showModal = false">Закрыть</button>
          </div>
        </form>
        <div v-if="errorMessage" class="error-message">
          {{ errorMessage }}
        </div>
      </div>
    </div>

  </div>
</template>

<script setup lang="ts">
import {ref, type Ref, onMounted, watch} from 'vue';
import {useRoute} from 'vue-router';
import {type Class, type Client, type Studio, type Trainer} from "@/types/types";

type Entity = Client | Class | Studio | Trainer
const route = useRoute();
const entityType = ref(route.params.entityType as 'client' | 'trainer' | 'class' | 'studio');
const filters: Ref<Record<string, string>> = ref({});

const items: Ref<Entity[]> = ref([]);
const showModal = ref(false);
const errorMessage = ref("");

const formData = ref<Record<string, any>>({});
const today = new Date().toISOString().split('T')[0];

const URI = `${window.location.protocol}//${window.location.hostname}`


async function loadData() {
  try {
    const response = await fetch(`${URI}/api/v1/${entityType.value}`, {
      method: 'GET'
    });
    const data = await response.json();

    switch (entityType.value) {
      case 'client':
        items.value = data.clients || [];
        break;
      case 'trainer':
        items.value = data.trainers || [];
        break;
      case 'class':
        items.value = data.classes || [];
        break;
      case 'studio':
        items.value = data.studios || [];
        break;
    }

    // console.log(items.value);
  } catch (error) {
    console.error("Ошибка при загрузке данных:", error);
  }
}

onMounted(loadData);

watch(() => route.params.entityType, (newType) => {
  entityType.value = newType as 'client' | 'trainer' | 'class' | 'studio';
  loadData();
});

function getColumnConfig(type: string) {
  switch (type) {
    case 'client':
      return [
        {key: 'id', label: 'ID'},
        {key: 'name', label: 'Name', isNeeded: true},
        {key: 'phone', label: 'Phone', isNeeded: true},
        {key: 'gender', label: 'Gender', isNeeded: true},
        {key: 'birthDate', label: 'Birth Date', isDate: true, isTime: false, isNeeded: true},
        {key: 'createdAt', label: 'Created At', isDate: true, isTime: true},
        {key: 'updatedAt', label: 'Updated At', isDate: true, isTime: true},
        {key: 'password', label: 'Password', isNeeded: true, isPassword: true},
        {key: 'pictureUri', label: 'Picture', isLink: true},
        {key: 'classIds', label: 'Classes', isList: true}
      ];
    case 'trainer':
      return [
        {key: 'id', label: 'ID'},
        {key: 'name', label: 'Name', isNeeded: true},
        {key: 'phone', label: 'Phone', isNeeded: true},
        {key: 'gender', label: 'Gender', isNeeded: true},
        {key: 'birthDate', label: 'Birth Date', isDate: true, isTime: false, isNeeded: true},
        {key: 'createdAt', label: 'Created At', isDate: true, isTime: true},
        {key: 'updatedAt', label: 'Updated At', isDate: true, isTime: true},
        {key: 'studioId', label: 'Studio ID', isNeeded: true},
        {key: 'pictureUri', label: 'Picture', isLink: true},
        {key: 'classIds', label: 'Classes', isList: true}
      ];
    case 'class':
      return [
        {key: 'id', label: 'ID'},
        {key: 'name', label: 'Class Type', isNeeded: true},
        {key: 'time', label: 'Time', isDate: true, isTime: true, isNeeded: true},
        {key: 'studioId', label: 'Studio ID', isNeeded: true},
        {key: 'trainerId', label: 'Trainer ID', isNeeded: true},
        {key: 'clients', label: 'Clients', isList: true}
      ];
    case 'studio':
      return [
        {key: 'id', label: 'ID'},
        {key: 'address', label: 'Address', isNeeded: true},
        {key: 'classes', label: 'Classes', isList: true},
        {key: 'trainers', label: 'Trainers', isList: true}
      ];
    default:
      return [];
  }
}

function getDateTime(date: string) {
  return new Date(date).toLocaleString()
}

function getDate(date: string) {
  return new Date(date).toDateString()
}

const addNewItem = async () => {
  try {
    if (formData.value.birthDate) {
      formData.value.birthDate = new Date(formData.value.birthDate).toISOString();
    }

    if (formData.value.time) {
      formData.value.time = new Date(formData.value.time).toISOString();
    }

    let payload = {
      [entityType.value.toLowerCase()]: formData.value
    };

    console.log(JSON.stringify(payload));

    const response = await fetch(`${URI}/api/v1/${entityType.value.toLowerCase()}`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify(payload)
    });

    if (!response.ok) {
      throw new Error(`${response.statusText}`);
    }

    showModal.value = false;
    formData.value = {};
    await loadData();

    errorMessage.value = "";
    // location.reload();
  } catch (error) {
    if (error instanceof Error) {
      errorMessage.value = `Ошибка при отправке данных: ${error.message}`;
    } else {
      errorMessage.value = "Произошла неизвестная ошибка";
    }
  }
}

const applyFilters = async () => {
  const queryParams = Object.entries(filters.value)
      .filter(([_, value]) => value)
      .map(([key, value]) => `${key}=${encodeURIComponent(value)}`)
      .join('&');

  try {
    const response = await fetch(`${URI}/api/v1/${entityType.value}?${queryParams}`, {
      method: 'GET',
    });
    const data = await response.json();
    items.value = data[entityType.value.toLowerCase()] || [];
  } catch (error) {
    console.error('Ошибка фильтрации:', error);
  }
};

</script>

<style scoped>

.error-message {
  color: #f44336;
  background-color: #ffebee;
  padding: 10px;
  border: 1px solid #f44336;
  border-radius: 4px;
  margin-top: 15px;
  text-align: center;
}

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

.table-container {
  padding-bottom: 60px;
}

.add-button {
  position: fixed;
  bottom: 20px;
  right: 20px;
  padding: 12px 24px;
  background-color: #f44336;
  color: white;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  font-size: 16px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  transition: background-color 0.3s ease, transform 0.2s ease;
  z-index: 100;
}

.add-button:hover {
  background-color: #d32f2f;
  transform: translateY(-3px);
}

.add-button:active {
  background-color: #b71c1c;
  transform: translateY(0);
}

.modal {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 200;
}

.modal-content {
  background-color: white;
  padding: 20px;
  border-radius: 8px;
  width: 90%;
  max-width: 500px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  overflow-y: auto;
  max-height: 90vh;
}

.modal-content h3 {
  text-align: center;
  font-size: 20px;
  margin-top: 0;
  margin-bottom: 20px;
}

.form-group {
  margin-bottom: 15px;
}

.form-group label {
  font-size: 14px;
  margin-bottom: 5px;
  display: block;
}

.modal-content input,
.modal-content select {
  width: 100%;
  padding: 10px;
  margin-top: 5px;
  border: 1px solid #ccc;
  border-radius: 4px;
  box-sizing: border-box;
  font-size: 14px;
}

.modal-actions {
  display: flex;
  justify-content: space-between;
  margin-top: 20px;
  gap: 10px;
}

.modal-content button {
  background-color: #f44336;
  color: white;
  border: none;
  padding: 10px;
  border-radius: 4px;
  cursor: pointer;
  width: 100%;
  font-size: 16px;
  transition: background-color 0.3s ease, transform 0.2s ease;
}

.modal-btn-create {
  background-color: #4CAF50;
}

.modal-btn-create:hover {
  background-color: #388E3C;
}

.modal-btn-close {
  background-color: #757575;
}

.modal-btn-close:hover {
  background-color: #616161;
}

.filters {
  display: flex;
  flex-wrap: nowrap; /* Размещение фильтров в одну строку */
  overflow-x: auto;
  gap: 15px; /* Отступы между фильтрами */
  margin-bottom: 20px;
  padding: 15px;
  background: #ffffff;
  border: 1px solid #ddd;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.filter-group {
  display: flex;
  flex-direction: column;
  min-width: 150px;
  max-width: 200px;
}

.filters input,
.filters select {
  width: 100%;
  padding: 8px;
  border: 1px solid #ccc;
  border-radius: 4px;
  box-sizing: border-box;
  font-size: 14px;
}

.apply-filters {
  flex: 0 0 auto;
  background-color: #a4a4a4;
  color: white;
  padding: 10px 20px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.3s ease, transform 0.2s ease;
  font-size: 16px;
}

.apply-filters:hover {
  background-color: #8d8d8d;
}
</style>