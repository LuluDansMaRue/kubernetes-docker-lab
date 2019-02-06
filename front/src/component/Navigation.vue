<template>
  <div class="navbar">
    <router-link to="/">Home</router-link>
    <router-link to="/add">Add</router-link>
  </div>
</template>

<style lang="scss">
.navbar {
  width: 100px;
  height: 300px;
  background-color: white;
  position: fixed;
  top: calc(50%);
  margin-top: -150px;
  left: 0;
  padding-left: 20px;
  
  a {
    display: block;
    color: black;
    font-size: 1.2em;
    text-decoration: none;
    padding: 10px 0px;
  }
}
</style>

<script>
export default {}
</script>
