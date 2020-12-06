<template>
  <v-container>
    <transition-group name="cell" tag="div" class="touchNumbers">
      <div v-for="(cell,index) in cells" @click="judgeAndHide(index)" :key="cell.id" class="cell rounded-circle" :class="{ show: cell.show }">
        {{ cell.number }}
      </div>
    </transition-group>
    <button @click="start">
      Start
    </button>
    <button @click="shuffle">
      Shuffle
    </button>
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
    openNum:0
  }),
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
    timerStart(){

    },
    start(){
      this.shuffle()
      this.timerStart()
    }
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
