"""
python3 -m venv tmp-env
source tmp-env/bin/activate
#
# pip3 show plumbum
# pip3 index versions plumbum
# pip3 list
#
pip3 install plumbum==1.9.0
"""

# https://github.com/tomerfiliba/plumbum
from plumbum import local

# RUNNING BASIC OPERATING SYSTEM COMMANDS
cmd = local["echo"]
print(cmd)
result, stdout, stderr = cmd["Hello, World!"].run()
print(result)
print(stdout)
print(10*"-")

# RUNNING AN O/S COMMAND WITH ARGUMENTS
from plumbum import local
from glob import glob
import os
cmd = local["ls"]["-l"]
files = glob(os.path.expanduser("./*"))
if files:
  result, stdout, stderr = cmd[files].run()
  print(result)
  print(stdout)
print(10*"-")

# CREATING PIPELINES
# create the `ls` and `grep` commands
cmd1 = local["ls"]
cmd2 = local["grep"]
# create the pipeline and execute it
pipeline = (cmd1["-l"] | cmd2["-E"]["."])
print(pipeline)
stdout = pipeline()  # capture the output from the pipeline
# print the output
print(stdout)
print(10*"-")

# RE-DIRECTING INPUT AND OUTPUT
cmd1 = local["ls"]
cmd2 = local["cat"]
pipeline = (cmd1["-l"] | cmd2) > "output.txt"
print(pipeline)
stdout = pipeline()  # capture the output from the pipeline
print(stdout)
print(10*"-")

# CHECK OUTPUT REDIRECTION
cmd1 = local["cat"]
# run the command and capture the output, including stderr if needed
result = cmd1["output.txt"].run()
code, stdout, stderr = result
# print the standard output
print(stdout)
print(10*"-")

# DEALING WITH EXIT CODES AND ERRORS
cmd=local["cat"]
try:
  result = cmd1["/home/tom/nofile.txt"].run()
  result
except Exception as e:
  print(e)
print(10*".")
# to avoid this error, we can pass the special parameter `retcode = xxx` as a keyword argument to our commands.
# setting `xxx` to `None` means that plumbum will not check for errors when executing a command.
result = cmd1["/home/tom/nofile.txt"].run(retcode=None)
print(result)
print(10*"-")

# FOREGROUND, BACKGROUND AND NOHUP PROCESSES
from plumbum import local, BG
# create the `ls` and `sed` commands
ls = local["ls"]
sed = local["sed"]
# create a pipeline to run a recursive file listing
# `ls -lR /mnt/d` and limit the output to 10 lines
process = (ls["-lR", "/"] | sed ["10q"]) & BG(timeout=5, retcode=None)
# wait for the process to complete
process.wait()
# print the output and errors directly
if process.stdout:
    print("Output:")
    print(process.stdout)
if process.stderr:
    print("Errors:")
    print(process.stderr)
print("the process has completed.")
print(10*"-")

"""
REFERENCE:
- GitHub:
  - https://github.com/tomerfiliba/plumbum
- Useful Python libraries you might not know existed — Plumbum.
  - https://levelup.gitconnected.com/useful-python-libraries-you-might-not-know-existed-plumbum-01be50863bd3
"""
