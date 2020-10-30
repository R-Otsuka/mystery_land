<template>
  <v-content class="py-10">
    <div class="mx-auto text-center py-5">
      <v-responsive class=" rounded-circle d-inline-flex success" height="64" width="64"></v-responsive>
    </div>
    <v-container>
<!--      start,center,など設定によりseparateを変えられる(公式参照)-->
      <v-row justify="space-around" class="py-5">
        <v-col cols="3" class="text-sm-left">
          <v-responsive class="rounded-circle d-inline-flex primary lighten-2 text-sm-left" height="64" width="64"></v-responsive>
        </v-col>
        <v-col cols="3" class="text-sm-right">
          <v-responsive class="rounded-circle d-inline-flex grey lighten-2" height="64" width="64" ></v-responsive>
        </v-col>
      </v-row>
      <p class="mx-auto text-center pb-10" style="width: 200px;">⠁⠡⠁⠊⠩⠚⠷⠐⠞⠓</p>
      <v-row justify="center" class="py-5">
        <v-col cols="3" class="text-sm-left">
          <v-responsive class="rounded-circle d-inline-flex purple lighten-2" height="64" width="64"></v-responsive>
        </v-col>
        <v-col cols="3" class="text-sm-right">
          <v-responsive class="rounded-circle d-inline-flex warning lighten-2" height="64" width="64"></v-responsive>
        </v-col>
      </v-row>
    </v-container>
  </v-content>
</template>

<script>
export default {

}
</script>
