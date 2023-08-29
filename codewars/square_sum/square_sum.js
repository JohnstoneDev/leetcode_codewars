function SquareSum(numbers) {
	let sum = 0;

	numbers.forEach(number => {
		sum += (number * number)
	});

	return sum
}

console.log(SquareSum([0,3,4,5]))