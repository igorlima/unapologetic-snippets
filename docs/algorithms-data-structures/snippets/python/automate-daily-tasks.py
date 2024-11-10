"""
PYTHON SCRIPTS TO AUTOMATE DAILY TASKS
"""


"""
Clipboard History Logger (No More Lost Copy-Paste)
Sat, Nov 09, 2024 - 2024a11m09d
- https://blog.stackademic.com/13-python-scripts-to-automate-your-daily-tasks-d52f79430146
Ever copied something important, only to accidentally overwrite it? This script logs everything you copy, so you’ll never lose it again.
Fun Fact: The average person presses “Ctrl+C” at least 100 times a day — probably more if you’re a developer.
"""
import pyperclip
import time
def log_clipboard():
  previous = ""
  while True:
    current = pyperclip.paste()
    if current != previous:
      with open("clipboard_log.txt", "a") as log:
        log.write(f"{current}\n")
      previous = current
    time.sleep(1)
log_clipboard()

"""
Smart Break Reminder (Take Care of That Spine)
Sat, Nov 09, 2024 - 2024a11m09d
- https://blog.stackademic.com/13-python-scripts-to-automate-your-daily-tasks-d52f79430146
Keep your posture in check with gentle reminders for breaks — automatically.
"""
import time
from plyer import notification
def remind_to_move():
  while True:
    notification.notify(
      title='Time to Move!',
      message='Stand up, stretch, and save your back.',
      timeout=10
    )
    time.sleep(3600)  # Every hour
remind_to_move()
