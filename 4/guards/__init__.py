
class GuardSleep():
    guard = 0
    date = ""
    duration = 0
    start_minute = 0

    def __init__(self, guard, date, start_minute, duration=0):
        self.guard = guard
        self.date = date
        self.duration = duration

    def wakes(self, minute):
        self.duration = minute - self.start_minute 


def parse_sorted(text_input):
    return []

def find_most_sleepy(guards_asleep):
    return 0

def find_most_sleepy_minute(guard, guards_asleep):
    return 0
