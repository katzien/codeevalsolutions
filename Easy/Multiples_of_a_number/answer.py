import sys

test_cases = open(sys.argv[1], 'r')

for test in test_cases:
    if test:
        testList = test.split(',')

        x = int(testList[0])
        n = int(testList[1])

        r = n

        while (r < x):
            r += n

        print(r)

test_cases.close()
