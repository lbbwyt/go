<!DOCTYPE html>

<html>
<head>
  <title>school</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
 	<!-- import css -->
	<link rel="stylesheet" href="https://unpkg.com/element-ui/lib/theme-chalk/index.css">	
    <script src="https://unpkg.com/axios/dist/axios.min.js"></script>
</head>
<body>
  <div id="school">
	<el-input v-model="inputValue" 
		v-on:input="inputEvent"
	
	  placeholder="请输入内容"></el-input>	
	
	<ol>
    <li v-for="school in schools">
      {{ school.school_name }}
    </li>
  </ol>
  </div>

	

	<div id="app">
	<el-row class="demo-autocomplete">
	  <el-col :span="12">
	    <div class="sub-title">激活即列出输入建议</div>
	    <el-autocomplete class="inline-input" v-model="state1" :fetch-suggestions="querySearch" placeholder="请输入内容" @select="handleSelect"></el-autocomplete>
	  </el-col>
	  <el-col :span="12">
	    <div class="sub-title">输入后匹配输入建议</div>
	    <el-autocomplete class="inline-input" v-model="state2" :fetch-suggestions="querySearch" placeholder="请输入内容" :trigger-on-focus="false" @select="handleSelect"></el-autocomplete>
	  </el-col>
	</el-row>
	</div>




</body>
 	<!-- import Vue before Element -->
 	 <script src="https://unpkg.com/vue/dist/vue.js"></script>
 	 <!-- import JavaScript -->
 	 <script src="https://unpkg.com/element-ui/lib/index.js"></script>
	
	<script src="../static/js/schoolcontroller.js"></script>
	

</html>
