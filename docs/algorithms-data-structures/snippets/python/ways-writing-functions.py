"""
LEVELS OF WRITING PYTHON FUNCTIONS

reference:
  - https://python.plainenglish.io/10-levels-of-writing-python-functions-a0a29997ae79
"""

"""
FUNCTIONS THAT TAKE IN ZERO PARAMETERS
"""
def hello():
  print("hello")
hello() # hello


"""
FUNCTIONS THAT TAKE IN 1 OR MORE PARAMETERS
"""
def add10(a):
  # a function that takes in 1 parameter:
  return a + 10
x = add10(7) # 17

def add(a, b):
  # a function that takes in 2 parameters:
  return a + b
x = add(4, 5) # 9

"""
FUNCTIONS THAT TAKE IN DEFAULT PARAMETERS

Here, if we choose not to pass in the `greeting` argument, it defaults to
`"hello"` as defined in our function definition. Conversely, if we choose to
pass in the `greeting` argument, it takes the value we pass in instead of the
default value.
"""
def greet(name, greeting="hello"):
  print(greeting, name)
greet("bob")  # hello bob
greet("bob", greeting="hi")  # hi bob


"""
FUNCTIONS THAT TAKE IN A VARIABLE NUMBER OF PARAMETERS

By adding a `*` in front of the parameter `names`, we tell Python that `names`
should be a tuple that catches and stores any additional positional arguments.
For instance, if we pass in 3 positional arguments, `names` will be a tuple
containing these 3 arguments.
"""
def greet(*names):
  print("hello", names)
greet("apple")
# hello ('apple',)
greet("apple", "bob", "cat")
# hello ('apple', 'bob', 'cat')


"""
FUNCTIONS THAT TAKE IN A VARIABLE NUMBER OF KEYWORD ARGUMENTS

Similarly, if we put `**` before `kwargs`, we are telling Python that `kwargs`
should be a dictionary that catches and stores multiple keyword arguments in
the form of key-value pairs.
"""
def test(**kwargs):
  print(kwargs)
test()  # {}
test(a=4)  # {"a":4}
test(a=4, b=5, c="hello")  # {"a":4, "b":5, "c":"hello"}


"""
LAMBDA FUNCTIONS

A lambda function is a small anonymous function. We call it anonymous because
it is optional to give it a name. The syntax:

```
lambda inputs: output
```
"""

def add10(a):
  return a + 10
# THIS IS THE SAME AS:
add10 = lambda a: a+10
'''
The function inputs go before the colon, while the output (whatever the
function returns) goes after the colon. For instance:
'''

# another example:
def add(a, b):
  return a + b
add = lambda a,b: a+b
'''
Here, notice that we are giving the lambda functions a name eg. `add10 =
lambda ...`, `add = lambda ...`. Note that this is optional.
'''


"""
FUNCTIONS THAT CALL THEMSELVES (recursive functions)


A recursive function is a function that calls itself inside of its function
body. Here, we have a function `factorial`. Notice that inside the function body,
the `factorial` function itself is called. This makes it a recursive function.
```
factorial(5) = 5 * factorial(4)
             = 5 * (4 * factorial(3))
             = 5 * (4 * (3 * factorial(2)))
             = 5 * (4 * (3 * (2 * factorial(1))))
             = 5 * (4 * (3 * (2 * 1)))
             = 120
```
"""
def factorial(n):
  if n == 1:
    return 1
  return n * factorial(n-1)
factorial(1)  # 1
factorial(2)  # 2
factorial(3)  # 6
factorial(4)  # 24
factorial(5)  # 120

"""
FUNCTIONS THAT TAKES IN OTHER FUNCTIONS

A function is known as a higher order function if it takes in another function
as a parameter, or if it returns a function as an output.

In this case, the `transform` function takes in 1) a normal list and 2) another
function. It then applies the function to every element inside the list, and
returns a new resultant list.
"""
def transform(lis, function):
  output = []
  for n in lis:
    x = function(n)
    output.append(x)
  return output

def square(number):
  return number**2

lis = [1, 2, 3, 4, 5]
new = transform(lis, square)
# [1, 4, 9, 16, 25]


"""
A FUNCTION THAT RETURNS ANOTHER FUNCTION

A function that returns another function as an output is also known as a higher
order function. One prevalent usage of this would be decorators. Here’s a
simple example:
"""
def add_exclamation(function):
  '''
  The `add_exclamation` function is known as a decorator — it takes in a
  function, does some stuff to the output of the original function, and returns
  a new function with additional capability.
  '''
  def inner(*args):
    return function(*args) + "!"
  return inner

@add_exclamation
def greet(name):
  '''
  In this exact case, `add_exclamation` takes in a function that returns a
  string, and adds 1 exclamation mark behind whatever string the original
  function returns.
  '''
  return "hello " + name
x = greet("bob")
print(x)  # hello bob!

"""
A FUNCTION THAT RETURNS ANOTHER FUNCTION THAT RETURNS ANOTHER FUNCTION

This is a useful way for us to customize our decorator functions.
"""
def add(char):
  '''
  Here, calling the `add` function `returns a decorator function`. In this
  exact case, we pass a string `char` to the `add` function. The decorator that
  is returned adds `char` to the back of whatever function it decorates.
  '''
  def decorator(function):
    def inner(*args):
      return function(*args) + char
    return inner
  return decorator

@add("!")
def greet(name):
  return "hello " + name
x = greet("bob")
print(x)  # hello bob!

