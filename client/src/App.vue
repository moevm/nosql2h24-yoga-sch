<template>
  <div>
    <nav class="navbar">
      <div class="nav-icons-left">
        <div v-if="isAdminLoggedIn">
          <!-- Значок выхода -->
          <router-link to="/" class="nav-link" @click="logout">
            <img src="https://img.icons8.com/?size=100&id=2445&format=png&color=ffffff" alt="Выход" class="nav-icon" />
          </router-link>
          <!-- Значок дом -->
          <router-link to="/admin" class="nav-link">
            <img src="https://img.icons8.com/?size=100&id=73&format=png&color=ffffff" alt="Дом" class="nav-icon" />
          </router-link>
        </div>
        <div v-if="isUserLoggedIn">
          <!-- Значок выхода -->
          <router-link to="/" class="nav-link" @click="logout">
            <img src="https://img.icons8.com/?size=100&id=2445&format=png&color=ffffff" alt="Выход" class="nav-icon" />
          </router-link>
          <!-- Значок дом -->
          <router-link to="/home" class="nav-link">
            <img src="https://img.icons8.com/?size=100&id=73&format=png&color=ffffff" alt="Дом" class="nav-icon" />
          </router-link>
        </div>
      </div>
      <div class="nav-icons-right" v-if="isUserLoggedIn">
        <!-- Значок профиля -->
        <router-link to="/profile" class="nav-link">
          <img src="https://img.icons8.com/?size=100&id=7823&format=png&color=ffffff" alt="Профиль" class="nav-icon" />
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

const isAdminLoggedIn = ref(false);
const isUserLoggedIn = ref(false);

function getCookie(name: string): string | null {
  const matches = document.cookie.match(new RegExp(`(?:^|; )${name.replace(/([.$?*|{}()[\]\\/+^])/g, '\\$1')}=([^;]*)`));
  return matches ? decodeURIComponent(matches[1]) : null;
}

function checkAuthorization() {
  isAdminLoggedIn.value = getCookie('Authorization') === 'admin';
  isUserLoggedIn.value = getCookie('Authorization') === 'user';
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
  isAdminLoggedIn.value = false;
  isUserLoggedIn.value = false;
}
</script>

<style scoped>
#app {
  font-family: 'Montserrat', sans-serif;
  text-align: center;
  margin-top: 20px;
}

main {
  justify-content: center;
  padding: 20px;
}

h1 {
  text-align:center;
}

.navbar {
  background-color: #6A5862;
  height: 60px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0 20px;
  color: white;
}

.nav-icons-left {
  display: flex;
  gap: 20px;
  align-items: center;
}

.nav-icons-right {
  display: flex;
  align-items: center;
}

.nav-icon {
  width: 30px;
  height: 30px;
  transition: transform 0.3s ease;
  margin: 5px;
}

.nav-icon:hover {
  transform: scale(1.1);
}

</style>
