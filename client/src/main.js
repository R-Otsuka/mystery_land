import Vue from 'vue' //packageのVue
import App from './App.vue' //App.vueのコンポーネントファイル
import router from './router'
import vuetify from './plugins/vuetify';
import axios from 'axios' //追記
import VueAxios from 'vue-axios' //追記

Vue.config.productionTip = false //お助け、productionモードではfalseであるべき。開発モードではtrueでも可
Vue.use(VueAxios, axios) //追記

new Vue({
  router,
  vuetify,
  // render: h => h("div","こんにちは")
  render: h => h(App) //ES6の書き方
  // render: function(h){ ES5の書き方
  //   return h(App)
  // }
}).$mount('#app') //mount
