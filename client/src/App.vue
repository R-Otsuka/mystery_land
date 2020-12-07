<template>
  <v-app>
    <nav>
      <v-main>
        <v-toolbar color="primary" class="elevation-2">
          <v-app-bar-nav-icon @click="drawer = !drawer"></v-app-bar-nav-icon>
          <v-toolbar-title>Question</v-toolbar-title>
          <v-spacer></v-spacer>
          <v-btn icon>
            <v-icon>mdi-keyboard-backspace</v-icon>
          </v-btn>
        </v-toolbar>
        <div class="mt-8">
          <v-card class="mx-auto" max-width="1000">
            <router-view></router-view>
          </v-card>
        </div>
      </v-main>

      <v-navigation-drawer v-model="drawer" app>
        <v-list-item v-for="link in links" :key="link.text" router :to="link.route">
          <v-list-item-icon class="itemIcon">
            <v-icon v-if=link.icon>
              {{ icon.unlock }}
            </v-icon>
            <v-icon v-else>
              {{ icon.lock }}
            </v-icon>
          </v-list-item-icon>
          <v-list-item-content>
            <v-list-item-title>{{ link.text }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-navigation-drawer>
    </nav>
  </v-app>
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
