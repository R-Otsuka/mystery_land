<template>
  <v-container class="text-center">
    <v-row justify="center" class="ma-5">
      <v-col cols="3" v-for="(number, index) in numRow1" v-bind:key="number.id">
        <div class="rounded-circle mx-auto balloon"
             :class="[number.visibility ? 'trans':'untrans']"
             @click="judgeAndHide(0, index, number.num)">
          {{ number.num }}
        </div>
      </v-col>
    </v-row>
    <v-row justify="center" class="ma-5">
      <v-col cols="3" v-for="(number, index) in numRow2" v-bind:key="number.id">
        <div class="rounded-circle mx-auto balloon"
             :class="[number.visibility ? 'trans':'untrans']"
             @click="judgeAndHide(1, index, number.num)">
          {{ number.num }}
        </div>
      </v-col>
    </v-row>
    <v-row justify="center" class="ma-5">
      <v-col cols="3" v-for="(number, index) in numRow3" v-bind:key="number.id">
        <div class="rounded-circle mx-auto balloon"
             :class="[number.visibility ? 'trans':'untrans']"
             @click="judgeAndHide(2, index, number.num)">
          {{ number.num }}
        </div>
      </v-col>
    </v-row>
    <v-row justify="center" class="ma-5">
      <v-col cols="3" v-for="(number, index) in numRow4" v-bind:key="number.id">
        <div class="rounded-circle mx-auto balloon"
             :class="[number.visibility ? 'trans':'untrans']"
             @click="judgeAndHide(3, index, number.num)">
          {{ number.num }}
        </div>
      </v-col>
    </v-row>
    <div v-if="clear" >
      クリア！<br>
      <router-link class="routerLink" to="/sample/3">
        <v-btn rounded color="success" class="text-center margin">
          次へ
        </v-btn>
      </router-link>
    </div>
  </v-container>
</template>

<script>
import _ from 'lodash';
export default {
  data:()=> ({
    length:4,
    hit:0,
    opened:false,
    openedNum:null,
    openedNumIndex:null,
    numbers:[
      {id:0,num:0,visibility:true},
      {id:1,num:1,visibility:true},
      {id:2,num:2,visibility:true},
      {id:3,num:3,visibility:true},
      {id:4,num:4,visibility:true},
      {id:5,num:5,visibility:true},
      {id:6,num:6,visibility:true},
      {id:7,num:7,visibility:true},
      {id:8,num:0,visibility:true},
      {id:9,num:1,visibility:true},
      {id:10,num:2,visibility:true},
      {id:11,num:3,visibility:true},
      {id:12,num:4,visibility:true},
      {id:13,num:5,visibility:true},
      {id:14,num:6,visibility:true},
      {id:15,num:7,visibility:true}
    ]
  }),
  computed:{
    numRow1(){
      let numRow = this.numbers.slice(0,4);
      return numRow
    },
    numRow2(){
      let numRow = this.numbers.slice(4,8);
      return numRow
    },
    numRow3(){
      let numRow = this.numbers.slice(8,12);
      return numRow
    },
    numRow4(){
      const numRow = this.numbers.slice(12,16);
      return numRow
    },
    clear(){
      if(this.hit == 16){
        return true
      }else{
        return false
      }
    }
  },
  methods:{
    numberShuffle() {
      this.numbers = _.shuffle(this.numbers);
    },
    judgeAndHide(id, index, num){
      //クリックした風船の配列内順位
      const arrayNo = id*4+index
      if(this.opened == false) {
        //1ターン目は残っている風船しかクリックできないようにする。
        if(this.numbers[arrayNo].visibility == true) {
          this.deleteBalloon(arrayNo)
          this.opened = true
          this.openedNum = num
          this.openedNumIndex = arrayNo
        }
      }else{
        if(this.openedNumIndex == arrayNo){
          //同じ場所を二回クリック。風船を復活させる
          this.reviveBalloon(arrayNo);
        }else if(this.openedNum == num){
          //hit
          this.deleteBalloon(arrayNo)
          this.hit += 2
        }else{
          //外れ
          this.deleteBalloon(arrayNo)
          setTimeout(this.reviveBalloon, 500, arrayNo);
          setTimeout(this.reviveBalloon, 500, this.openedNumIndex);
        }
        this.opened = false
        this.openedNum = null
        this.openedNumIndex = null
      }
    },
    deleteBalloon:function(index){
      //クリック箇所の背景色を透明に。classの追加かな
      this.numbers[index].visibility = false
    },
    reviveBalloon:function(index){
      //クリック箇所の背景色を戻す
      this.numbers[index].visibility = true
      console.log("完了")
      console.log(index)
    }
  },

  created(){
    this.numberShuffle()
    let numarray = []
    for(let i=0;i<16;i++){
      numarray.push({id:i,num:i,visibility:true})
    }
    return numarray
  }
}
</script>
<style scoped>
.balloon{
  align-items:center;
  display:flex;
  justify-content:center;
}
.rounded-circle {
  border-radius: 50%;
  width: 50px;
  height: 50px;
  transition: all 500ms 0s ease;
}
.trans{
  background-color:rgba(0,0,0,1);
}
.untrans{

}

</style>