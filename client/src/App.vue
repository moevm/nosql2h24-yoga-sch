<template>
  <div>
    <nav class="navbar">
      <div v-if="isLoggedIn" class="nav-links">
        <router-link to="/" class="nav-link" @click="logout">Выход</router-link>
        <router-link to="/data" class="nav-link">База данных</router-link>
      </div>
    </nav>
    <main>
      <router-view />
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue';

const isLoggedIn = ref(false);

function getCookie(name: string): string | null {
  const matches = document.cookie.match(new RegExp(`(?:^|; )${name.replace(/([.$?*|{}()[\]\\/+^])/g, '\\$1')}=([^;]*)`));
  return matches ? decodeURIComponent(matches[1]) : null;
}

function checkAuthorization() {
  isLoggedIn.value = getCookie('Authorization') === 'admin';
}

let intervalId: number | null = null;

onMounted(() => {
  checkAuthorization();

  intervalId = window.setInterval(() => {
    checkAuthorization();
  }, 1000);
});

onUnmounted(() => {
  if (intervalId !== null) {
    clearInterval(intervalId);
  }
});

function logout() {
  document.cookie = "Authorization=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/";
  isLoggedIn.value = false;
}
</script>

<style scoped>
#app {
  font-family: 'Montserrat', sans-serif;
  text-align: center;
  margin-top: 20px;
}

.navbar {
  background-color: #6A5862;
  height: 60px;
  padding: 0 20px;
  display: flex;
  justify-content: center;
  align-items: center;
  color: white;
  overflow: hidden;
}

.nav-links {
  display: flex;
  gap: 15px;
}

.nav-link {
  color: white;
  text-decoration: none;
  font-size: 1.2rem;
  transition: color 0.3s;
}

.nav-link:hover {
  color: #969696;
}

main {
  margin-top: 20px;
  padding: 20px;
  background-color: #F9F9F9;
  min-height: 80vh;
  text-align: center;
  display: block;
}
</style>
