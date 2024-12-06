<template>
  <div class="admin-container">
    <div class="header-container">
      <h1 class="page-title">Admin Page</h1>
    </div>
    <div class="button-container">
      <button class="admin-button" @click="viewDatabase">View Database</button>
      <button class="admin-button" @click="importDatabase">Import Database</button>
      <button class="admin-button" @click="exportDatabase">Export Database</button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import router from '@/router';

function viewDatabase() {
  router.push('/admin/data');
}

function importDatabase() {
  alert('Import database functionality here');
}

function exportDatabase() {
  alert('Export database functionality here');
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
