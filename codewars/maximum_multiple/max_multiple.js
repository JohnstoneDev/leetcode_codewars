/*
	Given a divisor & a Bound. Find the largest integer N
	such that :

	N is divisible by divisor
	N <= bound
	n > 0
*/

function MaxMultiple(divisor, bound) {
	let numbers = []

	for (let i = 1; i <= bound; i++) {
		if (i % divisor === 0) {
			numbers.push(i)
		}
	}

	return numbers[numbers.length - 1]
}


console.log(MaxMultiple(2,7), "<= Got, Expected 6")