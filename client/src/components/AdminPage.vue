<template>
  <div class="admin-container">
    <div class="header-container">
      <h1 class="page-title">Admin Page</h1>
    </div>
    <div class="button-container">
      <button class="admin-button" @click="viewDatabase">View Database</button>
      <button class="admin-button" @click="triggerFileInput">Import Database</button>
      <button class="admin-button" @click="exportDatabase">Export Database</button>
      <input
          type="file"
          ref="fileInput"
          accept="application/json"
          style="display: none"
          @change="importDatabase"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import router from '@/router';

const URI = `${window.location.protocol}//${window.location.hostname}`;
const fileInput = ref<HTMLInputElement | null>(null);

function viewDatabase() {
  router.push('/admin/data');
}

function triggerFileInput() {
  fileInput.value?.click();
}

async function importDatabase(event: Event) {
  const input = event.target as HTMLInputElement;
  const file = input.files?.[0];

  if (!file) {
    alert('No file selected.');
    return;
  }

  try {
    const fileContent = await file.text();
    const jsonData = JSON.parse(fileContent);

    console.log('Importing database:', jsonData);

    const response = await fetch(`${URI}/api/admin/v1/db/import`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(jsonData),
    });

    if (!response.ok) {
      throw new Error(`Failed to import database: ${response.statusText}`);
    }

    alert('Database imported successfully!');
  } catch (error) {
    console.error(error);
    alert('An error occurred while importing the database. Please check the file format.');
  } finally {
    input.value = '';
  }
}

async function exportDatabase() {
  try {
    const response = await fetch(`${URI}/api/admin/v1/db/export`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
    });

    if (!response.ok) {
      throw new Error(`Failed to export database: ${response.statusText}`);
    }

    const data = await response.json();
    const blob = new Blob([JSON.stringify(data, null, 2)], { type: 'application/json' });
    const url = URL.createObjectURL(blob);

    const a = document.createElement('a');
    a.href = url;
    a.download = 'database_export.json';
    document.body.appendChild(a);
    a.click();
    document.body.removeChild(a);
    URL.revokeObjectURL(url);
  } catch (error) {
    console.error(error);
    alert('An error occurred while exporting the database.');
  }
}
</script>

<style scoped>
.admin-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: flex-start;
  padding: 20px 0;
  margin: 0;
  height: 100vh;
}

.header-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 150px;
  width: 100%;
}

.page-title {
  font-family: 'Montserrat', sans-serif;
  color: #6A5862;
  font-size: 2rem;
  margin: 0;
}

.button-container {
  display: flex;
  flex-direction: column;
  gap: 20px;
  margin-top: 20px;
}

.admin-button {
  background-color: #D9D9D9;
  color: #6A5862;
  border: none;
  padding: 15px 30px;
  border-radius: 10px;
  cursor: pointer;
  font-size: 18px;
  font-family: 'Montserrat', sans-serif;
  font-weight: 600;
  transition: transform 0.2s ease, box-shadow 0.2s ease;
  width: 200px;
}

.admin-button:hover {
  transform: scale(1.05);
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.1);
}

.admin-button:active {
  transform: scale(0.98);
}
</style>
