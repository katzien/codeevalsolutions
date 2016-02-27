<?php
$fh = fopen($argv[1], "r");

while (($test = fgets($fh)) !== false) {

  list($x, $n) = explode(',', trim($test));

  $r = $n;

  while ($r < $x) {
    $r += $n;
  }

  echo $r . PHP_EOL;
}

fclose($fh);
exit(0);
