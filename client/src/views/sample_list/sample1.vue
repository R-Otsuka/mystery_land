<template>
  <v-container>
    <div class="mx-auto text-center py-5 circle-top">
      <transition name="bounce">
        <div class="rounded-circle mx-auto"
              :class=items[0].color
              @click="judgeAndHide(0)"
              v-show="items[0].visibility"
        >
        </div>
      </transition>
    </div>
    <v-row justify="space-around" class="py-5 circle-center">
      <v-col cols="3" class="text-sm-left">
        <transition name="bounce">
          <div class="rounded-circle"
                v-bind:class=items[1].color
                @click="judgeAndHide(1)"
                v-show="items[1].visibility"
          >
          </div>
        </transition>
      </v-col>
      <v-col cols="3" class="text-right">
        <transition name="bounce">
          <div class="rounded-circle float-right"
                v-bind:class=items[2].color
                @click="judgeAndHide(2)"
                v-show="items[2].visibility"
          >
          </div>
        </transition>
      </v-col>
    </v-row>
<!--    クリアすると点字->パスワードに表示が変わる-->
    <p v-if="clear" class="mx-auto text-center">
      クリア！<br>
      <router-link class="routerLink" to="/sample/2">
        <v-btn rounded color="success" class="text-center margin">
          次へ
        </v-btn>
      </router-link>
    </p>
    <p v-else class="mx-auto text-center tenji">⠊⠛⠴⠐⠳⠽⠑⠱⠣⠐⠩⠛⠒⠊⠊⠷⠐⠞⠓</p>
    <v-row justify="center" class="circle-bottom">
      <v-col cols="3" class="text-sm-left">
        <transition name="bounce">
          <div
                class="rounded-circle"
                v-bind:class=items[3].color
                @click="judgeAndHide(3)"
                v-show="items[3].visibility"
          ></div>
        </transition>
      </v-col>
      <v-col cols="3" class="text-right">
        <transition name="bounce">
          <div class="rounded-circle float-right"
                v-bind:class=items[4].color
                @click="judgeAndHide(4)"
                v-show="items[4].visibility"
          >
          </div>
        </transition>
      </v-col>
    </v-row>
    <br>
  </v-container>
</template>
<script>
import _ from 'lodash';
export default {
  data:()=> ({
    //次に消すべきボールのid
    BalloonIdToClickNext:0,
    //各ボールのhash情報
    items:[
      {id:0,color:'warning',visibility:true},
      {id:1,color:'purple',visibility:true},
      {id:2,color:'grey',visibility:true},
      {id:3,color:'primary',visibility:true},
      {id:4,color:'success',visibility:true}
    ]
  }),
  computed:{
    //クリア時にtrueを返す
    clear(){
      if(this.BalloonIdToClickNext==5){
        console.log("clear")
        return true
      }else{
        return false
      }
    }
  },
  methods: {
    //ボールをシャッフルする。(vueインスタンスcreate時に実行)
    listShuffle() {
      this.items = _.shuffle(this.items);
    },
    //ボールをクリックした時に実行される。
    //クリック順が正しければ消す、そうでなければリセットし初期状態に戻す。
    judgeAndHide(id){
      if(this.items[id].id == this.BalloonIdToClickNext) {
        if (this.items[id].visibility == true) {
          this.items[id].visibility = !this.items[id].visibility
        }
        this.BalloonIdToClickNext += 1
      }else if(this.items[id].id > this.BalloonIdToClickNext){
        for(let i=0;i<this.items.length;i++){
          this.items[i].visibility = true
        }
        this.BalloonIdToClickNext = 0
        this.listShuffle()
      }
    },
    next(){

    }
  },
  created() {
    this.listShuffle()
  }
}
</script>
<style scoped>
.tenji{
  padding-top:10px;
  padding-bottom:15px;
}

.rounded-circle {
  border-radius: 50%;
  width: 40px;
  height: 40px;
  background: #000;
}
.circle-top {
  height:50px;
  margin:10px;
}
.circle-center {
  height: 60px;
  margin: 10px;
}
.circle-bottom {
  height: 60px;
  margin: 10px;
}

@media screen and (min-width:480px) {
  /*画面サイズが480pxからはここを読み込む*/
  .rounded-circle {
    border-radius: 50%;
    width: 50px;
    height: 50px;
    background: #000;
  }
  .circle-top {
    height:80px;
    margin:20px;
  }
  .circle-center {
    height: 90px;
    margin: 20px;
  }
  .circle-bottom {
    height: 90px;
    margin:20px;
  }
  .tenji{
    padding-top:10px;
    padding-bottom:30px;
  }

}
@media screen and (min-width:1024px) {
   /*画面サイズが1024pxからはここを読み込む */
  .rounded-circle {
    border-radius: 50%;
    width: 64px;
    height: 64px;
    background: #000;
  }
  .circle-top {
    height:100px;
    margin:20px;
  }
  .circle-center {
    height: 120px;
    margin: 20px;
  }
  .circle-bottom {
    height: 120px;
    margin:20px;
  }
  .tenji{
    padding-top:10px;
    padding-bottom:50px;
  }
}

.bounce-enter-active {
  animation: bounce-in 0.5s;
}
.bounce-leave-active {
  animation: bounce-in 0.5s reverse;
}
@keyframes bounce-in {
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
</style>
