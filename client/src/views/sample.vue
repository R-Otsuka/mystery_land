<template>
  <nav>
    <v-toolbar dark="true" class="elevation-1">
      <v-app-bar-nav-icon @click="drawer = !drawer"></v-app-bar-nav-icon>
      <v-toolbar-title>Title</v-toolbar-title>
      <v-spacer></v-spacer>
      <v-btn icon>
        <v-icon>mdi-keyboard-backspace</v-icon>
      </v-btn>
    </v-toolbar>

    <v-navigation-drawer class="navigation" v-model="drawer" app color="blue lighten-1">
      <v-list-item class="item" v-for="question in questions" :key="question.number">
        <v-list-item-icon>
          <v-icon v-if="question.open">
              {{ icon.unlock }}
          </v-icon>
          <v-icon v-else>
            {{ icon.lock }}
          </v-icon>
          <v-list-item-content>Q{{ question.number }}</v-list-item-content>
        </v-list-item-icon>
      </v-list-item>
    </v-navigation-drawer>
  </nav>
</template>

<script>
export default {
  data:() =>({
    drawer:false,
    icon:{
      "unlock":"mdi-lock-open-variant",
      "lock":"mdi-lock",
    },
    questions:[
      {number:1,open:true},
      {number:2,open:true},
      {number:3,open:false},
    ]
  }),
}
</script>
<style>
.navigation{
  width:1000px;
  background-color:white;
}
</style>