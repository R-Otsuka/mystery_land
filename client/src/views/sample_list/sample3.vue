<template>
  <v-container>
    <transition-group name="cell" tag="div" class="touchNumbers">
      <div v-for="(cell,index) in cells" @click="judgeAndHide(index)" :key="cell.id" class="cell rounded-circle" :class="{ show: cell.show }">
        {{ cell.number }}
      </div>
    </transition-group>
    <button v-if="playing" @click="reset">
      Reset
    </button>
    <button v-else @click="start">
      Start
    </button>
    <div>
      {{checkHours | zeroPadding}}：{{checkMinutes | zeroPadding}}：{{checkSeconds | zeroPadding}}：{{checkMiliSeconds | showMiliseconds}}
    </div>
    <div v-if="finish">
      <div v-if="clear">
        clear
      </div>
      <div v-else>
        fail
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
      if (this.cells[index].id == this.openNum) {
        this.deleteballon(index)
        this.openNum += 1
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
    checkHours(){
      return Math.floor(this.diffTime / 1000 / 60 / 60);
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
<style>
.touchNumbers {
  display: flex;
  flex-wrap: wrap;
  width: 238px;
  margin-top: 10px;
}
.cell {
  display: flex;
  justify-content: space-around;
  align-items: center;
  width: 45px;
  height: 45px;
  /*border: 1px solid #aaa;*/
  margin-right: -1px;
  margin-bottom: -1px;
}

.cell-move {
  transition: transform 0.6s;
}

.rounded-circle {
  border-radius: 50%;
  width: 45px;
  height: 45px;
  background-color: #ffffff;
  transition:all 0.2s ease-out;
}
.show{
  background-color: #00acc1;
}
</style>
