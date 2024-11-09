<template>
  <div class="entity-list-container">
    <h1 class="page-title">Данные из базы данных</h1>
    <ul class="entity-list">
      <li class="entity-item">
        <router-link
            :to="{ name: 'EntityList', params: { entityType: 'clients' } }"
            class="entity-link"
        >
          Список клиентов
        </router-link>
      </li>
      <li class="entity-item">
        <router-link
            :to="{ name: 'EntityList', params: { entityType: 'trainers' } }"
            class="entity-link"
        >
          Список тренеров
        </router-link>
      </li>
      <li class="entity-item">
        <router-link
            :to="{ name: 'EntityList', params: { entityType: 'studios' } }"
            class="entity-link"
        >
          Список студий
        </router-link>
      </li>
      <li class="entity-item">
        <router-link
            :to="{ name: 'EntityList', params: { entityType: 'classes' } }"
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
  padding: 0;
  text-align: center;
}

.page-title {
  font-family: 'Montserrat', sans-serif;
  color: #6A5862;
  font-size: 2rem;
  margin-bottom: 2rem;
}

ul.entity-list {
  list-style-type: none;
  padding: 0;
  margin: 0;
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
  transition: background-color 0.3s, transform 0.3s;
  text-align: center;
  width: auto;
}

.entity-link:hover {
  background-color: #D9D9D9;
}

.entity-link:active {
  background-color: #6A5862;
}

.entity-link:focus {
  outline: none;
  box-shadow: 0 0 0 2px #6A5862;
}
</style>
