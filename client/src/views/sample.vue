<template>
  <nav>
    <v-toolbar color="primary" dark="true">
      <v-app-bar-nav-icon @click="drawer = !drawer"></v-app-bar-nav-icon>
      <v-toolbar-title>Title</v-toolbar-title>
      <v-spacer></v-spacer>
      <v-btn icon>
        <v-icon>mdi-keyboard-backspace</v-icon>
      </v-btn>
    </v-toolbar>

    <v-navigation-drawer v-model="drawer" app class="success">
      <v-list-item>
        <v-list-item-icon>
          <v-icon></v-icon>
          <v-list-item-content></v-list-item-content>
        </v-list-item-icon>
      </v-list-item>
    </v-navigation-drawer>
  </nav>
</template>

<script>
export default {
  data:() =>({
    drawer:false,
  }),
}
</script>
