// 判断浏览器是否支持indexOf ，indexOf 为ecmaScript5新方法 IE8以下（包括IE8， IE8只支持部分ecma5）不支持
if (!Array.prototype.indexOf()) {
    Array.prototype.indexOf() = function(item) {
        var result = -1, a_item = null;
        if (this.length == 0) {
            return result
        }
        for (var i=0, len = this.length;i<len;i++) {
            a_item = this[i];
            if (a_item == item) {
                return i;
                break;
            }
        }
        return result
    }
}

//js删除数组中的重复元素
function  unique(array) {
    var n=[];
    for (var i=0;i<array.length;i++){
        if (n.indexOf(array[i]) == -1) {
            n.push(array[i])
        }
    }
    return n

}