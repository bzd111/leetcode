import sys
from typing import List


class MinStack:
    def __init__(self) -> None:
        """
        initialize your data structure here.
        """
        self.stack: List[int] = []
        self.min = sys.maxsize

    def push(self, x: int) -> None:
        if self.min >= x:
            self.stack.append(self.min)
            self.min = x
        self.stack.append(x)

    def pop(self) -> int:
        t = self.stack.pop()
        if self.min == t:
            self.min = self.stack.pop()
        return t

    def top(self) -> int:
        return self.stack[-1]

    def getMin(self) -> int:
        return self.min


if __name__ == "__main__":
    m = MinStack()
    m.push(-2)
    m.push(0)
    m.push(-3)
    print(m.getMin())
    m.pop()
    print(m.top())
    print(m.getMin())
    # MinStack minStack = new MinStack();
    # minStack.push(-2);
    # minStack.push(0);
    # minStack.push(-3);
    # minStack.getMin();   --> Returns -3.
    # minStack.pop();
    # minStack.top();      --> Returns 0.
    # minStack.getMin();   --> Returns -2.
