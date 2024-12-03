<template>
  <div class="entity-list-container">
    <div class="header-container">
      <h1 class="page-title">Данные из базы данных</h1>
    </div>
    <ul class="entity-list">
      <li class="entity-item">
        <router-link
            :to="{ name: 'EntityList', params: { entityType: 'client' } }"
            class="entity-link"
        >
          Список клиентов
        </router-link>
      </li>
      <li class="entity-item">
        <router-link
            :to="{ name: 'EntityList', params: { entityType: 'trainer' } }"
            class="entity-link"
        >
          Список тренеров
        </router-link>
      </li>
      <li class="entity-item">
        <router-link
            :to="{ name: 'EntityList', params: { entityType: 'studio' } }"
            class="entity-link"
        >
          Список студий
        </router-link>
      </li>
      <li class="entity-item">
        <router-link
            :to="{ name: 'EntityList', params: { entityType: 'class' } }"
            class="entity-link"
        >
          Список занятий
        </router-link>
      </li>
    </ul>
  </div>
</template>

<script setup lang="ts">
</script>

<style scoped>
.entity-list-container {
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

ul.entity-list {
  list-style-type: none;
  padding: 0;
  margin: 20px 0 0;
}

li.entity-item {
  margin: 1rem 0;
}

.entity-link {
  display: inline-block;
  padding: 0.8rem 1rem;
  background-color: #6A5862;
  color: #fff;
  font-family: 'Montserrat', sans-serif;
  font-size: 1.2rem;
  border-radius: 8px;
  text-decoration: none;
  transition: transform 0.3s ease;
  text-align: center;
  width: auto;
}

.entity-link:hover {
  transform: scale(1.1);
}

.entity-link:active {
  transform: scale(1.05);
}

.entity-link:focus {
  outline: none;
  box-shadow: 0 0 0 2px #6A5862;
}
</style>
