"""
C style string formatting with % operator
"""
name='Yang'
print('Hello, %s!' % name)
# Hello, Yang!

protocol = 7
print('The protocol version is %03d' % protocol)
# The protocol version is 007

name='Yang'
description='a software engineer'
print('%s is %s' % (name, description))
# Yang is a software engineer

"""
Built-in string formatting with str.format()
"""
name='Yang'
print('Hello, {}!'.format(name))
# Hello, Yang!

protocol = 7
print('The protocol version is {:03d}'.format(protocol))
# The protocol version is 007

name='Yang'
description='a software engineer'
print('{} is {}'.format(name, description))
print('{n} is {d}'.format(n=name, d=description))
# Yang is a software engineer

"""
String Literals
String interpolation / f-strings
"""
name='Yang'
description='a software engineer'
print(f'{name} is {description}')
# Yang is a software engineer

protocol = 7
print(f'The protocol version is {protocol:03d}')
# The protocol version is 007

"or even run a function inside the string"
from datetime import datetime
print(f'The current time is {datetime.now()}')
# The current time is 2021-08-21 16:02:05.238000

"""
Template Strings: An Unpopular Way
https://docs.python.org/3/library/string.html#template-strings
"""
from string import Template
t = Template('Hey, $name!')
print(t.substitute(name=name))
# Hey, Yang!



"""
REFERENCE:
• 4 String Formatting Techniques in Python — Feel the Evolution
  - https://medium.com/techtofreedom/4-string-formatting-techniques-in-python-feel-the-evolution-a1b880e36466
"""
