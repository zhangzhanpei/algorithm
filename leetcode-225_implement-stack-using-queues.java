import java.util.LinkedList;
import java.util.Queue;

/**
 * 使用两个队列实现一个栈
 * 入栈：直接追加到输入队列尾部
 * 出栈：将输入队列的元素转移到输出队列中，并返回最后一个元素，即后入先出
 */
class MyStack {

    Queue<Integer> inQueue = new LinkedList<>();
    Queue<Integer> outQueue = new LinkedList<>();

    public MyStack() {

    }

    public void push(int x) {
        inQueue.add(x);
    }

    public int pop() {
        while (inQueue.size() > 1) {
            outQueue.add(inQueue.poll());
        }
        int ret = inQueue.poll();
        while (!outQueue.isEmpty()) {
            inQueue.add(outQueue.poll());
        }
        return ret;
    }

    public int top() {
        while (inQueue.size() > 1) {
            outQueue.add(inQueue.poll());
        }
        int ret = inQueue.poll();
        outQueue.add(ret);
        while (!outQueue.isEmpty()) {
            inQueue.add(outQueue.poll());
        }
        return ret;
    }

    public boolean empty() {
        return inQueue.isEmpty();
    }
}
