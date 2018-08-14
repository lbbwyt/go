
//https://segmentfault.com/a/1190000012789253#articleHeader7
//var say = require('./util.js');
//say();
import Vue from 'vue';
import './style/common.css';
import getData from './util';

import Ppa from  './App.vue';
//使vue的devtools可用
Vue.config.devtools = true;
//vue 单页组件

new Vue({
	el :'#ppa',
	template:'<Ppa/>',
	components: { Ppa }
});

