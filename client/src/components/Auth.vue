<template>
  <h1 class="site-title">Youga Places</h1>
  <div class="auth-container">
    <div class="auth-form">
      <div v-if="!isLogin" class="arrow-button">
        <button @click="toggleAuthMode" class="back-button">
          <span class="back-icon">←</span>
        </button>
      </div>
      <h2 class="auth-header">{{ isLogin ? 'Log In' : 'Sign Up' }}</h2>
      <form @submit.prevent="isLogin ? loginUser() : registerUser()">
        <div class="form-group">
          <label for="phone">Phone number</label>
          <input
              v-model="formattedPhone"
              type="tel"
              id="phone"
              placeholder="+7(XXX)XXX-XXXX"
              required
          />
        </div>
        <div class="form-group">
          <label for="password">Password</label>
          <input
              v-model="formData.password"
              type="password"
              id="password"
              placeholder="Enter your password"
              required
          />
        </div>
        <div v-if="!isLogin" class="form-group">
          <label for="name">Name</label>
          <input
              v-model="formData.name"
              type="text"
              id="name"
              placeholder="Enter your name"
              required
          />
        </div>
        <div v-if="!isLogin" class="form-group">
          <label for="gender">Gender</label>
          <select
              v-model="formData.gender"
              id="gender"
              required
          >
            <option value="F">Female</option>
            <option value="M">Male</option>
          </select>
        </div>
        <div class="auth-actions">
          <button
              v-if="isLogin"
              type="submit"
              class="auth-button"
          >
            Log In
          </button>
          <button
              v-if="isLogin"
              type="button"
              class="toggle-button"
              @click="toggleAuthMode"
          >
            Sign Up
          </button>
          <button
              v-else
              type="submit"
              class="auth-button"
          >
            Sign Up
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue';
import router from "@/router";

const URI = `${window.location.protocol}//${window.location.hostname}`;
const isLogin = ref(true);
const formData = ref({
  phone: '',
  password: '',
  name: '',
  gender: 'F',
});

const formattedPhone = computed({
  get: () => formatPhone(formData.value.phone),
  set: (value: string) => {
    formData.value.phone = value;
  },
});

function formatPhone(value: string): string {
  let input = value.replace(/\D/g, ''); // Удаляем всё, кроме цифр
  const size = input.length;

  if (size === 0) {
    return `+7(`;
  }

  if (input.startsWith('7')) {
    input = input.substring(1);
  }

  if (size < 4) {
    return `+7(${input}`;
  } else if (size < 7) {
    return `+7(${input.substring(0, 3)})${input.substring(3)}`;
  } else if (size <= 10) {
    return `+7(${input.substring(0, 3)})${input.substring(3, 6)}-${input.substring(6)}`;
  } else {
    return `+7(${input.substring(0, 3)})${input.substring(3, 6)}-${input.substring(6, 10)}`;
  }
}

function toggleAuthMode() {
  isLogin.value = !isLogin.value;
}

async function loginUser() {
  console.log('Logging in with phone:', formData.value.phone);
  try {
    const response = await fetch(`${URI}/api/v1/auth`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        phone: formData.value.phone,
        password: formData.value.password,
      }),
    });

    if (!response.ok) {
      throw new Error('Login failed. Please check your credentials.');
    }

    const result = await response.json();
    console.log('Login response:', result);

    document.cookie = `Authorization=admin; path=/; max-age=86400; secure; SameSite=Strict`;
    await router.push('/admin');
  } catch (error) {
    alert('Login failed. Please try again.');
  }
}

async function registerUser() {
  console.log('Registering with phone:', formData.value.phone);
  try {
    const response = await fetch(`${URI}/api/v1/client`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        client: {
          phone: formData.value.phone,
          password: formData.value.password,
          name: formData.value.name,
          gender: formData.value.gender === 'F' ? 'FEMALE' : 'MALE',
        }
      }),
    });

    if (!response.ok) {
      throw new Error('Registration failed. Please check your input.');
    }

    const result = await response.json();
    console.log('Registration response:', result);

    alert('Registration successful! Please log in.');
    isLogin.value = true;
  } catch (error) {
    alert('Registration failed. Please try again.');
  }
}
</script>

<style scoped>
.auth-container {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  margin: 100px;
}

.site-title {
  font-size: 36px;
  font-family: 'Montserrat', sans-serif;
  font-weight: 600;
  color: #000000;
  margin-bottom: 20px;
}

.auth-form {
  background-color: #6a5862;
  top: 0px;
  padding: 40px;
  border-radius: 12px;
  width: 100%;
  max-width: 400px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  text-align: center;
  position: relative;
}

.auth-form h2 {
  margin: 0;
  font-size: 24px;
  font-family: 'Montserrat', sans-serif;
  font-weight: 600;
  color: #ffffff;
  margin-bottom: 20px;
}

form {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.form-group {
  display: flex;
  flex-direction: column;
  width: 100%;
}

.form-group label {
  margin-bottom: 5px;
  font-size: 14px;
  color: #ffffff;
  font-family: 'Montserrat', sans-serif;
  text-align: left;
}

.form-group input,
.form-group select {
  width: 100%;
  padding: 12px;
  border: none;
  border-radius: 4px;
  background-color: #d9d9d9;
  font-size: 14px;
  box-sizing: border-box;
}

.form-group select {
  appearance: none; /* Remove default dropdown arrow */
  background-color: #d9d9d9; /* Same as the input field */
  color: #6a5862; /* Match input text color */
  padding-right: 30px; /* Space for custom dropdown arrow */
}

.form-group select:focus {
  outline: none;
  box-shadow: 0 0 5px rgba(0, 0, 0, 0.2);
}

.auth-actions {
  display: flex;
  justify-content: center;
  gap: 10px;
}

.auth-button,
.toggle-button {
  background-color: #d9d9d9;
  color: #6a5862;
  border: none;
  padding: 12px 20px;
  border-radius: 20px;
  cursor: pointer;
  font-size: 16px;
  font-family: 'Montserrat', sans-serif;
  font-weight: 600;
  transition: transform 0.2s ease;
}

.auth-button:hover,
.toggle-button:hover {
  transform: scale(1.05);
}

.back-button {
  background-color: transparent;
  border: none;
  font-size: 24px;
  color: #ffffff;
  cursor: pointer;
  transition: transform 0.2s ease;
  position: absolute;
  top: 20px;
  left: 20px;
}

.back-button:hover {
  transform: scale(1.1);
}

.back-icon {
  font-size: 24px;
  color: #ffffff;
}

.auth-header {
  margin: 0;
  text-align: center;
}

.select-wrapper {
  position: relative;
}

.select-wrapper::after {
  content: '▼';
  position: absolute;
  right: 10px;
  top: 50%;
  transform: translateY(-50%);
  font-size: 16px;
  color: #6a5862;
  pointer-events: none;
}
</style>
