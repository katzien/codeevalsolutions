<?php

$results = array();

$file = fopen($argv[1], "r");

$testCasesCount = 0;

while ((($line = fgets($file)) !== false) && $testCasesCount < 20) {
    $lineNumbers = explode(' ', $line);
    $x = (int)$lineNumbers[0]; // Fizz
    $y = (int)$lineNumbers[1]; // Buzz
    $max = (int)$lineNumbers[2];

    $lineResults = '';

    if ($x < 1 || $x > 20) {
        $lineResults = 'X is outside of the allowed range [1, 20]';
    } else if ($y < 1 || $y > 20) {
        $lineResults = 'Y is outside of the allowed range [1, 20]';
    } else if ($max < 21 || $max > 100) {
        $lineResults = 'N is outside of the allowed range [21, 100]';
    } else {
        for ($i = 1; $i <= $max; $i++) {

            $output = $i;

            if ($i % $x === 0) {
                $output = 'F';

                if ($i % $y === 0) {
                    $output .= 'B';
                }
            } else {
                if ($i % $y === 0) {
                    $output = 'B';
                }
            }

            $lineResults .= $output . ' ';
        }
    }

    $results[] = trim($lineResults);
    $testCasesCount++;
}

fclose($file);

$numResults = count($results);
for ($j = 0; $j < $numResults; $j++) {
    echo $results[$j];

    if ($j < $numResults-1) {
        echo PHP_EOL;
    }
}

exit(0);