#!/usr/bin/env python


def one_find_longest_common_prefix(str_lst=None):
    res = ""
    if not str_lst:
        return res

    for item in zip(*str_lst):
        if len(set(item)) == 1:
            res += item[0]
        else:
            return res


def two_find_longest_common_prefix(str_lst=None):
    res = ""
    if not str_lst:
        return res

    min_str = min(str_lst)
    max_str = max(str_lst)
    for i, char in enumerate(min_str):
        if char != max_str[i]:
            return min_str[:i]
    return str_lst


def three_find_longest_common_prefix(str_lst=None):
    res = ""
    if not str_lst:
        return res

    res_lst = []
    new_str_lst = sorted(str_lst, key=len, reverse=False)
    for i in range(len(new_str_lst[0])):

        char_lst = []
        for j in str_lst:
            char_lst.append(j[i])

        char_lst = sorted(char_lst)
        if char_lst[0] == char_lst[-1]:
            res_lst.append(char_lst[0])

    return ''.join(res_lst)


class TestFuncOne:
    def test_one(self):
        arr = ["flower", "flow", "flight"]
        assert one_find_longest_common_prefix(arr) == "fl"

    def test_two(self):
        arr = ["dog", "racecar", "car"]
        assert one_find_longest_common_prefix(arr) == ""


class TestFuncTwo:
    def test_one(self):
        arr = ["flower", "flow", "flight"]
        assert two_find_longest_common_prefix(arr) == "fl"

    def test_two(self):
        arr = ["dog", "racecar", "car"]
        assert two_find_longest_common_prefix(arr) == ""


class TestFuncThree:
    def test_one(self):
        arr = ["flower", "flow", "flight"]
        assert three_find_longest_common_prefix(arr) == "fl"

    def test_two(self):
        arr = ["dog", "racecar", "car"]
        assert three_find_longest_common_prefix(arr) == ""
