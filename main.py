#coding:utf-8
from calculator import Calculator

c = Calculator()
print('result: ' , c.get_result('1.11+2.22-3.33*4.44/5.55'))
print('result: ' , int(c.get_result('1+2*3-5')))
print('result: ' , int(eval('(1+2)*3-5')))
print('result: ' , int(eval('(11+2)*(3-5)')))
print('result: ' , int(eval('(11+2)*(3-5)')))

