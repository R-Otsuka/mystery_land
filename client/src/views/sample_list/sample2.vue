<template>
  <v-container>
    <v-row justify="center" class="ma-5">
      <v-col cols="3" v-for="number in numList1" v-bind:key="number.id">
        <div class="rounded-circle mx-auto balloon">
          <span>{{ number.num }}</span>
        </div>
      </v-col>
    </v-row>
    <v-row justify="center" class="ma-5">
      <v-col cols="3" v-for="number in numList2" v-bind:key="number.id">
        <div class="rounded-circle mx-auto balloon">
          {{ number.num }}
        </div>
      </v-col>
    </v-row>
    <v-row justify="center" class="ma-5">
      <v-col cols="3" v-for="number in numList3" v-bind:key="number.id">
        <div class="rounded-circle mx-auto balloon">
          {{ number.num }}
        </div>
      </v-col>
    </v-row>
    <v-row justify="center" class="ma-5">
      <v-col cols="3" v-for="number in numList4" v-bind:key="number.id">
        <div class="rounded-circle mx-auto balloon">
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
      1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16
    ],
  }),
  computed:{
    numList1(){
      const num1 = this.numbers.slice(0,4);
      let array = []
      for(let i=0;i<4;i++){
        if(num1[i] > 8) {
          array.push({id: num1[i], num: num1[i]-8, visibility: true})
        }else{
          array.push({id: num1[i], num: num1[i], visibility: true})
        }
      }
      return array
    },
    numList2(){
      const num2 = this.numbers.slice(4,8);
      let array = []
      for(let i=0;i<4;i++){
        if(num2[i] > 8) {
          array.push({id: num2[i], num: num2[i]-8, visibility: true})
        }else{
          array.push({id: num2[i], num: num2[i], visibility: true})
        }
      }
      return array
    },
    numList3(){
      const num3 = this.numbers.slice(8,12);
      let array = []
      for(let i=0;i<4;i++){
        if(num3[i] > 8) {
          array.push({id: num3[i], num: num3[i]-8, visibility: true})
        }else{
          array.push({id: num3[i], num: num3[i], visibility: true})
        }
      }
      return array
    },
    numList4(){
      const num4 = this.numbers.slice(12,16);
      let array = []
      for(let i=0;i<4;i++){
        if(num4[i] > 8) {
          array.push({id: num4[i], num: num4[i]-8, visibility: true})
        }else{
          array.push({id: num4[i], num: num4[i], visibility: true})
        }
      }
      return array
    },
  },
  methods:{
    listShuffle() {
      this.numbers = _.shuffle(this.numbers);
    },
  },
  created(){
    this.listShuffle()
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
  background-color:rgba(11,10,120,0.2);
}
</style>