<template>
  <div class="navbar">
    <div class="navbar__character">
      <p>珍珠奶茶</p>
    </div>
    <div class="navbar__link">
      <router-link to="/">Home</router-link>
      <router-link to="/add">Add</router-link>
    </div>
  </div>
</template>

<style lang="scss">
.navbar {
  width: 100%;
  height: 70px;
  display: flex;

  &__link {
    background-color: white;
    width: 50%;
    padding: 20px 10px;
    text-align: right;

    a {
      color: black;
      font-size: 1.2em;
      text-decoration: none;
      padding: 0px 20px;
      border-bottom: 1px solid white;
      transition: all ease .25s;

      &:hover {
        border-bottom: 1px solid black;
      }
    }
  }

  &__character {
    background-color: black;
    width: 50%;

    p {
      color: white;
      font-size: 1.3em;
      padding-left: 15px;
    }
  }
}
</style>

<script>
export default {}
</script>
