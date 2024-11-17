"""
Python provides a nice, built-in functions called open() that can help you
interact with pre-existing files or write out files yourself.

In a way you can:
- open files
- read files
- write files
- append to files

The open() function's arguments and defaults:
  ```python
  open(file, mode='r', buffering=-1, encoding=None,
       errors=None, newline=None, closefd=True, opener=None)
  ```

The following table goes over the other modes that can be used when opening a file:
- 'r' open file for reading (default)
- 'w' open for writing. If file exists, replace its contents
- 'a' open for writing. If file exists, appends to end
- 'b' binary mode
- 't' text mode (default)
- '+' reading and writing

There are two primary methods used to open a file. You can do something like this:
  ```python
  file_handler = open('example.txt')
  # do something with the file
  file_handler.close()
  ```

  ```python
  try:
    file_handler = open('example.txt')
  except:
    # ignore the error, print a warning, or log the exception
    pass
  finally:
    file_handler.close()
  ```
Using a context manager:
  ```python
  with open('example.txt') as file_handler:
    # do something with the handler here
    data = file_handler.read()
  ```

"""

DIVIDER = 10*'-'

# READING FILES
print(DIVIDER)
print('most common way to read files')
print(DIVIDER)
with open('./aux/file-sample.txt') as file_handler:
  for line in file_handler:
    print(line)
print('\n')

# an alternative way to loop over the lines
# ...
# but be careful with this method because
# depending on how much RAM your machine has in it, you may run
# out of memory. This is why the first method is recommended.
print(DIVIDER)
print('an alternative way to loop over the lines at once')
print(DIVIDER)
with open('./aux/file-sample.txt') as file_handler:
  lines = file_handler.readlines()
  for line in lines:
    print(line)
print('\n')

# if you know the file is pretty small, there is another way to
# read the entire file into memory:
# ...
# the `read()` method will read the entire file into memory and
# assign it to your variable.
print(DIVIDER)
print('read the entire file into memory')
print(DIVIDER)
with open('./aux/file-sample.txt') as file_handler:
  file_contents = file_handler.read()
print('\n')

# OCCASIONALLY you may want to read a file in smaller or larger
# chunks. This can be done by specifying the the size in bytes to
# `read()`.
# quit()
print(DIVIDER)
print('specifying the size in bytes to `read()`')
print(DIVIDER)
with open('./aux/file-sample.txt') as file_handler:
  while True:
    data = file_handler.read(128)
    # OR:
    # data = file_handler.read(1024)
    if not data:
      break
    print(data)
print('\n')

# READING BINARY FILES
print(DIVIDER)
print('reading binary files')
print(DIVIDER)
with open('./aux/python.jpeg', 'rb') as file_handler:
  file_contents = file_handler.read()
print('\n')


# WRITING FILES
"""
WRITING FILES

Writing a new file in Python uses pretty much the exact same syntax as reading.
But instead of setting the mode to `r`, you set it to `w` for write-mode. If you
need to write in binary mode, then you would open the file in `wb` mode.

WARNING: When using the `w` and `wb` modes, if the file already exists, you will
end up overwriting it. Python does not warn you in any way. Python does provide
a way to check for a file's existence by using the os module via
`os.path.exists()`.
"""
print(DIVIDER)
print('writing files')
print(DIVIDER)
with open('./aux/temp.txt', 'w') as file_handler:
  file_handler.write('This is a test. ')
# verify it worked by reading the file and printing out its contents:
with open('./aux/temp.txt') as file_handler:
  print(file_handler.read())
print('\n')

# APPENDING TO FILES
print(DIVIDER)
print('appending to files')
print(DIVIDER)
with open('./aux/temp.txt', 'a') as file_handler:
  file_handler.write('Here is some more text')
with open('./aux/temp.txt') as file_handler:
  print(file_handler.read())
print('\n')

# CATCHING FILE EXCEPTIONS
print(DIVIDER)
print('catching file exceptions')
print(DIVIDER)
try:
  with open('./aux/example.txt') as file_handler:
    for line in file_handler:
      print(line)
except OSError:
  print('An error has occurred')
except Exception as e:
  print('The error type is:', e.__class__.__name__)
  print(f'An error has occurred: {str(type(e))}')
  print(f'Error message: {str(e)}')
print('\n')
