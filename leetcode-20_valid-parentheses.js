//检查字符串中的括号是否一一闭合
var isValid = function(s) {
    let d = ['[]', '{}', '()'];
    let stack = [];
    for (let i = 0; i < s.length; i++) {
        stack.push(s[i]);
        if (stack.length >= 2 && d.indexOf(stack[stack.length - 2] + stack[stack.length - 1]) !== -1) { //如果闭合则出栈
            stack.pop();
            stack.pop();
        }
    }
    return stack.length === 0; //如果全部闭合, 那么所有括号都会出栈, 最终栈的长度为0
};

console.log(isValid('()'));
