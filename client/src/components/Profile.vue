<template>
  <div class="client-profile">
    <div v-if="loading" class="loading">Loading...</div>
    <div v-else-if="error" class="error">{{ error }}</div>
    <div v-else-if="client" class="profile-container">
      <!-- Левая колонка: Занятия -->
      <div class="classes-container">
        <h2>Classes</h2>
        <div class="class-cards">
          <div
              v-for="(classInfo, index) in client.classesInfo"
              :key="index"
              class="class-card"
          >
            <button class="delete-button" @click="deleteClass(classInfo.id)">
              <img
                  src="https://img.icons8.com/?size=100&id=3062&format=png&color=993737"
                  alt="Delete"
                  class="delete-icon"
              />
            </button>

            <div class="class-left">
              <p class="class-time">{{ new Date(classInfo.time).toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' }) }}</p> <!-- Время с увеличенным шрифтом -->
            </div>
            <div class="class-center">
              <h3>{{ classInfo.name }}</h3>
              <p v-if="classInfo.studioName">Студия: {{ classInfo.studioName }}</p>
              <p v-if="classInfo.studioName">Тренер: {{ classInfo.trainer }}</p>
            </div>
          </div>
          <div v-if="client.classesInfo.length === 0">
            <p>No classes information available.</p>
          </div>
        </div>
      </div>

      <div class="profile">
        <div class="profile-content">
          <img
              :src="client.pictureUri"
              alt="Profile Picture"
              class="profile-picture"
          />
          <h1 class="profile-name">{{ client.name }}</h1>
          <p class="profile-phone">Телефон: {{ client.phone }}</p>
          <p class="profile-birthdate">
            Дата рождения: {{ new Date(client.birthDate).toLocaleDateString() }}
          </p>
          <p class="profile-gender">Пол: {{ client.gender }}</p>
          <p class="profile-created-at">
            Участник с: {{ new Date(client.createdAt).toLocaleDateString() }}
          </p>
          <p class="profile-updated-at">
            Последнее обновление: {{ new Date(client.updatedAt).toLocaleDateString() }}
          </p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";
import type { PropType } from "vue";

const URI = `${window.location.protocol}//${window.location.hostname}:${window.location.port}`;

interface Studio {
  id: string;
  name: string;
}

interface ClassInfo {
  id: string;
  name: string;
  time: string;
  studioId: string;
  studioName?: string;
  trainer: string;
}

interface Client {
  id: string;
  phone: string;
  name: string;
  pictureUri: string;
  birthDate: string;
  gender: "FEMALE" | "MALE";
  createdAt: string;
  updatedAt: string;
  classesInfo: ClassInfo[];
}

const props = defineProps({
  id: {
    type: String as PropType<string>,
    required: true,
  },
});

const client = ref<Client | null>(null);
const loading = ref<boolean>(true);
const error = ref<string | null>(null);

const loadClient = async () => {
  try {
    loading.value = true;
    error.value = null;

    const response = await fetch(`${URI}/api/v1/client/${props.id}`);
    if (!response.ok) {
      throw new Error(`Failed to fetch client data: ${response.status}`);
    }

    const data = await response.json();
    client.value = data.client;

    const classesInfo: ClassInfo[] = [];
    for (const classId of data.client.classesInfo) {
      const classResponse = await fetch(`${URI}/api/v1/class/${classId.id}`);
      if (classResponse.ok) {
        const classData = await classResponse.json();
        classesInfo.push({
          id: classData.class.id,
          name: classData.class.name,
          time: classData.class.time,
          studioId: classData.class.studioId,
          studioName: classData.class.studioInfo.name,
          trainer: classData.class.trainerInfo.name,
        });
      } else {
        console.error(`Failed to fetch class data for classId ${classId.id}`);
      }
    }

    if (client.value) {
      client.value.classesInfo = classesInfo;
    }
  } catch (err) {
    error.value = "Failed to load client data.";
  } finally {
    loading.value = false;
  }
};

const deleteClass = async (classId: string) => {
  if (client.value) {
    try {
      const response = await fetch(`${URI}/api/v1/appointment/${client.value.id}/${classId}`, {
        method: 'DELETE',
      });
      if (response.ok) {
        // Удаляем класс из локального состояния
        client.value.classesInfo = client.value.classesInfo.filter(
            (classInfo) => classInfo.id !== classId
        );
      } else {
        console.error("Failed to delete class.");
      }
    } catch (err) {
      console.error("Error while deleting class:", err);
    }
  }
};

onMounted(() => {
  loadClient();
});
</script>

<style scoped>
.client-profile {
  font-family: Arial, sans-serif;
  padding: 20px;
  max-width: 1200px;
  margin: 0 auto;
}

.loading {
  text-align: center;
  color: #888;
}

.error {
  text-align: center;
  color: red;
}

.profile-container {
  display: flex;
  gap: 20px;
  justify-content: space-between;
  align-items: flex-start;
}

.classes-container {
  flex: 1;
  background: #6A5862;
  border: 1px solid #ddd;
  border-radius: 8px;
  padding: 20px;
}

.class-cards {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.class-card {
  background: #ffffff;
  padding: 15px;
  border: 1px solid #ddd;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  display: flex;
  align-items: center;
  position: relative; /* Для позиционирования кнопки */
}

.delete-button {
  position: absolute;
  right: 10px;
  top: 50%;
  transform: translateY(-50%);
  background: transparent;
  border: none;
  padding: 0;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
}

.delete-icon {
  width: 50px;
  height: 50px;
}

.class-left {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
}

.class-time {
  font-size: 48px; /* Очень большой шрифт для времени */
  font-weight: bold;
  color: #333;
}

.class-center {
  flex: 1;
  text-align: center;
}

.class-center h3 {
  margin: 0;
  font-size: 18px;
  font-weight: bold;
  color: #333;
}

.class-center p {
  margin-top: 5px;
  font-size: 14px;
  color: #555;
}

.profile {
  flex: 1;
  display: flex;
  justify-content: center;
  align-items: center;
  border-radius: 8px;
  padding: 20px;
}

.profile-content {
  text-align: center;
}

.profile-name {
  font-size: 30px;
  margin: 10px 0;
}

.profile-phone,
.profile-birthdate,
.profile-gender,
.profile-created-at,
.profile-updated-at {
  margin: 5px 0;
  font-size: 18px;
  color: #555;
}

.profile-picture {
  width: 120px;
  height: 120px;
  border-radius: 50%;
  object-fit: cover;
  margin-bottom: 20px;
}

.classes-container h2 {
  color: white;
}
</style>
