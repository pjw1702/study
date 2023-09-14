package chapter05_Queue;

import java.util.ArrayList;

public class ArrayQueue<T> {
    private ArrayList<T> queue = new ArrayList<T>();

    public void enqueue(T item) {
        queue.add(item);
    }

    public T dequeue() {
        if(queue.isEmpty()) {
            return null;
        }
        return queue.remove(0); // Array 기반의 큐 이므로, 큐의 맨 앞의 데이터를 먼저 삭제(0번 인덱스)
    }

    public boolean isEmpty() {
        return queue.isEmpty();
    }

    public static void main(String[] args) {
        ArrayQueue<Integer> mq = new ArrayQueue<Integer>();
        mq.enqueue(1);
        mq.enqueue(2);
        mq.enqueue(3);

        System.out.println(mq.dequeue());
        System.out.println(mq.dequeue());
        System.out.println(mq.dequeue());
    }
}