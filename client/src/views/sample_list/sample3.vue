<template>
  <v-container class="text-center">
    <transition-group name="cell" tag="div" class="touchNumbers d-inline-flex">
      <div v-for="(cell,index) in cells" @click="judgeAndHide(index)" :key="cell.id" class="rounded-circle" :class="{ show: cell.show }">
        {{ cell.number }}
      </div>
    </transition-group>
    <div class="pa-4">
      <v-btn rounded color="success" v-if="nowPlaying" @click="reset">
        Reset
      </v-btn>
      <v-btn rounded color="success" v-else @click="start">
        Start
      </v-btn>
      <div class="pa-2">
        <span :class="{'red--text accent-2':timeup}">{{checkMinutes | zeroPadding}}：{{checkSeconds | zeroPadding}}：{{checkMiliSeconds | showMiliseconds}}</span> / 00：30：000
      </div>
      <br>
<!--      <button @click="Send">Get</button><br>-->
<!--      <button @click="Post">Post</button>-->
      <div v-if="finish">
        <div v-if="!timeup">
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
import axios from 'axios' //追記

export default {
  data:()=>({
    cells: Array.apply(null, { length:25 }).map(function(_, index) {
      return {
        id: index,
        number: index+1,
        show:true
      };
    }),
    openNumCount:0, //消した数字の数
    nowPlaying:false,
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
    shuffle() { //数字のシャッフル、start時に実行
      this.cells = _.shuffle(this.cells);
    },
    deleteNum(id) { //数字を消す。
      this.cells[id].show = false
    },
    judgeAndHide(index) { //数字クリック時に正誤判定を行う
      if(this.nowPlaying == true){
        if (this.cells[index].id == this.openNumCount) {
          this.deleteNum(index)
          this.openNumCount += 1
        }
      }
    },
    start(){ //スタート時、シャッフルとタイマーを起動
      this.shuffle()
      this.timeStart()
      this.nowPlaying = !this.nowPlaying
    },
    resetNumber(){ //数字の再表示
      this.openNumCount = 0
      // cellの全てを再表示する
      for(let cell of this.cells){
        cell.show = true
      }
    },
    reset(){ //ゲームリセット
      this.resetNumber()
      this.timeStop()
      this.timeReset()
      this.nowPlaying = !this.nowPlaying
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
    async Post() { //ゲーム終了時にAPIを叩いて、タイムの記録とランキング表示を行う
      axios.post('http://localhost:8800/record/register', {
        name: "default",
        time: Math.round(this.diffTime), //millisecondで送信
        success: this.clear
      })
          .then(function (response) {
            console.log(response);
          })
          .catch(function (error) {
            console.log(error);
          });
    }
  },
  computed:{
    timeup(){ //timeLimitに間に合うとtrueを返す
      const timeLimit = 30
      if(this.diffTime > timeLimit){
        return true
      }else{
        return false
      }
    },
    finish(){ //ゲームが終わるとtrueを返す
      if(this.openNumCount === 25){
        this.timeStop()
        //後々、名前の入力を受け付けるように仕様変更したい。
        this.Post()
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
