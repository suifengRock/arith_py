#coding:utf-8
class Stack(object):
    """
    The structure of a Stack.
    The user don't have to know the definition.
    """
    def __init__(self):
        self.__container = list()
    def __is_empty(self):
        """
        Test if the stack is empty or not
        :return: True or False
        """
        return len(self.__container) == 0
    def push(self, element):
        """
        Add a new element to the stack
        :param element: the element you want to add
        :return: None
        """
        self.__container.append(element)
    def top(self):
        """
        Get the top element of the stack
        :return: top element
        """
        if self.__is_empty():
            return None
        return self.__container[-1]
    def pop(self):
        """
        Remove the top element of the stack
        :return: None or the top element of the stack
        """
        return None if self.__is_empty() else self.__container.pop()
    def clear(self):
        """
        We'll make an empty stack
        :return: self
        """
        self.__container.clear()
        return self