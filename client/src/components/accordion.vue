<template>
  <div>
    <div v-on:click="toggle" v-bind:class="[active? 'close':'open','togglebutton']">{{item.title}}</div>
    <slide-up-down :active="active">
      <div class="element"><span v-html="item.text"></span></div>
    </slide-up-down>
  </div>
</template>

<script>
export default {
  data(){
    return {
      active:false
    }
  },
  props:['item'],
  methods:{
    toggle () {
      this.active = !this.active;
    }
  }
}
</script>

<style scoped>
.togglebutton{
  padding:1em;
  background-color:#2196F3;
  color:#FFF;
  cursor:pointer;
  position:relative;
  border-top:1px solid #FFF;
}
.togglebutton:after{
  position:absolute;
  right:0.5em;
  top:50%;
  transform:translateY(-50%);
  font-size:80%;
}
.togglebutton.close:after{
  content:"△close";
}
.togglebutton.open:after{
  content:"▽open";
}

.element {
  padding : 1em;
  background-color:#B3E5FC;
}
</style>