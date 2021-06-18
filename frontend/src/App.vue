<template>
  <div id="app">
    <router-view></router-view>
  </div>
</template>

<style lang="scss">
@import url('./common/common.scss');
html,
body {
  margin: 0;
  padding: 0;
  font-family: PingFangSC-Light, PingFang SC;
  // font-size: 20px;
  background-color: #F6F5FC;
}
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  min-height: 100%;
  background-color: #F6F5FC;
  max-width: 500PX;
  margin: 0 auto;
}
</style>
