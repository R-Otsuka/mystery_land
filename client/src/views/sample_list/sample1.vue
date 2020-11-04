<template>
<!--  v-contentタグはサイドバー表示の際にpadding-leftの自動追加機能あり??-->
<!--  <v-content class="py-10">-->
  <v-container>
    <div class="mx-auto text-center py-5">
      <transition>
        <span class="rounded-circle d-inline-flex"
              :class=items[0].color
              @click="hide(0)"
              v-show="items[0].visibility"
              :style="{opacity: items[0].visibility ? 1 : 0}"
        >
        </span>
      </transition>
    </div>
<!--      start,center,など設定によりseparateを変えられる(公式参照)-->
    <v-row justify="space-around" class="py-5">
      <v-col cols="3" class="text-sm-left">
        <transition>
          <span class="rounded-circle d-inline-flex"
                v-bind:class=items[1].color
                @click="hide(1)"
                v-show="items[1].visibility"
                :style="{opacity: items[1].visibility ? 1 : 0}"
          >
          </span>
        </transition>
      </v-col>
      <v-col cols="3" class="text-sm-right">
        <transition>
          <span class="rounded-circle d-inline-flex"
                v-bind:class=items[2].color
                @click="hide(2)"
                v-show="items[2].visibility"
                :style="{opacity: items[2].visibility ? 1 : 0}"
          >
          </span>
        </transition>
      </v-col>
    </v-row>
    <p class="mx-auto text-center pb-10" style="width: 200px;">⠁⠡⠁⠊⠩⠚⠷⠐⠞⠓</p>
    <v-row justify="center" class="py-5">
      <v-col cols="3" class="text-sm-left">
        <transition>
          <span class="rounded-circle d-inline-flex"
                v-bind:class=items[3].color
                @click="hide(3)"
                v-show="items[3].visibility"
                :style="{opacity: items[3].visibility ? 1 : 0}"
          >
          </span>
        </transition>
      </v-col>
      <v-col cols="3" class="text-sm-right">
        <transition>
          <span class="rounded-circle d-inline-flex"
                v-bind:class=items[4].color
                @click="hide(4)"
                v-show="items[4].visibility"
                :style="{opacity: items[4].visibility ? 1 : 0}"
          >
          </span>
        </transition>
      </v-col>
    </v-row>
    <button @click="listShuffle">shuffle</button>
    <br>
  </v-container>
<!--  </v-content>-->
</template>
<script>
import _ from 'lodash';
export default {
  data:()=> ({
    setfunc:false,
    items:[
      {id:0,color:'warning',visibility:true},
      {id:1,color:'purple',visibility:true},
      {id:2,color:'grey',visibility:true},
      {id:3,color:'primary',visibility:true},
      {id:4,color:'success',visibility:true}
    ]
  }),
  methods: {
    listShuffle() {
      this.items = _.shuffle(this.items);
    },
    hide(id){
      if(this.items[id].visibility == true) {
        this.items[id].visibility = !this.items[id].visibility
        console.log(this.items[id].visibility)
      }
    },
  }
}
</script>
<style>
.rounded-circle {
  border-radius: 50%;
  width: 64px;
  height: 64px;
  background: #000;
}
.v-leave-active{
  transition: opacity 1s;
}
.v-enter,
.v-leave-to {
  opacity: 0;
}
@keyframes v-in {
  0% {
    transform: scale(0);
  }
  50% {
    transform: scale(1.5);
  }
  100% {
    transform: scale(1);
  }
}
@keyframes v-out {
  0% {
    transform: scale(1);
  }
  50% {
    transform: scale(1.5);
  }
  100% {
    transform: scale(0);
  }
}
</style>
