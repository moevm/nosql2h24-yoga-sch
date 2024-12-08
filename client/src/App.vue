<template>
  <div>
    <nav class="navbar">
      <div v-if="isLoggedIn" class="nav-icons">
        <!-- Значок выхода -->
        <router-link to="/" class="nav-link" @click="logout">
          <img src="https://img.icons8.com/?size=100&id=2445&format=png&color=ffffff" alt="Выход" class="nav-icon" />
        </router-link>
        <!-- Значок дом -->
        <router-link to="/admin" class="nav-link">
          <img src="https://img.icons8.com/?size=100&id=73&format=png&color=ffffff" alt="Дом" class="nav-icon" />
        </router-link>
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
  }, 200);
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
  display: flex;
  align-items: center;
  padding: 0 20px;
  color: white;
}

.nav-icons {
  display: flex;
  gap: 20px;
  align-items: center;
}

.nav-link {
  text-decoration: none;
  font-size: 1.2rem;
  transition: transform 0.3s;
}

.nav-icon {
  width: 30px;
  height: 30px;
  transition: transform 0.3s ease;
}

.nav-icon:hover {
  transform: scale(1.1);
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
