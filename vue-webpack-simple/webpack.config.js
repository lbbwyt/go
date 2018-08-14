const path = require('path');
var webpack = require('webpack');
const VueLoaderPlugin = require('vue-loader/lib/plugin');

module.exports = {
	entry : ['babel-polyfill','./src/main.js'],
	output : {
		filename:'build.js',
		path: path.resolve(__dirname, 'dist'),
		publicPath: '/dist/' // 通过devServer访问路径
		
	},
	
	plugins: [
        // make sure to include the plugin for the magic
        new VueLoaderPlugin()
    ],
	//JavaScript 提供了 source map 功能，将编译后的代码映射回原始源代码。如果一个错误来自于 b.js，source map 就会明确的告诉你。
	devtool: 'inline-source-map',
	
	//代码发生变化后自动编译代码,同时需在package,json的scripts中添加"dev": "webpack-dev-server --open --hot",
	devServer: {
        historyApiFallback: true,
        overlay: true
    },
	resolve : {
		alias:{
			'vue$':'vue/dist/vue.esm.js'
		}
	},
	module:{
		rules:[
			{
				test :/\.css$/,
				use:[
				'vue-style-loader',
				'css-loader'
				]
			},
			{
				test :/\.js$/,
				loader:'babel-loader',
				//exclude表示忽略node_modules文件夹下的文件，不用转码
				exclude:/node_modules/
			},
			{
				test: /\.(png|jpg|gif|svg)$/,
				loader: 'file-loader',
				options: {
					name: '[name].[ext]?[hash]'
				}
			},
			{
				test: /\.vue$/,
				loader: 'vue-loader',
				options: {
					loaders: {
						'scss': [
							'vue-style-loader',
							'css-loader',
							'sass-loader'
						],
						'sass': [
							'vue-style-loader',
							'css-loader',
							'sass-loader?indentedSyntax'
						]
					}
				}
			}
		]
	}
};