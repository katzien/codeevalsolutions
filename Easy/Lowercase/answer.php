<?php
$fh = fopen($argv[1], "r");

while (true) {
    $test = fgets($fh);

    if (empty($test)) {
        break;
    }

    echo strtolower($test) . PHP_EOL;
}

fclose($fh);
exit(0);