<template>
  <div class="home">
    <h1>This is HomePage</h1>
    <p class="red blue--text">
      いかがでしたでしょうか。 Vuetifyなどのデザインフレームワークを使わずにスクラッチでデザインを構築する場合も、Vuetifyの命名規則や構成はとても参考になると思います！是非一度使ってみて下さい。
      今回は以上になります。最後までお読み頂きありがとうございました！
    </p>
    <p class="pink lighten-4 red--text text--darken4">
      hello
    </p>
    <div class="display-4">
      Massive
    </div>
    <v-row
        align="center"
        justify="space-around"
    >
      <v-btn class="pink white--text">Click Me</v-btn>
      <v-btn depressed color="pink">Click Me</v-btn>
    </v-row>
  </div>
</template>

<script>

export default {

}

</script>
