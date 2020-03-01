import java.util.Stack;

/**
 * 使用两个栈实现一个队列
 * 进队：先将输出栈的元素放回输入栈，此时输入栈的顺序即入队顺序
 * 出队：输出栈此时为空，将输入栈的元素放到输出栈，此时栈顶元素即输入栈的最前面那个元素，达到输入栈的元素先入先出
 */
class MyQueue {

    Stack<Integer> inStack = new Stack<>();
    Stack<Integer> outStack = new Stack<>();

    public MyQueue() {

    }

    public void push(int x) {
        while (!outStack.isEmpty()) {
            inStack.push(outStack.pop());
        }
        inStack.push(x);
    }

    public int pop() {
        while (!inStack.isEmpty()) {
            outStack.push(inStack.pop());
        }
        return outStack.pop();
    }

    public int peek() {
        while (!inStack.isEmpty()) {
            outStack.push(inStack.pop());
        }
        return outStack.peek();
    }

    public boolean empty() {
        while (!inStack.isEmpty()) {
            outStack.push(inStack.pop());
        }
        return outStack.isEmpty();
    }
}
