
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
    for line in text_input:
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

def find_guards_most_sleepy_minute(guard, guards_asleep):
    guards_times = filter( lambda x: x.guard == guard, guards_asleep)

    minutes = {}

    for gt in guards_times:
        max_min = gt.date.minute + gt.duration 

        for minute in range(gt.date.minute, max_min):
            if minute in minutes:
                minutes[minute] += 1
            else:
                minutes[minute] = 1

    return max(minutes, key=lambda d: minutes[d])

def find_most_sleepy_minute(guards_asleep):
    # FIXME: This is very similar to most sleepy min
    guards_sleep_minutes = {}

    for gt in guards_asleep:
        guard = gt.guard

        if guard not in guards_sleep_minutes:
            guards_sleep_minutes[guard] = {}

        max_min = gt.date.minute + gt.duration

        for minute in range(gt.date.minute, max_min+1):
            if minute in guards_sleep_minutes[guard]:
                guards_sleep_minutes[guard][minute] += 1
            else:
                guards_sleep_minutes[guard][minute] = 1

    guards_max_sleep = {}
    for guard, minutes in guards_sleep_minutes.items():
        if minutes == {}:
            guards_max_sleep[guard] = (0,0)
            continue

        most_frequent_minute = max(minutes, key=lambda d: minutes[d])
        minute_freq = minutes[most_frequent_minute]
        guards_max_sleep[guard] = (most_frequent_minute, minute_freq)

    print(guards_max_sleep)
    max_sleeper = max(guards_max_sleep, key=lambda d: guards_max_sleep[d][1])
    max_sleep = guards_max_sleep[max_sleeper][0]

    return (max_sleep, max_sleeper)


def main():
    lines = ""
    with open("input.txt") as infile:
        lines = infile.readlines()

    guards = parse_sorted(lines)
    guard = find_most_sleepy(guards)
    sleepiest = find_guards_most_sleepy_minute(guard, guards)
    print(sleepiest * guard)

    sleepy_minute, sleepy_guard = find_most_sleepy_minute(guards)
    print(sleepy_minute, sleepy_guard)
    print(sleepy_minute * sleepy_guard)


if __name__ == "__main__":
    main()
