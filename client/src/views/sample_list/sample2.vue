<template>
  <div>
  <!--    16個数の風船-->
    <div>
<!--      <transition-group name="cell" tag="div" class="container">-->
<!--        <div v-for="cell in cells" :key="cell.id" class="cell">-->
<!--          {{ cell.number }}-->
<!--        </div>-->
<!--      </transition-group>-->
<!--      <div v-for="number in numbers1"-->
<!--           v-bind:key="number.id"-->
<!--           class="rounded-circle text-center float-left"-->
<!--           @click="judgeAndHide(number.id)"-->
<!--           v-show="number.visibility">-->
<!--        {{ number.num }}-->
<!--      </div>-->
      <v-row justify="center">
        <v-col cols="3" v-for="number in numbers1" v-bind:key="number.id">
          <div class="rounded-circle mx-auto balloon">
            <span>{{ number.num }}</span>
          </div>
        </v-col>
      </v-row>
      <v-row justify="center">
        <v-col cols="3" v-for="number in numbers2" v-bind:key="number.id">
          <div class="rounded-circle mx-auto balloon">
            {{ number.num }}
          </div>
        </v-col>
      </v-row>
      <v-row justify="center">
        <v-col cols="3" v-for="number in numbers3" v-bind:key="number.id">
          <div class="rounded-circle mx-auto balloon">
            {{ number.num }}
          </div>
        </v-col>
      </v-row>
      <v-row justify="center">
        <v-col cols="3" v-for="number in numbers4" v-bind:key="number.id">
          <div class="rounded-circle mx-auto balloon">
            {{ number.num }}
          </div>
        </v-col>
      </v-row>
    </div>
  </div>
</template>

<script>
// import _ from 'lodash';
export default {
  data:()=> ({
    length:4,
    // numbers:[
    //   1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16
    // ],
    numbers1:[
      {id:1,num:1,visibility:true},
      {id:2,num:2,visibility:true},
      {id:3,num:3,visibility:true},
      {id:4,num:4,visibility:true},
    ],
    numbers2:[
      {id:5,num:5,visibility:true},
      {id:6,num:6,visibility:true},
      {id:7,num:7,visibility:true},
      {id:8,num:8,visibility:true},
    ],
    numbers3:[
      {id:9,num:9,visibility:true},
      {id:10,num:10,visibility:true},
      {id:11,num:11,visibility:true},
      {id:12,num:12,visibility:true},
    ],
    numbers4:[
      {id:13,num:13,visibility:true},
      {id:14,num:14,visibility:true},
      {id:15,num:15,visibility:true},
      {id:16,num:16,visibility:true},
    ],
  }),
  methods:{
    judgeAndHide(id){
      this.numbers1[id-1].visibility = false
    }
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
  width: 64px;
  height: 64px;
  background-color:rgba(11,10,120,0.2);
}
</style>