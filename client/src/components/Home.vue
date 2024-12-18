<template>
  <div>
    <h1 class="site-title">Youga Places</h1>
    <button class="register-button" @click="goToBookingPage">Записаться на занятие</button>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';

const router = useRouter();

function goToBookingPage() {
  router.push('/appointment');
}
</script>

<style scoped>

h3 {
  font-size: 1.2rem;
}

.greetings h1,
.greetings h3 {
  text-align: center;
}

@media (min-width: 1024px) {
  .greetings h1,
  .greetings h3 {
    text-align: left;
  }
}

.site-title {
  font-size: 36px;
  font-family: 'Montserrat', sans-serif;
  font-weight: 600;
  color: #000000;
  margin-bottom: 20px;
  text-align: center;
}

.register-button {
  display: block;
  margin: 20px auto;
  padding: 10px 20px;
  font-size: 16px;
  font-family: 'Montserrat', sans-serif;
  font-weight: 500;
  color: #ffffff;
  background-color: #6A5862;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  text-align: center;
}

.register-button:hover {
  transform: scale(1.05);
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.1);
}

</style>
