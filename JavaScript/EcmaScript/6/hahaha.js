// A classic for loop
let ha = "";
for(let i = 0; i < 3; i++){
	ha += "ha";
}
console.log(ha);

// A messed up nested for loop
ha = "";
for(let i = 0; i < 1; i++){
	ha += "ha";
	for(let i = 0; i < 1; i++){
		ha += "ha";
		for(let i = 0; i < 1; i++){
			ha += "ha";
		}
	}
}
console.log(ha);

// A simple while loop
ha = "";
let counter = 0;
while(counter < 3){
	ha += "ha";
	counter += 1;
}
console.log(ha);

// A messed up nested while loop
ha = "";
let counter1 = 0;
while(counter1 < 1){
	counter2 = 0;
	ha += "ha";
	while(counter2 < 1){
		counter3 = 0;
		ha += "ha";
		while(counter3 < 1){
			ha += "ha";
			counter3 += 1;
		}
		counter2 += 1;
	}
	counter1 += 1;
}
console.log(ha);
