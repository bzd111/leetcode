#
# @lc app=leetcode id=146 lang=python3
#
# [146] LRU Cache
#

# @lc code=start


class DLinkNode:
    def __init__(self, key, val):
        self.key = key
        self.val = val
        self.prev = None
        self.next = None


class LRUCache:
    def __init__(self, capacity: int):
        self.cap = capacity
        self.map = {}
        self.top = DLinkNode(None, -1)
        self.tail = DLinkNode(None, -1)
        self.top.next = self.tail
        self.tail.prev = self.top

    def get(self, key: int) -> int:
        if key in self.map:
            cur = self.map[key]
            # move cur node
            cur.next.prev = cur.prev
            cur.prev.next = cur.next
            # move cur to first
            top = self.top.next
            self.top.next = cur
            cur.prev = self.top
            cur.next = top
            top.prev = cur

            return self.map[key].val
        return -1

    def put(self, key: int, value: int) -> None:
        if key in self.map:
            print("key: ", key)
            cur = self.map[key]
            cur.val = value
            # move cur node
            cur.next.prev = cur.prev
            cur.prev.next = cur.next
            # move cur to first
            top = self.top.next
            self.top.next = cur
            cur.prev = self.top
            cur.next = top
            top.prev = cur
        else:
            cur = DLinkNode(key, value)
            self.map[key] = cur
            top = self.top.next
            self.top.next = cur
            cur.prev = self.top
            cur.next = top
            top.prev = cur
            if len(self.map.keys()) > self.cap:
                self.map.pop(self.tail.prev.key)
                # remove tail Node
                self.tail.prev.prev.next = self.tail
                self.tail.prev = self.tail.prev.prev

    def __repr__(self):
        vals = []
        p = self.top.next
        while p.next:
            vals.append(str(p.val))
            p = p.next
        return "->".join(vals)


# Your LRUCache object will be instantiated and called as such:
# obj = LRUCache(capacity)
# param_1 = obj.get(key)
# obj.put(key,value)
# @lc code=end

if __name__ == "__main__":
    #     ["LRUCache","put","put","get","put","put","get"]
    # ' +
    #   '[[2],[2,1],[2,2],[2],[1,1],[4,1],[2]]
    # Answer
    # [null,null,null,1,null,null,-1]
    # Expected Answer
    # [null,null,null,2,null,null,-1]
    lru = LRUCache(2)
    lru.put(2, 1)
    lru.put(2, 2)
    lru.get(2)
    lru.put(1, 1)
    lru.put(4, 1)
    lru.get(2)
    print(lru)
