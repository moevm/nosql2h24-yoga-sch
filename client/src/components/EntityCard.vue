<template>
  <div class="card-container" v-if="item">
    <h2 class="title">Details: {{ entityType }}</h2>
    <div class="card">
      <template v-for="field in getColumnConfig(entityType)" :key="field.key">
        <div v-if="field.key !== 'password'" class="card-field">
          <strong class="field-label">{{ field.label }}:</strong>
          <span v-if="field.isLink">
            <a :href="item[field.key as keyof Entity]" target="_blank" class="link">
              {{ item[field.key as keyof Entity] }}
            </a>
          </span>
          <span v-else-if="field.isList && field.isInfo">
            <ul class="list">
              <li
                  class="list-item"
                  v-for="elem in item[field.key as keyof Entity]"
                  :key="elem.id || elem"
              >
                {{ elem.name || elem }}
              </li>
            </ul>
          </span>
          <span v-else-if="field.isInfo">
            {{ item[field.key as keyof Entity].name }}
          </span>
          <span v-else-if="field.isDate && field.isTime">
            {{ getDateTime(item[field.key as keyof Entity]) }}
          </span>
          <span v-else-if="field.isDate && !field.isTime">
            {{ getDate(item[field.key as keyof Entity]) }}
          </span>
          <span v-else>
            {{ item[field.key as keyof Entity] }}
          </span>
        </div>
      </template>
    </div>
  </div>
  <p v-else class="no-element">No element found for ID {{ entityId }}</p>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";
import { useRoute } from "vue-router";
import { type Client, type Class, type Studio, type Trainer } from "@/types/types";

type Entity = Client | Class | Studio | Trainer | Record<string, any>;

const route = useRoute();
const entityType = ref(route.params.entityType as 'client' | 'trainer' | 'class' | 'studio');
const entityId = ref(route.params.id);

const URI = `${window.location.protocol}//${window.location.hostname}:${window.location.port}`;
let item = ref<Entity | null>(null);

onMounted(async () => {
  await loadItem();
});

async function loadItem() {
  try {
    const response = await fetch(`${URI}/api/v1/${entityType.value}/${entityId.value}`, {
      method: "GET",
    });
    const data = await response.json();
    item.value = data[entityType.value] || null;
    console.log(item);
  } catch (error) {
    console.error("Error while loading item:", error);
  }
}

type ColumnConfig = {
  key: string;
  label: string;
  isDate?: boolean;
  isTime?: boolean;
  isList?: boolean;
  isInfo?: boolean;
  isLink?: boolean;
};

function getColumnConfig(type: string): ColumnConfig[] {
  switch (type) {
    case "client":
      return [
        { key: "id", label: "ID" },
        { key: "name", label: "Name" },
        { key: "phone", label: "Phone" },
        { key: "gender", label: "Gender" },
        { key: "birthDate", label: "Birth Date", isDate: true },
        { key: "createdAt", label: "Created At", isDate: true, isTime: true },
        { key: "updatedAt", label: "Updated At", isDate: true, isTime: true },
        { key: "pictureUri", label: "Picture", isLink: true },
        { key: "classesInfo", label: "Classes", isList: true, isInfo: true },
      ];
    case "trainer":
      return [
        { key: "id", label: "ID" },
        { key: "name", label: "Name" },
        { key: "phone", label: "Phone" },
        { key: "gender", label: "Gender" },
        { key: "birthDate", label: "Birth Date", isDate: true },
        { key: "createdAt", label: "Created At", isDate: true, isTime: true },
        { key: "updatedAt", label: "Updated At", isDate: true, isTime: true },
        { key: "studioInfo", label: "Studio", isInfo: true },
        { key: "pictureUri", label: "Picture", isLink: true },
        { key: "classesInfo", label: "Classes", isList: true, isInfo: true },
      ];
    case "class":
      return [
        { key: "id", label: "ID" },
        { key: "name", label: "Class Type" },
        { key: "time", label: "Time", isDate: true, isTime: true },
        { key: "studioInfo", label: "Studio", isInfo: true },
        { key: "trainerInfo", label: "Trainer", isInfo: true },
        { key: "clientsInfo", label: "Clients", isList: true, isInfo: true },
      ];
    case "studio":
      return [
        { key: "id", label: "ID" },
        { key: "address", label: "Address" },
        { key: "classesInfo", label: "Classes", isList: true, isInfo: true },
        { key: "trainersInfo", label: "Trainers", isList: true, isInfo: true },
      ];
    default:
      return [];
  }
}

function getDateTime(date: string) {
  return new Date(date).toLocaleString();
}

function getDate(date: string) {
  return new Date(date).toDateString();
}
</script>

<style scoped>
.card-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 1.5rem;
  border: 1px solid #ccc;
  border-radius: 12px;
  background: linear-gradient(135deg, #fdfbfb, #ebedee);
  max-width: 600px;
  margin: 2rem auto;
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.1);
  position: relative;
}

.title {
  font-size: 1.8rem;
  margin-bottom: 1.5rem;
  color: #333;
  text-align: center;
}

.card {
  display: flex;
  flex-direction: column;
  gap: 1rem;
  position: relative;
}

.card-field {
  display: flex;
  flex-direction: column;
  align-items: center;
  text-align: center;
}

.field-label {
  font-weight: bold;
  color: #555;
}

.link {
  color: #007bff;
  text-decoration: none;
  transition: color 0.3s ease;
}

.link:hover {
  color: #0056b3;
}

.list {
  list-style-type: none;
  padding: 0;
  margin: 0;
}

.list-item {
  background: #f7f7f7;
  padding: 0.5rem;
  margin: 0.3rem 0;
  border-radius: 6px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.no-element {
  text-align: center;
  color: #999;
  font-size: 1.2rem;
  margin-top: 2rem;
}
</style>
