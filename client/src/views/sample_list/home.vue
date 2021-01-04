<template>
  <div id="app" class="py-10 mx-5">
    <p>Vue.js、Goで作った脱出ゲームです。是非チャレンジしてみてください。</p>
    <p>感想・コメントも受け付けております。</p>
    <span>
      追加予定の機能<br>
      ・ステージの増設<br>
      ・クリア時に表示される文言やステージ選択をユーザーがカスタマイズができるようにする<br>
      ・クリア時間の測定とランキング機能の実装<br>
    </span>
    <router-link class="routerLink" to="/sample/1">
      <v-btn rounded color="success" class="text-center margin">
        Start
      </v-btn>
    </router-link>
  </div>
</template>

<script>
export default {

}
</script>
<style scoped>
.margin{
  margin:200px;
}

</style>