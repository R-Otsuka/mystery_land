<template>
  <v-container>
    <v-row justify="center" class="ma-5">
      <v-col cols="3" v-for="(number, index) in numRow1" v-bind:key="number.id">
        <div class="rounded-circle mx-auto balloon"
             :class="[number.visibility ? 'trans':'untrans']"
             @click="showNumber(0, index)">
          {{ number.num }}
        </div>
      </v-col>
    </v-row>
    <v-row justify="center" class="ma-5">
      <v-col cols="3" v-for="(number, index) in numRow2" v-bind:key="number.id">
        <div class="rounded-circle mx-auto balloon"
             :class="[number.visibility ? 'trans':'untrans']"
             @click="showNumber(1, index)">
          {{ number.num }}
        </div>
      </v-col>
    </v-row>
    <v-row justify="center" class="ma-5">
      <v-col cols="3" v-for="(number, index) in numRow3" v-bind:key="number.id">
        <div class="rounded-circle mx-auto balloon"
             :class="[number.visibility ? 'trans':'untrans']"
             @click="showNumber(2, index)">
          {{ number.num }}
        </div>
      </v-col>
    </v-row>
    <v-row justify="center" class="ma-5">
      <v-col cols="3" v-for="(number, index) in numRow4" v-bind:key="number.id">
        <div class="rounded-circle mx-auto balloon"
             :class="[number.visibility ? 'trans':'untrans']"
             @click="showNumber(3, index)">
          {{ number.num }}
        </div>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import _ from 'lodash';
export default {
  data:()=> ({
    length:4,
    numbers:[
      {id:1,num:1,visibility:true},
      {id:2,num:2,visibility:true},
      {id:3,num:3,visibility:true},
      {id:4,num:4,visibility:true},
      {id:5,num:5,visibility:true},
      {id:6,num:6,visibility:true},
      {id:7,num:7,visibility:true},
      {id:8,num:8,visibility:true},
      {id:9,num:1,visibility:true},
      {id:10,num:2,visibility:true},
      {id:11,num:3,visibility:true},
      {id:12,num:4,visibility:true},
      {id:13,num:5,visibility:true},
      {id:14,num:6,visibility:true},
      {id:15,num:7,visibility:true},
      {id:16,num:8,visibility:true}
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
  },
  methods:{
    numberShuffle() {
      this.numbers = _.shuffle(this.numbers);
    },
    showNumber(id, index){
      //クリック箇所の背景色を透明に。classの追加かな
      const No = id*4+index
      this.numbers[No].visibility = !this.numbers[No].visibility
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
}
.trans{
  background-color:rgba(0,0,0,0.1);
}
.untrans{

}

</style>