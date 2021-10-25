#!/usr/bin/env python

class Node(object):
    def __init__(self, item):
        self.item = item
        self.next = None


class SinglyLinkedList(object):
    def __init__(self):
        self._head = None

    @property
    def is_empty(self):
        if self._head:
            return False
        return True

    @property
    def length(self):
        cur = self._head
        n = 0
        if not cur:
            return n

        n = 1
        while cur.next:
            n += 1
            cur = cur.next
        return n

    def ergodic(self):
        result = []
        cur = self._head
        if not cur:
            return result

        while cur.next:
            result.append(cur.item)
            cur = cur.next
        result.append(cur.item)
        return result

    def add(self, item):
        node = Node(item)
        node.next = self._head
        self._head = node

    def append(self, item):
        cur = self._head
        if not cur:
            self.add(item)
        else:
            while cur.next:
                cur = cur.next
            cur.next = Node(item)

    def insert(self, index, item):
        if index == 0:
            self.add(item)
        elif index >= self.length:
            self.append(item)
        else:
            cur = self._head
            n = 1
            while cur.next:
                if n == index:
                    break
                cur = cur.next
                n += 1

            node = Node(item)
            node.next = cur.next
            cur.next = node

    def delete(self, item):
        if self.is_empty:
            raise ValueError("null")

        cur = self._head
        if cur.item == item:
            self._head = cur.next
        else:
            while cur.next:
                pre = cur
                cur = cur.next
                if cur.item == item:
                    pre.next = cur.next

    def remove_nth_from_end(self, n):
        if n < 1 or n > self.length:
            raise ValueError("error")

        if n == self.length == 1:
            self._head = None
            return

        remove_index = self.length - n
        cur = self._head
        index = 0
        while cur.next:
            index += 1
            pre = cur
            cur = cur.next
            if remove_index == index:
                pre.next = cur.next


class TestClass:
    def test_one(self):
        s = SinglyLinkedList()
        for i in map(lambda x: x + 1, list(range(5))):
            s.append(i)

        assert s.ergodic() == [1, 2, 3, 4, 5]

        s.remove_nth_from_end(2)
        assert s.ergodic() == [1, 2, 3, 5]

    def test_two(self):
        s = SinglyLinkedList()
        for i in map(lambda x: x + 1, list(range(1))):
            s.append(i)
        assert s.ergodic() == [1]

        s.remove_nth_from_end(1)
        assert s.ergodic() == []

    def test_three(self):
        s = SinglyLinkedList()
        for i in map(lambda x: x + 1, list(range(2))):
            s.append(i)
        assert s.ergodic() == [1, 2]

        s.remove_nth_from_end(1)
        assert s.ergodic() == [1]
