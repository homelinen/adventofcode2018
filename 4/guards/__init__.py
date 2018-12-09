
from datetime import datetime
import re

class GuardSleep():
    guard = 0
    duration = 0
    date = None
    slept = False

    def __init__(self, guard, date=datetime.today(), duration=0):
        self.guard = int(guard)
        self.duration = duration
        self.date = date

    def wakes(self, date):
        self.duration = parse_date(date).minute - self.date.minute

    def sleep(self, date):
        self.date = parse_date(date)
        self.slept = True

    def __str__(self):
        return f'{self.guard} {self.date} {self.duration}'

    def __repr__(self):
        return f'GuardSleep({self.guard}, {self.date}, {self.duration}, {self.slept})'

    def __eq__(self, other):
        return int(self.guard) == int(other.guard) and self.date == other.date and self.duration == other.duration


def parse_date(datestr):
    date = datetime.strptime(datestr, '%Y-%m-%d %H:%M')
    return date


def parse_sorted(text_input):

    guards_sleep = []
    guard_sleep = None
    for line in text_input.split("\n"):
        match = re.search(r'\[(\d+-\d+-\d+ \d+:\d+)] (.*)$', line)
        date = match[1]
        action = match[2]

        if "Guard" in action:
            guard = re.search(r'#(\d+)', action)
            guard_sleep = GuardSleep(guard[1])
            guards_sleep.append(guard_sleep)

        if "falls asleep" in action and guard_sleep.slept:
            guard_sleep = GuardSleep(guard_sleep.guard)
            guards_sleep.append(guard_sleep)

        if "falls asleep" in action:
            guard_sleep.sleep(date)
        if "wakes up" in action:
            guard_sleep.wakes(date)

        # print(f"line: {line}\n{guard_sleep}")

    return guards_sleep

def find_most_sleepy(guards_asleep):

    sleeper = {}
    for asleep in guards_asleep:
        if asleep.guard in sleeper:
            sleeper[asleep.guard] += asleep.duration
        else:
            sleeper[asleep.guard] = asleep.duration
    return max(sleeper, key=lambda d: sleeper[d])

def find_most_sleepy_minute(guard, guards_asleep):
    return 0
