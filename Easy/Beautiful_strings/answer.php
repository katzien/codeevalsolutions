<?php

$fh = fopen($argv[1], "r");
while (($test = fgets($fh)) !== false) {

  $lettersOnly = preg_replace("/[^a-zA-Z]+/", "", $test);
  $letters = str_split(strtolower($lettersOnly));

  $totals = array();

  foreach ($letters as $l) {
    if (isset($totals[$l])) {
      $totals[$l] += 1;
    } else {
      $totals[$l] = 1;
    }
  }

  rsort($totals);

  $score = 0;
  $highestPoints = 26;
  foreach($totals as $frequency) {
    $score += ($highestPoints * $frequency);
    $highestPoints--;
  }

  echo $score . PHP_EOL;
}

fclose($fh);
exit(0);