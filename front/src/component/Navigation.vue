<template>
  <div class="navbar">
    <router-link to="/">Home</router-link>
    <router-link to="/add">Add</router-link>
  </div>
</template>

<style lang="scss">
.navbar {
  width: 300px;
  background-color: white;
  position: fixed;
  top: calc(50% - 300px);
  left: 0;
}
</style>

<script>
export default {}
</script>
