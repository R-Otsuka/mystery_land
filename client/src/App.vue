<template>
  <div>
    <nav>
      <v-content>
        <v-toolbar app flat>
          <v-app-bar-nav-icon @click="drawer = !drawer"></v-app-bar-nav-icon>
          <v-toolbar-title>Question</v-toolbar-title>
          <v-spacer></v-spacer>
          <v-btn icon>
            <v-icon>mdi-keyboard-backspace</v-icon>
          </v-btn>
        </v-toolbar>
      </v-content>

      <v-navigation-drawer v-model="drawer" app class="success">
        <v-list-item v-for="link in links" :key="link.text" router :to="link.route">
          <v-list-item-icon>
            <v-icon class="white--text" v-if=link.icon>
              {{ icon.unlock }}
            </v-icon>
            <v-icon class="white--text" v-else>
              {{ icon.lock }}
            </v-icon>
          </v-list-item-icon>
          <v-list-item-content>
            <v-list-item-title class="white--text">{{ link.text }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-navigation-drawer>
    </nav>
    <v-content>
      <router-view></router-view>
    </v-content>
  </div>
</template>

<script>
export default {
  data:() =>({
    drawer:false,
    icon:{
      "unlock":"mdi-lock-open-variant",
      "lock":"mdi-lock",
    },
    links: [
      {icon:true, text:'Q1',route:'/sample/1'},
      {icon:true, text:'Q2',route:'/sample/2'},
      {icon:false, text:'Q3',route:'/sample/3'},
    ]
  }),
}
</script>