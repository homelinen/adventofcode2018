from guards import GuardSleep, parse_sorted, find_most_sleepy

def test_parse_input():


    text_input = """
        [1518-11-01 00:00] Guard #10 begins shift
        [1518-11-01 00:05] falls asleep
        [1518-11-01 00:25] wakes up
        [1518-11-01 00:30] falls asleep
        [1518-11-01 00:55] wakes up
        [1518-11-01 23:58] Guard #99 begins shift
        [1518-11-02 00:40] falls asleep
        [1518-11-02 00:50] wakes up
        [1518-11-03 00:05] Guard #10 begins shift
        [1518-11-03 00:24] falls asleep
        [1518-11-03 00:29] wakes up
        [1518-11-04 00:02] Guard #99 begins shift
        [1518-11-04 00:36] falls asleep
        [1518-11-04 00:46] wakes up
        [1518-11-05 00:03] Guard #99 begins shift
        [1518-11-05 00:45] falls asleep
        [1518-11-05 00:55] wakes up
    """

    guards_asleep = [
                GuardSleep(10, "1518-11-01", 5, 20),
                GuardSleep(99, "1518-11-02", 40, 10),
                GuardSleep(10, "1518-11-03", 24, 5),
                GuardSleep(99, "1518-11-04", 36, 10),
                GuardSleep(99, "1518-11-05", 45, 10),
                ]

    assert guards_asleep == parse_sorted(text_input)
    

def test_find_most_sleep():
    guards_asleep = [
                GuardSleep(10, "1518-11-01", 5, 20),
                GuardSleep(99, "1518-11-02", 40, 10),
                GuardSleep(10, "1518-11-03", 24, 5),
                GuardSleep(99, "1518-11-04", 36, 10),
                GuardSleep(99, "1518-11-05", 45, 10),
                ]

    most_asleep = 10

    assert most_asleep == find_most_sleepy(guards_asleep)
