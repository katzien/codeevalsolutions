<?php
$fh = fopen($argv[1], "r");

while (($test = fgets($fh)) !== false) {

    $parts = explode(',', $test);

    $str = trim($parts[0]);
    $chars = str_split(trim($parts[1]));

    $pattern = '/';
    foreach ($chars as $c) {
        $pattern .= $c . '|';
    }

    $pattern = rtrim($pattern, '|') . '/';

    echo preg_replace($pattern, '', $str) . PHP_EOL;
}

fclose($fh);
exit(0);