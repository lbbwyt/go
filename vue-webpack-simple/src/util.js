 function getData() {
	 //Promise 是一个对象，从它可以获取异步操作的消息
	return new Promise((resolve, reject)=>{
//		if (/* 异步操作成功 */){
//			resolve(value);
//		} else {
//			reject(error);
//		}
	resolve("this ia a  single page vue componnent");

	})
}

export default getData;