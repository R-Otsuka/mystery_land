<template>
  <div id="app">
    <button @click="toggleBtn">表示／非表示ボタン</button>
    <transition name="simple">
      <p v-show="show" style="width:150px;height:150px;background-color:green">本日は晴天なり</p>
    </transition>
    <button @click="count++">+ju</button>
    <transition name="datachange">
      <div :key="count">{{ count }}</div>
    </transition>
    <button @click="order=!order">切り替え</button>
    <transition-group name="group" tag="ul" class="list">
      <li v-for="item in sortedList" v-bind:key="item.id">
        {{ item.name }} {{ item.price }}円
      </li>
    </transition-group>
    <button @click="change=!change">切り替え</button>
    <transition name="hook" v-on:enter="enter" v-on:after-enter="afterEnter">
      <div v-if="change">example</div>
    </transition>
  </div>
</template>

<script>
import _ from 'lodash';
export default {
  data:()=> ({
    show: true,
    change:true,
    order: false,
    count: 0,
    list:[
      {id:1,name:'りんご',price:100},
      {id:2,name:'ばなな',price:200},
      {id:3,name:'いちご',price:300}
    ]
  }),
  methods: {
    toggleBtn: function () {
      this.show = !this.show
    },
    enter:function(done){
      console.log('enter')
      setTimeout(done,1000)
    },
    afterEnter:function(){
      console.log('afterEnter')
    }
  },
  computed:{
    sortedList:function(){
      return _.orderBy(this.list,'price',this.order ? 'desc':'asc')
    }
  }
}
</script>
<style>
.simple-enter-active,.simple-leave-active {
  transition:opacity 1s;
}

.simple-enter,.simple-leave-to {
  opacity:0;
}
.datachange-enter-active,.datachange-leave-active {
  transition:opacity 1s;
}
.datachange-leave-active{
  position:absolute;
}
.datachange-enter,.datachange-leave-to{
  opacity:0;
}
.group-move{
  transition:transform 1s;
}


</style>