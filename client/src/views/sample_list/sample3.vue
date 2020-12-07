<template>
  <v-container class="text-center">
    <transition-group name="cell" tag="div" class="touchNumbers d-inline-flex">
      <div v-for="(cell,index) in cells" @click="judgeAndHide(index)" :key="cell.id" class="rounded-circle" :class="{ show: cell.show }">
        {{ cell.number }}
      </div>
    </transition-group>
    <div class="pa-4">
      <v-btn rounded color="success" v-if="playing" @click="reset">
        Reset
      </v-btn>
      <v-btn rounded color="success" v-else @click="start">
        Start
      </v-btn>
      <div class="pa-2">
        <span :class="{'red--text accent-2':timeup}">{{checkMinutes | zeroPadding}}：{{checkSeconds | zeroPadding}}：{{checkMiliSeconds | showMiliseconds}}</span> / 00：25：000
      </div>
      <div v-if="finish">
        <div v-if="clear">
          clear
        </div>
        <div v-else>
          fail
        </div>
      </div>
    </div>
  </v-container>
</template>

<script>
import _ from 'lodash';
export default {
  data:()=>({
    cells: Array.apply(null, { length:25 }).map(function(_, index) {
      return {
        id: index,
        number: index+1,
        show:true
      };
    }),
    openNum:0,
    playing:false,
    animationId: 0,
    hours: 0,
    minutes: 0,
    second: 0,
    milisecond: 0,
    startTime: 0,// スタート時間
    endTime: 0,// ストップ押した時間
    diffTime: 0, //スタートとストップ押した時の差分
    flag: false,
  }),
  filters:{
    zeroPadding(value){
      return value.toString().padStart(2, 0);
    },
    showMiliseconds(value){
      return value.toString().padStart(3, 0);
    }
  },
  methods: {
    shuffle() {
      this.cells = _.shuffle(this.cells);
    },
    deleteballon(id) {
      this.cells[id].show = false
    },
    judgeAndHide(index) {
      if(this.playing == true){
        if (this.cells[index].id == this.openNum) {
          this.deleteballon(index)
          this.openNum += 1
        }
      }
    },
    start(){
      this.shuffle()
      this.timeStart()
      this.playing = !this.playing
    },
    resetNumber(){
      this.openNum = 0
      // cellの全てをshowにする
      for(var num of this.cells){
        num.show = true
      }
    },
    reset(){
      this.resetNumber()
      this.timeStop()
      this.timeReset()
      this.playing = !this.playing
    },
    setStartTime(time){
      //performance.now()自体は前回のページから今回のページへと、遷移を開始した瞬間からの経過時間を算出する
      this.startTime = performance.now() - time;
    },
    timeStart(){
      if(this.flag){
        return false;
      }
      const vm_data = this;
      this.flag = true; // スタートしたらタイマーflagをtrueに
      this.setStartTime(vm_data.diffTime);
      (function progress(){
        vm_data.endTime = performance.now();
        vm_data.diffTime = vm_data.endTime - vm_data.startTime;
        [vm_data.second, vm_data.milisecond] = [Math.floor(vm_data.diffTime / 1000), Math.floor(vm_data.diffTime % 1000)];
        vm_data.animationId = window.requestAnimationFrame(progress);
      }());
    },
    timeStop(){
      this.flag = false;   // ストップしたらタイマーflagをfalseに
      window.cancelAnimationFrame(this.animationId);
    },
    timeReset(){
      console.log("reset")
      if(this.flag){
        return false;
      }
      this.startTime = this.diffTime = 0;
    },
  },
  computed:{
    timeup(){
      if(this.diffTime > 25000){
        return true
      }else{
        return false
      }
    },
    finish(){
      if(this.openNum === 25){
        this.timeStop()
        return true
      }else{
        return false
      }
    },
    clear() {
      if (this.diffTime < 25000) {
        return true
      }else{
        return false
      }
    },
    checkMinutes(){
      return Math.floor(this.diffTime / 1000 / 60) % 60;
    },
    checkSeconds(){
      return Math.floor(this.diffTime / 1000) % 60;
    },
    checkMiliSeconds(){
      return Math.floor(this.diffTime % 1000);
    },
  }
}
</script>
<style scoped>

.touchNumbers {
  display: flex;
  flex-wrap: wrap;
  width:300px;
  margin-top: 10px;
}
.rounded-circle {
  display: flex;
  justify-content: space-around;
  align-items: center;
  border-radius: 50%;
  width: 55px;
  height: 55px;
  margin-right: 5px;
  margin-bottom: 5px;
  background-color: #ffffff;
  transition:all 0.2s ease-out;
}

@media screen and (min-width:480px) {
  .touchNumbers {
    display: flex;
    flex-wrap: wrap;
    width:480px;
    margin-top: 10px;
  }
  .rounded-circle {
    display: flex;
    border-radius: 50%;
    width: 80px;
    height: 80px;
    margin-right: 10px;
    margin-bottom: 10px;
    background-color: #ffffff;
    transition:all 0.2s ease-out;
  }
}
@media screen and (min-width:1024px) {
  .touchNumbers {
    display: flex;
    flex-wrap: wrap;
    width:600px;
    margin-top: 10px;
  }
  .rounded-circle {
    display: flex;
    justify-content: space-around;
    align-items: center;
    border-radius: 50%;
    width: 90px;
    height: 90px;
    margin-right: 15px;
    margin-bottom: 15px;
    background-color: #ffffff;
    transition:all 0.2s ease-out;
  }
}
.cell-move {
  transition: transform 0.6s;
}
.show{
  background-color: #00acc1;
}
</style>
