<template>
<!--  v-contentタグはサイドバー表示の際にpadding-leftの自動追加機能あり??-->
<!--  <v-content class="py-10">-->
  <v-container>
    <div class="mx-auto text-center py-5">
      <v-responsive class="rounded-circle d-inline-flex"
                    :class=items[0].color
                    :style="{visibility: items[0].visibility ? 'visible' : 'hidden'}"
                    @click="hide(0)"
                    height="64" width="64">
      </v-responsive>
    </div>
<!--      start,center,など設定によりseparateを変えられる(公式参照)-->
    <v-row justify="space-around" class="py-5">
      <v-col cols="3" class="text-sm-left">
        <v-responsive class="rounded-circle d-inline-flex"
                      v-bind:class=items[1].color
                      :style="{visibility: items[1].visibility ? 'visible' : 'hidden'}"
                      @click="hide(1)"
                      height="64" width="64">

        </v-responsive>
      </v-col>
      <v-col cols="3" class="text-sm-right">
        <v-responsive class="rounded-circle d-inline-flex"
                      v-bind:class=items[2].color
                      :style="{visibility: items[2].visibility ? 'visible' : 'hidden'}"
                      @click="hide(2)"
                      height="64" width="64">

        </v-responsive>
      </v-col>
    </v-row>
    <p class="mx-auto text-center pb-10" style="width: 200px;">⠁⠡⠁⠊⠩⠚⠷⠐⠞⠓</p>
    <v-row justify="center" class="py-5">
      <v-col cols="3" class="text-sm-left">
        <v-responsive class="rounded-circle d-inline-flex"
                      v-bind:class=items[3].color
                      :style="{visibility: items[3].visibility ? 'visible' : 'hidden'}"
                      @click="hide(3)"
                      height="64" width="64">

        </v-responsive>
      </v-col>
      <v-col cols="3" class="text-sm-right">
        <v-responsive class="rounded-circle d-inline-flex"
                      v-bind:class=items[4].color
                      :style="{visibility: items[4].visibility ? 'visible' : 'hidden'}"
                      @click="hide(4)"
                      height="64" width="64">

        </v-responsive>
      </v-col>
    </v-row>
    <button @click="startshuffle">shuffle</button>
    <br>
    <button @click="stopshuffle">stop</button>
  </v-container>
<!--  </v-content>-->
</template>
<script>
import _ from 'lodash';
export default {
  data:()=> ({
    setfunc:false,
    showOrNot:false,
    items:[
      {id:0,color:'warning',visibility:true},
      {id:1,color:'purple',visibility:true},
      {id:2,color:'grey',visibility:true},
      {id:3,color:'primary',visibility:true},
      {id:4,color:'success',visibility:true}
    ]
  }),
  methods: {
    onShuffleClick() {
      this.items = _.shuffle(this.items);
    },
    startshuffle(){
      if(this.setfunc == true){
        setTimeout(this.startshuffle, 3000);
        this.onShuffleClick()
      }else{
        this.setfunc = true
      }
    },
    stopshuffle(){
      this.setfunc = false
    },
    hide(id){
      this.items[id].visibility = !this.items[id].visibility
      //そいつのvisibilityをfalseに変更、しかし順番あり
    }
  }
}
</script>
