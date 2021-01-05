<template>
  <div id="app" class="py-10 mx-10">
    <div class="text-center">
      <h1 class="py-15 text-h3">PAZZLVE</h1>
      <router-link class="routerLink" to="/sample/1">
        <v-btn rounded color="success" class="text-center margin">
          Start
        </v-btn>
      </router-link>
    </div>
    <accordion v-for="item in items" v-bind:item="item" v-bind:key="item.id"></accordion>
  </div>
</template>

<script>

export default {
  data(){
    return{
      items:[
        { id:0,title:"何のアプリ？",text:"学習中のVue.js、Goで作った脱出ゲームです。" },
        { id:1,title:"追加予定の機能は？",text:"・ステージの増設<br>・クリア時に表示される文言やステージ選択をユーザーがカスタマイズができるようにする<br>・クリア時間の測定とランキング機能の実装"},
      ]
    }
  }
}
</script>
<style scoped>
.margin{
  margin-top:100px;
  margin-bottom:200px;
}

</style>