<template>
  <nav class="navbar">
    <router-link to="/" class="nav-link">Home</router-link> |
    <router-link to="/data" class="nav-link">Database</router-link>
  </nav>
  <main>
    <router-view />
  </main>
</template>

<style scoped>
#app {
  font-family: 'Montserrat', sans-serif;
  text-align: center;
  margin-top: 20px;
}

.navbar {
  background-color: #6A5862;
  padding: 10px 0;
  display: flex;
  justify-content: center;
  align-items: center;
}

.nav-link {
  color: white;
  text-decoration: none;
  font-size: 1.2rem;
  margin: 0 10px;
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
