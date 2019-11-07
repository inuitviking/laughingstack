<?php
/*
 * This is probably the simplest way to do a for loop
*/
for($i = 0; $i < 3; $i++){
	echo "ha";
}

// For the folks who don't like the above way of doing things
for($i = 0; $i < 3; $i++)
{
	echo "ha";
}

/*
 * Why not have THREE NESTED FOR LOOPS EACH WRITING "ha"?!
*/
for($i = 0; $i < 1; $i++){
	echo "ha";
	for($i = 0; $i < 1; $i++){
		echo "ha";
		for($i = 0; $i < 1; $i++){
			echo "ha";
		}
	}
}

echo "hahaha"; // this is just boring

/*
 * While loops are plenty fun
*/
$counter = 0;
while($counter !== 3){
	echo "ha";
	$counter += 1;
}

/*
 * Oh whoops, while loops got a bit too fun
*/
$firsti = 0;
while($firsti !== 1){
	echo "ha";
	$secondi = 0;
	while($secondi !== 1){
		echo "ha";
		$thirdi = 0;
		while($thirdi !== 1){
			echo "ha";
			$thirdi += 1;
		}
		$secondi += 1;
	}
	$firsti += 1;
}
