<script setup lang="ts">
defineProps<{
  msg: string;
}>();
</script>

<template>
  <div>
    <h1 class="site-title">Youga Places</h1>
  </div>
</template>

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
</style>
