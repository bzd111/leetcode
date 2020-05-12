from typing import List


class MyQueue:
    def __init__(self) -> None:
        """
        Initialize your data structure here.
        """
        self.data: List[int] = []
        self.reverse: List[int] = []

    def push(self, x: int) -> None:
        """
        Push element x to the back of queue.
        """
        self.data.append(x)

    def pop(self) -> int:
        """
        Removes the element from in front of queue and returns that element.
        """
        self.reverse = self.data[::-1]
        item = self.reverse.pop()
        self.data = self.reverse[::-1]
        return item

    def peek(self) -> int:
        """
        Get the front element.
        """
        item = self.data[0]
        return item

    def empty(self) -> bool:
        """
        Returns whether the queue is empty.
        """
        return len(self.data) == 0


# Your MyQueue object will be instantiated and called as such:
# obj = MyQueue()
# obj.push(x)
# param_2 = obj.pop()
# param_3 = obj.peek()
# param_4 = obj.empty()
if __name__ == "__main__":
    queue = MyQueue()

    queue.push(1)
    queue.push(2)
    print(queue.peek())  # returns 1
    print(queue.pop())  # // returns 1
    print(queue.empty())  # // returns false
